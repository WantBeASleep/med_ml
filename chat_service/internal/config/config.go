package config

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	HTTP       HTTP
	GRPC       GRPC
	WSClient   WSClient
	Postgres   Postgres
	Redis      Redis
	MedService MedService
}

type HTTP struct {
	Host string
	Port uint16
}

type GRPC struct {
	Host string
	Port uint16
}

type WSClient struct {
	PingInterval time.Duration
	PongWait     time.Duration
}

type Postgres struct {
	DSN               string
	MaxConns          int32
	MinConns          int32
	MaxConnLifetime   time.Duration
	MaxConnIdleTime   time.Duration
	HealthCheckPeriod time.Duration
	ConnectTimeout    time.Duration
}

type Redis struct {
	Addr         string
	Password     string
	DB           int
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PoolSize     int
	MinIdleConns int
}

type MedService struct {
	Addr     string
	CacheTTL time.Duration
}

func Parse(configPath string) (Config, error) {
	v := viper.New()

	if configPath != "" {
		v.SetConfigFile(configPath)

		if err := v.ReadInConfig(); err != nil {
			log.Error().
				Err(err).
				Msg("read config file")
		}
	}

	v.AutomaticEnv()

	v.BindEnv("HTTP.Host", "HTTP_HOST")
	v.BindEnv("HTTP.Port", "HTTP_PORT")
	v.BindEnv("GRPC.Host", "GRPC_HOST")
	v.BindEnv("GRPC.Port", "GRPC_PORT")
	v.BindEnv("WSClient.PingInterval", "WS_PING_INTERVAL")
	v.BindEnv("WSClient.PongWait", "WS_PONG_WAIT")
	v.BindEnv("Postgres.DSN", "POSTGRES_DSN")
	v.BindEnv("Redis.Addr", "REDIS_ADDR")
	v.BindEnv("Redis.Password", "REDIS_PASSWORD")
	v.BindEnv("Redis.DB", "REDIS_DB")
	v.BindEnv("MedService.Addr", "MED_SERVICE_ADDR")
	v.BindEnv("MedService.CacheTTL", "MED_SERVICE_CACHE_TTL")

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return Config{}, fmt.Errorf("parse config: %w", err)
	}

	return config, nil
}
