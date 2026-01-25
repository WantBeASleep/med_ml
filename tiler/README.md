# Tiler Service

Микросервис для работы с тайлами изображений, аналогичный bidder из citology.

## Описание

Сервис предоставляет HTTP API для работы с большими изображениями через DeepZoom формат:
- Получение DZI XML метаданных
- Получение тайлов изображений на разных уровнях масштабирования

## API Endpoints

### GET /dzi/{file_path}
Возвращает DZI XML метаданные для изображения.

Пример: `GET /dzi/path/to/image.svs`

### GET /dzi/{file_path}/files/{level}/{col}_{row}.{format}
Возвращает тайл изображения.

Параметры:
- `level` - уровень масштабирования (0 - оригинальное изображение)
- `col` - номер колонки тайла
- `row` - номер строки тайла
- `format` - формат изображения (jpeg, png)

Пример: `GET /dzi/path/to/image.svs/files/5/10_20.jpeg`

**Примечание:** Формат URL отличается от citology/bidder (`/files` вместо `_files`) для совместимости с фронтендом.

## Конфигурация

Переменные окружения:
- `APP_URL` - адрес и порт сервера (по умолчанию: `localhost:50080`)
- `MEDIA_PATH` - путь к медиа файлам (не используется, файлы берутся из S3)
- `TILE_SIZE` - размер тайла в пикселях (по умолчанию: `510`, соответствует citology/bidder)
- `OVERLAP` - перекрытие тайлов в пикселях (по умолчанию: `1`, соответствует citology/bidder)
- `LIMIT_BOUNDS` - ограничивать границы (по умолчанию: `true`)
- `S3_ENDPOINT` - endpoint S3 хранилища (обязательно)
- `S3_TOKEN_ACCESS` - access token для S3 (обязательно)
- `S3_TOKEN_SECRET` - secret token для S3 (обязательно)
- `S3_BUCKET_NAME` - имя bucket в S3 (по умолчанию: `cytology`)

## Запуск

```bash
go run cmd/service/main.go
```

Или через Docker:
```bash
docker build -t tiler .
docker run -p 50080:50080 tiler
```
