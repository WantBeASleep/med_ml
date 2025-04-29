package config

type Config struct {
	App      App
	DB       DB
	Yookassa Yookassa
}

type App struct {
	Url string `env:"APP_URL" env-default:"localhost:50055"`
}

type DB struct {
	Dsn string `env:"DB_DSN" env-required:"true"`
}

type Yookassa struct {
	AccountID string `env:"YOOKASSA_ACCOUNT_ID" env-required:"true"`
	SecretKey string `env:"YOOKASSA_ACCESS_KEY" env-required:"true"`
	ReturnURL string `env:"YOOKASSA_RETURN_URL" env-required:"true"`
}
