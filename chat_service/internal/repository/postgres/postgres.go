package postgres

import (
	"context"
	"fmt"

	"chat_service/internal/config"

	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	*pgxpool.Pool
}

func New(cfg config.Postgres) (*Postgres, error) {
	pgConfig, err := pgxpool.ParseConfig(cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	pgConfig.MaxConns = cfg.MaxConns
	pgConfig.MinConns = cfg.MinConns
	pgConfig.MaxConnLifetime = cfg.MaxConnLifetime
	pgConfig.MaxConnIdleTime = cfg.MaxConnIdleTime
	pgConfig.HealthCheckPeriod = cfg.HealthCheckPeriod
	pgConfig.ConnConfig.ConnectTimeout = cfg.ConnectTimeout

	pool, err := pgxpool.NewWithConfig(context.Background(), pgConfig)
	if err != nil {
		return nil, fmt.Errorf("create connection pool: %w", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("ping: %w", err)
	}

	return &Postgres{
		pool,
	}, nil
}

func (p *Postgres) Close() {
	p.Close()
}
