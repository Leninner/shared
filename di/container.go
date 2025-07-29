package di

import (
	"database/sql"
	"log/slog"
	"sync"

	"github.com/leninner/infrastructure/kafka"
	"github.com/leninner/shared/config"
)

type SharedContainer struct {
	mu       sync.RWMutex
	config   config.Config
	logger   *slog.Logger
	db       *sql.DB
	kafka    *kafka.KafkaModule
	services map[string]interface{}
}

func NewSharedContainer(cfg config.Config, log *slog.Logger) *SharedContainer {
	return &SharedContainer{
		config:   cfg,
		logger:   log,
		services: make(map[string]interface{}),
	}
}

func (c *SharedContainer) SetDatabase(db *sql.DB) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.db = db
}

func (c *SharedContainer) SetKafka(kafkaModule *kafka.KafkaModule) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.kafka = kafkaModule
}

func (c *SharedContainer) RegisterService(name string, service interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.services[name] = service
}

func (c *SharedContainer) GetService(name string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	service, exists := c.services[name]
	return service, exists
}

func (c *SharedContainer) GetConfig() config.Config {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.config
}

func (c *SharedContainer) GetLogger() *slog.Logger {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.logger
}

func (c *SharedContainer) GetDatabase() *sql.DB {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.db
}

func (c *SharedContainer) GetKafka() *kafka.KafkaModule {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.kafka
}

func (c *SharedContainer) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	if c.db != nil {
		return c.db.Close()
	}
	return nil
} 