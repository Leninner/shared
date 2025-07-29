package config

import (
	"database/sql"
	"log/slog"
	"sync"
	"time"
)

type Config struct {
	Port int
	Env  string
	DB   struct {
		DSN          string
		MaxOpenConns int
		MaxIdleConns int
		MaxIdleTime  time.Duration
	}
	Limiter struct {
		RPS     float64
		Burst   int
		Enabled bool
	}
	CORS struct {
		TrustedOrigins []string
	}
}

type Application struct {
	Config Config
	Logger *slog.Logger
	WG     sync.WaitGroup
	DataSource *sql.DB
} 