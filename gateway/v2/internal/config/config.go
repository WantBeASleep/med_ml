package config

type Config struct {
	App      App
	Adapters Adapters
	S3       S3
	Dbus     Dbus
}

type App struct {
	Url string `env:"APP_URL" env-default:"localhost:8080"`
}

type Adapters struct {
	UziUrl string `env:"ADAPTERS_UZIURL" env-required:"true"`
}

type S3 struct {
	Endpoint     string `env:"S3_ENDPOINT" env-required:"true"`
	Access_Token string `env:"S3_TOKEN_ACCESS" env-required:"true"`
	Secret_Token string `env:"S3_TOKEN_SECRET" env-required:"true"`
}

type Dbus struct {
	Addrs []string `env:"DBUS_ADDRS" env-required:"true"`
}
