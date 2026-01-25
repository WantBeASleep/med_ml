package config

type Config struct {
	App App
	S3  S3
}

type App struct {
	URL         string `env:"APP_URL" env-default:"localhost:50080"`
	MediaPath   string `env:"MEDIA_PATH" env-default:"./media"`
	TileSize    int    `env:"TILE_SIZE" env-default:"510"` // Соответствует citology/bidder (tile_size=510)
	Overlap     int    `env:"OVERLAP" env-default:"1"`     // Соответствует citology/bidder (overlap=1)
	LimitBounds bool   `env:"LIMIT_BOUNDS" env-default:"true"` // Соответствует citology/bidder (limit_bounds=True)
}

type S3 struct {
	Endpoint     string `env:"S3_ENDPOINT" env-required:"true"`
	Access_Token string `env:"S3_TOKEN_ACCESS" env-required:"true"`
	Secret_Token string `env:"S3_TOKEN_SECRET" env-required:"true"`
	BucketName   string `env:"S3_BUCKET_NAME" env-default:"cytology"`
}
