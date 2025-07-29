package di

import (
	"database/sql"
	"log/slog"

	"github.com/leninner/infrastructure/kafka"
	"github.com/leninner/shared/config"
	"github.com/leninner/shared/server"
)

type SharedApplicationBuilder struct {
	config config.Config
	logger *slog.Logger
	db     *sql.DB
	kafka  *kafka.KafkaModule
}

func NewSharedApplicationBuilder() *SharedApplicationBuilder {
	return &SharedApplicationBuilder{}
}

func (b *SharedApplicationBuilder) WithConfig(cfg config.Config) *SharedApplicationBuilder {
	b.config = cfg
	return b
}

func (b *SharedApplicationBuilder) WithLogger(log *slog.Logger) *SharedApplicationBuilder {
	b.logger = log
	return b
}

func (b *SharedApplicationBuilder) WithDatabase() *SharedApplicationBuilder {
	if b.config.Port == 0 {
		panic("config must be set before database")
	}
	
	db, err := server.OpenDB(b.config)
	if err != nil {
		if b.logger != nil {
			b.logger.Error("Failed to open database", "error", err)
		}
		panic(err)
	}
	b.db = db
	return b
}

func (b *SharedApplicationBuilder) WithKafka() *SharedApplicationBuilder {
	b.kafka = kafka.NewKafkaModule()
	return b
}

func (b *SharedApplicationBuilder) Build() *SharedContainer {
	if b.logger == nil {
		panic("logger must be set")
	}
	
	container := NewSharedContainer(b.config, b.logger)
	
	if b.db != nil {
		container.SetDatabase(b.db)
	}
	
	if b.kafka != nil {
		container.SetKafka(b.kafka)
	}
	
	return container
}

func (b *SharedApplicationBuilder) GetConfig() config.Config {
	return b.config
}

func (b *SharedApplicationBuilder) GetLogger() *slog.Logger {
	return b.logger
}

func (b *SharedApplicationBuilder) GetDatabase() *sql.DB {
	return b.db
}

func (b *SharedApplicationBuilder) GetKafka() *kafka.KafkaModule {
	return b.kafka
} 