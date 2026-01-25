package domain

import "time"

type ImageInfo struct {
	Width    int
	Height   int
	Levels   int
	TileSize int
	Overlap  int
}

type Tile struct {
	Level  int
	Col    int
	Row    int
	Data   []byte
	Format string // "jpeg" or "png"
}

type DZI struct {
	XML       string
	ImageInfo ImageInfo
}

type CachedImage struct {
	Path       string
	Info       ImageInfo
	LastAccess time.Time
}
