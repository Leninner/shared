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
	Kafka KafkaConfig
}

type KafkaConfig struct {
	BootstrapServers string
	ClientID         string
	GroupID          string
	Topics           TopicConfig
}

type TopicConfig struct {
	PaymentRequest              string
	PaymentResponse             string
	RestaurantApprovalRequest   string
	RestaurantApprovalResponse  string
}

type Application struct {
	Config Config
	Logger *slog.Logger
	WG     sync.WaitGroup
	DataSource *sql.DB
}

func NewDefaultConfig() Config {
	return Config{
		Port: 4000,
		Env:  "development",
		DB: struct {
			DSN          string
			MaxOpenConns int
			MaxIdleConns int
			MaxIdleTime  time.Duration
		}{
			DSN:          "postgres://user:password@localhost/dbname?sslmode=disable",
			MaxOpenConns: 25,
			MaxIdleConns: 25,
			MaxIdleTime:  15 * time.Minute,
		},
		Limiter: struct {
			RPS     float64
			Burst   int
			Enabled bool
		}{
			RPS:     100,
			Burst:   100,
			Enabled: true,
		},
		CORS: struct {
			TrustedOrigins []string
		}{
			TrustedOrigins: []string{"http://localhost:3000"},
		},
		Kafka: KafkaConfig{
			BootstrapServers: "localhost:9092",
			ClientID:         "microservice",
			GroupID:          "microservice-group",
			Topics: TopicConfig{
				PaymentRequest:              "payment-request",
				PaymentResponse:             "payment-response",
				RestaurantApprovalRequest:   "restaurant-approval-request",
				RestaurantApprovalResponse:  "restaurant-approval-response",
			},
		},
	}
} 