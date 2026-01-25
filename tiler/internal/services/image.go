package services

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"image"
	"io"
	"log/slog"
	"math"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"tiler/internal/domain"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/disintegration/imaging"
	_ "golang.org/x/image/tiff" // Регистрирует TIFF декодер для fallback
)

type ImageService interface {
	GetImageInfo(ctx context.Context, imagePath string) (*domain.ImageInfo, error)
	GetTile(ctx context.Context, imagePath string, level, col, row int, format string) (*domain.Tile, error)
	GetDZI(ctx context.Context, imagePath string) (*domain.DZI, error)
}

type cachedImage struct {
	path      string
	image     *vips.ImageRef
	lastUsed  time.Time
	accessMux sync.RWMutex
}

type imageService struct {
	s3Client     S3Client
	tileSize     int
	overlap      int
	cache        map[string]*cachedImage
	cacheMux     sync.RWMutex
	infoCache    map[string]*domain.ImageInfo // Кэш информации об изображениях
	infoCacheMux sync.RWMutex
	cacheDir     string
	tileCacheDir string // Директория для кэширования тайлов
	maxCacheSize int
	cacheTTL     time.Duration
}

type S3Client interface {
	GetObject(ctx context.Context, bucketName, objectName string) ([]byte, error)
	GetObjectStream(ctx context.Context, bucketName, objectName string) (io.ReadCloser, error)
}

func NewImageService(s3Client S3Client, tileSize, overlap int) ImageService {
	// Инициализируем libvips (требуется только один раз)
	vips.Startup(nil)

	// Создаем временную директорию для кэша файлов
	cacheDir := filepath.Join(os.TempDir(), "tiler_cache")
	os.MkdirAll(cacheDir, 0755)

	// Создаем директорию для кэширования тайлов
	tileCacheDir := filepath.Join(os.TempDir(), "tiler_tile_cache")
	os.MkdirAll(tileCacheDir, 0755)

	return &imageService{
		s3Client:     s3Client,
		tileSize:     tileSize,
		overlap:      overlap,
		cache:        make(map[string]*cachedImage),
		infoCache:    make(map[string]*domain.ImageInfo),
		cacheDir:     cacheDir,
		tileCacheDir: tileCacheDir,
		maxCacheSize: 10,        // Максимум 10 открытых файлов в кэше
		cacheTTL:     time.Hour, // Время жизни кэша
	}
}

// GetImageInfo возвращает информацию об изображении, включая максимальный уровень масштабирования
//
// МАКСИМАЛЬНЫЙ УРОВЕНЬ ПРИБЛИЖЕНИЯ для SVS файлов:
//
// Максимальный уровень масштабирования зависит от трех факторов:
//
// 1. РАЗМЕР ИЗОБРАЖЕНИЯ (width x height)
//   - Чем больше исходное разрешение слайда, тем больше уровней
//   - SVS файлы обычно имеют очень высокое разрешение (десятки тысяч пикселей)
//   - Например: слайд 100000x80000 пикселей даст больше уровней, чем 10000x8000
//
// 2. РАЗМЕР ТАЙЛА (tileSize)
//
//   - Стандартный размер тайла: 256x256 пикселей
//
//   - Чем меньше tileSize, тем больше уровней (но больше тайлов на каждом уровне)
//
//   - Чем больше tileSize, тем меньше уровней (но меньше тайлов)
//
//     3. ФОРМУЛА ВЫЧИСЛЕНИЯ:
//     levels = ceil(log2(max(width, height) / tileSize)) + 1
//
//     Где:
//
//   - max(width, height) - максимальная сторона изображения
//
//   - tileSize - размер тайла (обычно 256)
//
//   - log2 - логарифм по основанию 2
//
//   - ceil - округление вверх
//
//   - +1 добавляет уровень 0 (полное разрешение)
//
// КАК РАБОТАЮТ УРОВНИ:
// - Level 0: полное разрешение (100% масштаб)
// - Level 1: 50% масштаб (изображение уменьшено в 2 раза)
// - Level 2: 25% масштаб (изображение уменьшено в 4 раза)
// - Level N: изображение уменьшено в 2^N раз
// - Максимальный уровень: когда изображение становится меньше или равно размеру тайла
//
// ДЛЯ SVS ФАЙЛОВ:
// - SVS файлы (формат Aperio) уже содержат встроенные пирамидальные уровни
// - libvips автоматически использует эти уровни при чтении файла (random access)
// - Это позволяет эффективно читать только нужные тайлы без загрузки всего файла
// - Максимальный уровень определяется исходным разрешением сканирования слайда
//
// ПРИМЕРЫ РАСЧЕТА:
// - Изображение 256x256, tileSize=256:  levels = ceil(log2(256/256)) + 1 = ceil(0) + 1 = 1 уровень (level 0)
// - Изображение 512x512, tileSize=256:  levels = ceil(log2(512/256)) + 1 = ceil(1) + 1 = 2 уровня (level 0-1)
// - Изображение 10000x8000, tileSize=256: levels = ceil(log2(10000/256)) + 1 = ceil(5.29) + 1 = 7 уровней (level 0-6)
// - Изображение 100000x80000, tileSize=256: levels = ceil(log2(100000/256)) + 1 = ceil(8.61) + 1 = 10 уровней (level 0-9)
//
// КАК УЗНАТЬ МАКСИМАЛЬНЫЙ УРОВЕНЬ:
// 1. Вызовите GetImageInfo() - вернет структуру ImageInfo с полем Levels
// 2. Максимальный доступный уровень = Levels - 1 (так как нумерация с 0)
// 3. Например, если Levels = 10, то доступны уровни от 0 до 9
func (s *imageService) GetImageInfo(ctx context.Context, imagePath string) (*domain.ImageInfo, error) {
	// Проверяем кэш информации об изображении
	s.infoCacheMux.RLock()
	cachedInfo, exists := s.infoCache[imagePath]
	s.infoCacheMux.RUnlock()

	if exists && cachedInfo != nil {
		// Возвращаем кэшированную информацию
		return &domain.ImageInfo{
			Width:    cachedInfo.Width,
			Height:   cachedInfo.Height,
			Levels:   cachedInfo.Levels,
			TileSize: cachedInfo.TileSize,
			Overlap:  cachedInfo.Overlap,
		}, nil
	}

	// Используем кэшированное изображение для получения метаданных
	// libvips читает только заголовок файла, не весь файл
	img, err := s.getOrLoadImage(ctx, imagePath)
	if err != nil {
		return nil, err
	}
	defer img.Close()

	width := img.Width()
	height := img.Height()

	// Вычисляем количество уровней (levels) по формуле Deep Zoom Image (DZI)
	// Формула: levels = ceil(log2(max(width, height) / tileSize)) + 1
	// Каждый уровень масштабирования уменьшает изображение в 2 раза
	// Максимальный уровень - это когда изображение становится меньше или равно размеру тайла
	levels := s.calculateMaxLevels(width, height)

	// Для очень маленьких изображений гарантируем минимум 1 уровень
	if levels < 1 {
		levels = 1
	}

	info := &domain.ImageInfo{
		Width:    width,
		Height:   height,
		Levels:   levels,
		TileSize: s.tileSize,
		Overlap:  s.overlap,
	}

	// Сохраняем в кэш
	s.infoCacheMux.Lock()
	s.infoCache[imagePath] = info
	s.infoCacheMux.Unlock()

	// Возвращаем копию
	return &domain.ImageInfo{
		Width:    info.Width,
		Height:   info.Height,
		Levels:   info.Levels,
		TileSize: info.TileSize,
		Overlap:  info.Overlap,
	}, nil
}

// calculateMaxLevels вычисляет максимальное количество уровней масштабирования
// для изображения заданного размера
//
// Формула: levels = ceil(log2(max(width, height) / tileSize)) + 1

func (s *imageService) calculateMaxLevels(width, height int) int {
	maxDim := math.Max(float64(width), float64(height))
	if maxDim <= 0 {
		return 1
	}
	// Базовый расчет уровней по формуле Deep Zoom Image
	levels := int(math.Ceil(math.Log2(maxDim/float64(s.tileSize)))) + 1

	// OpenSeadragon может запрашивать дополнительные уровни из-за особенностей
	// расчета координат и округления. Добавляем один дополнительный уровень,
	// если изображение достаточно большое, чтобы это имело смысл
	// Это позволяет обрабатывать запросы уровня, который теоретически не нужен,
	// но может быть запрошен OpenSeadragon при определенных условиях масштабирования
	if maxDim > float64(s.tileSize*2) {
		// Добавляем запас только для достаточно больших изображений
		// чтобы избежать создания ненужных уровней для маленьких изображений
		levels++
	}

	if levels < 1 {
		return 1
	}
	return levels
}

// getOrLoadImage получает изображение из кэша или загружает его
func (s *imageService) getOrLoadImage(ctx context.Context, imagePath string) (*vips.ImageRef, error) {
	// Проверяем кэш
	s.cacheMux.RLock()
	cached, exists := s.cache[imagePath]
	s.cacheMux.RUnlock()

	if exists {
		cached.accessMux.Lock()
		cached.lastUsed = time.Now()
		// Создаем копию изображения для использования
		imgCopy, err := cached.image.Copy()
		cached.accessMux.Unlock()
		if err == nil {
			return imgCopy, nil
		}
		// Если копирование не удалось, удаляем из кэша и загружаем заново
		s.cacheMux.Lock()
		delete(s.cache, imagePath)
		s.cacheMux.Unlock()
	}

	// Файл не в кэше, загружаем его
	localPath := filepath.Join(s.cacheDir, strings.ReplaceAll(imagePath, "/", "_"))

	// Проверяем, существует ли файл локально
	if _, err := os.Stat(localPath); os.IsNotExist(err) {
		// Используем streaming для загрузки из S3 (более эффективно для больших файлов)
		stream, err := s.s3Client.GetObjectStream(ctx, "", imagePath)
		if err != nil {
			return nil, fmt.Errorf("failed to get image stream from S3 (path: %s): %w", imagePath, err)
		}
		defer stream.Close()

		// Создаем файл и копируем данные напрямую из потока
		outFile, err := os.Create(localPath)
		if err != nil {
			return nil, fmt.Errorf("failed to create cache file: %w", err)
		}
		defer outFile.Close()

		// Копируем данные из потока в файл
		_, err = io.Copy(outFile, stream)
		if err != nil {
			os.Remove(localPath) // Удаляем частично загруженный файл
			return nil, fmt.Errorf("failed to save image to cache: %w", err)
		}
	}

	// Открываем файл через libvips с random access
	// libvips автоматически будет читать только нужные тайлы из tiled TIFF
	vipsImg, err := vips.NewImageFromFile(localPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open image with libvips: %w", err)
	}

	// Добавляем в кэш
	s.cacheMux.Lock()
	// Очищаем старые записи, если кэш переполнен
	if len(s.cache) >= s.maxCacheSize {
		s.cleanupCache()
	}
	s.cache[imagePath] = &cachedImage{
		path:     localPath,
		image:    vipsImg,
		lastUsed: time.Now(),
	}
	s.cacheMux.Unlock()

	// Возвращаем копию для использования
	return vipsImg.Copy()
}

// cleanupCache удаляет старые записи из кэша
func (s *imageService) cleanupCache() {
	now := time.Now()
	for path, cached := range s.cache {
		if now.Sub(cached.lastUsed) > s.cacheTTL {
			cached.accessMux.Lock()
			cached.image.Close()
			cached.accessMux.Unlock()
			delete(s.cache, path)
		}
	}
}

func (s *imageService) GetTile(ctx context.Context, imagePath string, level, col, row int, format string) (*domain.Tile, error) {
	// Сначала получаем информацию об изображении для валидации
	info, err := s.GetImageInfo(ctx, imagePath)
	if err != nil {
		return nil, err
	}

	// Проверяем, что уровень валиден
	// Включаем информацию об изображении в сообщение об ошибке для отладки
	if level < 0 || level >= info.Levels {
		return nil, fmt.Errorf("invalid level: %d (max: %d, image_size: %dx%d, calculated_levels: %d)",
			level, info.Levels-1, info.Width, info.Height, info.Levels)
	}

	// Вычисляем размеры изображения на данном уровне
	var levelWidth, levelHeight int
	if level == 0 {
		levelWidth = info.Width
		levelHeight = info.Height
	} else {
		scale := math.Pow(2, float64(level))
		levelWidth = int(float64(info.Width) / scale)
		levelHeight = int(float64(info.Height) / scale)
	}

	// Проверяем, что размеры изображения валидны
	if levelWidth <= 0 || levelHeight <= 0 {
		return nil, fmt.Errorf("invalid image dimensions at level %d: %dx%d", level, levelWidth, levelHeight)
	}

	// Вычисляем максимальное количество тайлов на данном уровне
	// Даже если изображение меньше тайла, должен быть хотя бы один тайл
	maxCol := int(math.Max(1, math.Ceil(float64(levelWidth)/float64(s.tileSize))))
	maxRow := int(math.Max(1, math.Ceil(float64(levelHeight)/float64(s.tileSize))))

	// Проверяем границы координат тайла
	// OpenSeadragon может запрашивать тайлы с координатами, которые выходят за границы,
	// особенно на высоких уровнях масштабирования, где изображение становится очень маленьким.
	// В таких случаях возвращаем пустой белый тайл, чтобы избежать наложения одинаковых тайлов.
	if col < 0 || row < 0 || col >= maxCol || row >= maxRow {
		// Возвращаем пустой белый тайл для запросов вне границ
		return s.createEmptyTile(level, col, row, format)
	}

	// Проверяем кэш тайлов на диске
	tileCacheKey := fmt.Sprintf("%s_%d_%d_%d.%s", strings.ReplaceAll(imagePath, "/", "_"), level, col, row, format)
	tileCachePath := filepath.Join(s.tileCacheDir, tileCacheKey)

	// Пытаемся прочитать тайл из кэша
	if cachedTile, err := os.ReadFile(tileCachePath); err == nil {
		return &domain.Tile{
			Level:  level,
			Col:    col,
			Row:    row,
			Data:   cachedTile,
			Format: format,
		}, nil
	}

	// Получаем изображение из кэша или загружаем его
	// libvips будет читать только нужные тайлы из файла благодаря random access
	baseImg, err := s.getOrLoadImage(ctx, imagePath)
	if err != nil {
		return nil, err
	}
	defer baseImg.Close()

	// Масштабируем изображение до нужного уровня
	// Для SVS файлов libvips автоматически использует встроенные пирамидальные уровни
	// при чтении файла, но для других форматов нужно масштабировать вручную
	var vipsImg *vips.ImageRef
	if level > 0 {
		scale := math.Pow(2, float64(level))
		scaleFactor := 1.0 / scale
		// Создаем копию для масштабирования, чтобы не модифицировать кэшированное изображение
		scaledImg, err := baseImg.Copy()
		if err != nil {
			return nil, fmt.Errorf("failed to copy image for scaling: %w", err)
		}
		defer scaledImg.Close()

		if err := scaledImg.Resize(scaleFactor, vips.KernelLanczos3); err != nil {
			return nil, fmt.Errorf("failed to resize image: %w", err)
		}
		vipsImg = scaledImg
	} else {
		// Для уровня 0 используем исходное изображение
		vipsImg = baseImg
	}

	// Вычисляем координаты тайла
	x := col * s.tileSize
	y := row * s.tileSize

	// Дополнительная проверка границ после масштабирования (на случай округления)
	if x >= vipsImg.Width() || y >= vipsImg.Height() {
		return nil, errors.New("tile coordinates out of bounds")
	}

	// Обрезаем тайл с учетом overlap
	cropX := int(math.Max(0, float64(x-s.overlap)))
	cropY := int(math.Max(0, float64(y-s.overlap)))
	cropWidth := int(math.Min(float64(vipsImg.Width()-cropX), float64(s.tileSize+2*s.overlap)))
	cropHeight := int(math.Min(float64(vipsImg.Height()-cropY), float64(s.tileSize+2*s.overlap)))

	// Извлекаем тайл
	// ExtractArea модифицирует исходное изображение, поэтому нужно создать копию
	tileImg, err := vipsImg.Copy()
	if err != nil {
		return nil, fmt.Errorf("failed to copy image: %w", err)
	}
	defer tileImg.Close()

	if err := tileImg.ExtractArea(cropX, cropY, cropWidth, cropHeight); err != nil {
		return nil, fmt.Errorf("failed to extract tile: %w", err)
	}

	// Кодируем тайл в нужный формат
	var tileData []byte
	var encodeErr error
	switch strings.ToLower(format) {
	case "jpeg", "jpg":
		ep := vips.NewJpegExportParams()
		ep.Quality = 85 // Оптимизируем качество для баланса между размером и качеством
		tileData, _, encodeErr = tileImg.ExportJpeg(ep)
	case "png":
		ep := vips.NewPngExportParams()
		tileData, _, encodeErr = tileImg.ExportPng(ep)
	default:
		return nil, fmt.Errorf("unsupported format: %s", format)
	}

	if encodeErr != nil {
		return nil, fmt.Errorf("failed to encode tile: %w", encodeErr)
	}

	// Сохраняем тайл в кэш на диске (асинхронно, чтобы не блокировать ответ)
	go func() {
		if err := os.WriteFile(tileCachePath, tileData, 0644); err != nil {
			// Игнорируем ошибки записи в кэш, чтобы не влиять на основной поток
			slog.Debug("failed to cache tile", "path", tileCachePath, "err", err)
		}
	}()

	return &domain.Tile{
		Level:  level,
		Col:    col,
		Row:    row,
		Data:   tileData,
		Format: format,
	}, nil
}

// createEmptyTile создает пустой белый тайл для запросов, которые выходят за границы изображения
// Это предотвращает наложение одинаковых тайлов в OpenSeadragon
// Вместо возврата тайла (0,0) для всех запросов, возвращаем пустой белый тайл,
// который OpenSeadragon правильно отобразит в нужной позиции без наложения
func (s *imageService) createEmptyTile(level, col, row int, format string) (*domain.Tile, error) {
	// Создаем пустое белое изображение размером тайла
	tileSizeWithOverlap := s.tileSize + 2*s.overlap

	// Создаем белое изображение через imaging (более надежный способ)
	whiteImg := imaging.New(tileSizeWithOverlap, tileSizeWithOverlap, image.White)

	// Кодируем в нужный формат
	var tileData []byte
	var encodeErr error
	switch strings.ToLower(format) {
	case "jpeg", "jpg":
		tileData, encodeErr = encodeJPEG(whiteImg)
	case "png":
		tileData, encodeErr = encodePNG(whiteImg)
	default:
		return nil, fmt.Errorf("unsupported format: %s", format)
	}

	if encodeErr != nil {
		return nil, fmt.Errorf("failed to encode empty tile: %w", encodeErr)
	}

	return &domain.Tile{
		Level:  level,
		Col:    col,
		Row:    row,
		Data:   tileData,
		Format: format,
	}, nil
}

// getTileFromStandardImage - fallback метод для стандартного image.Decode
func (s *imageService) getTileFromStandardImage(img image.Image, level, col, row int, format string) (*domain.Tile, error) {
	// Масштабируем изображение до нужного уровня
	scaledImg := s.scaleToLevel(img, level)

	// Вычисляем размеры тайла с учетом overlap
	tileSizeWithOverlap := s.tileSize + 2*s.overlap

	// Вычисляем координаты тайла
	x := col * s.tileSize
	y := row * s.tileSize

	// Проверяем границы
	bounds := scaledImg.Bounds()
	if x >= bounds.Dx() || y >= bounds.Dy() {
		return nil, errors.New("tile coordinates out of bounds")
	}

	// Вычисляем размеры для обрезки
	width := s.tileSize
	height := s.tileSize
	if x+width > bounds.Dx() {
		width = bounds.Dx() - x
	}
	if y+height > bounds.Dy() {
		height = bounds.Dy() - y
	}

	// Обрезаем тайл с учетом overlap
	cropX := math.Max(0, float64(x-s.overlap))
	cropY := math.Max(0, float64(y-s.overlap))
	cropWidth := math.Min(float64(bounds.Dx()-int(cropX)), float64(tileSizeWithOverlap))
	cropHeight := math.Min(float64(bounds.Dy()-int(cropY)), float64(tileSizeWithOverlap))

	tileImg := imaging.Crop(scaledImg, image.Rect(
		int(cropX),
		int(cropY),
		int(cropX)+int(cropWidth),
		int(cropY)+int(cropHeight),
	))

	// Кодируем тайл в нужный формат
	var tileData []byte
	var encodeErr error
	switch strings.ToLower(format) {
	case "jpeg", "jpg":
		tileData, encodeErr = encodeJPEG(tileImg)
	case "png":
		tileData, encodeErr = encodePNG(tileImg)
	default:
		return nil, fmt.Errorf("unsupported format: %s", format)
	}

	if encodeErr != nil {
		return nil, fmt.Errorf("failed to encode tile: %w", encodeErr)
	}

	return &domain.Tile{
		Level:  level,
		Col:    col,
		Row:    row,
		Data:   tileData,
		Format: format,
	}, nil
}

func (s *imageService) scaleToLevel(img image.Image, level int) image.Image {
	if level == 0 {
		return img
	}

	// Вычисляем масштаб для уровня
	scale := math.Pow(2, float64(level))
	bounds := img.Bounds()
	newWidth := int(float64(bounds.Dx()) / scale)
	newHeight := int(float64(bounds.Dy()) / scale)

	return imaging.Resize(img, newWidth, newHeight, imaging.Lanczos)
}

func (s *imageService) GetDZI(ctx context.Context, imagePath string) (*domain.DZI, error) {
	info, err := s.GetImageInfo(ctx, imagePath)
	if err != nil {
		return nil, err
	}

	// Генерируем DZI XML
	xml := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<Image xmlns="http://schemas.microsoft.com/deepzoom/2008"
       TileSize="%d"
       Overlap="%d"
       Format="jpeg"
       ServerFormat="Default">
  <Size Width="%d" Height="%d"/>
</Image>`, s.tileSize, s.overlap, info.Width, info.Height)

	return &domain.DZI{
		XML:       xml,
		ImageInfo: *info,
	}, nil
}

func encodeJPEG(img image.Image) ([]byte, error) {
	// Используем imaging для кодирования
	var buf bytes.Buffer
	err := imaging.Encode(&buf, img, imaging.JPEG)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func encodePNG(img image.Image) ([]byte, error) {
	var buf bytes.Buffer
	err := imaging.Encode(&buf, img, imaging.PNG)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// GetLevelInfo возвращает детальную информацию о конкретном уровне масштабирования
// Полезно для отладки и понимания структуры пирамиды изображения
func (s *imageService) GetLevelInfo(ctx context.Context, imagePath string, level int) (map[string]interface{}, error) {
	info, err := s.GetImageInfo(ctx, imagePath)
	if err != nil {
		return nil, err
	}

	if level < 0 || level >= info.Levels {
		return nil, fmt.Errorf("invalid level: %d (max: %d)", level, info.Levels-1)
	}

	// Вычисляем размеры изображения на данном уровне
	var levelWidth, levelHeight int
	if level == 0 {
		levelWidth = info.Width
		levelHeight = info.Height
	} else {
		scale := math.Pow(2, float64(level))
		levelWidth = int(float64(info.Width) / scale)
		levelHeight = int(float64(info.Height) / scale)
	}

	// Вычисляем количество тайлов на данном уровне
	maxCol := int(math.Max(1, math.Ceil(float64(levelWidth)/float64(s.tileSize))))
	maxRow := int(math.Max(1, math.Ceil(float64(levelHeight)/float64(s.tileSize))))
	totalTiles := maxCol * maxRow

	return map[string]interface{}{
		"level":         level,
		"width":         levelWidth,
		"height":        levelHeight,
		"tiles_cols":    maxCol,
		"tiles_rows":    maxRow,
		"total_tiles":   totalTiles,
		"tile_size":     s.tileSize,
		"scale_factor":  math.Pow(2, float64(level)),
		"is_max_level":  level == info.Levels-1,
		"original_size": fmt.Sprintf("%dx%d", info.Width, info.Height),
	}, nil
}
