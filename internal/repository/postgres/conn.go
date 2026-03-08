package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"

	"github.com/ayupov-ayaz/shortly/internal/helper/environment"
)

type ConnPool struct {
	*pgxpool.Pool
}

type Config struct {
	Env               environment.Env
	DSN               string
	MaxConns          int32
	MinConns          int32
	MaxConnLifetime   time.Duration
	MaxConnIdleTime   time.Duration
	HealthCheckPeriod time.Duration
}

func NewConnPool(ctx context.Context, cfg Config) (*ConnPool, error) {
	poolCfg, err := pgxpool.ParseConfig(cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("parsing config: %w", err)
	}

	configure(poolCfg, cfg)

	pool, err := pgxpool.NewWithConfig(ctx, poolCfg)
	if err != nil {
		return nil, fmt.Errorf("creating connection pool: %w", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		pool.Close()
		return nil, fmt.Errorf("pinging database: %w", err)
	}

	return &ConnPool{Pool: pool}, nil
}

func configure(poolCfg *pgxpool.Config, cfg Config) {
	poolCfg.MaxConns = cfg.MaxConns
	poolCfg.MinConns = cfg.MinConns
	poolCfg.MaxConnLifetime = cfg.MaxConnLifetime
	poolCfg.MaxConnIdleTime = cfg.MaxConnIdleTime
	poolCfg.HealthCheckPeriod = cfg.HealthCheckPeriod

	logLevel := tracelog.LogLevelInfo
	if cfg.Env.IsDevelopment() {
		logLevel = tracelog.LogLevelDebug
	}

	// log configuration for debugging purposes
	poolCfg.ConnConfig.Tracer = &tracelog.TraceLog{
		Logger:   &Logger{},
		LogLevel: logLevel,
		Config: &tracelog.TraceLogConfig{
			TimeKey: "time",
		},
	}
}

func (c *ConnPool) Close() error {
	if c.Pool != nil {
		c.Pool.Close()
	}

	return nil
}
