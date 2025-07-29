# Shared Configuration System

This module provides a centralized configuration system that can be reused across all microservices in the TanEats platform.

## Features

- **Unified Configuration**: Single configuration structure for all microservices
- **Environment Variable Support**: Load configuration from environment variables
- **Command Line Flags**: Override configuration with command line flags
- **Kafka Integration**: Built-in Kafka configuration for messaging
- **Database Configuration**: Standardized database connection settings
- **Rate Limiting**: Configurable rate limiting settings
- **CORS Support**: Cross-origin resource sharing configuration

## Configuration Structure

```go
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
```

## Usage

### Basic Usage

```go
import "github.com/leninner/shared/config"

// Load configuration for a specific service
cfg := config.LoadConfig("order-service")
```

### Using ConfigLoader

```go
loader := config.NewConfigLoader("order-service")
cfg := loader.LoadFromFlags().LoadFromEnv().Build()
```

### Environment Variables

The system supports the following environment variables:

#### Server Configuration
- `PORT` - API server port
- `ENV` - Environment (development|staging|production)

#### Database Configuration
- `DB_DSN` - Database connection string
- `DB_MAX_OPEN_CONNS` - Maximum number of open connections
- `DB_MAX_IDLE_CONNS` - Maximum number of idle connections
- `DB_MAX_IDLE_TIME` - Maximum idle time for connections

#### Rate Limiter Configuration
- `LIMITER_RPS` - Requests per second
- `LIMITER_BURST` - Burst limit
- `LIMITER_ENABLED` - Enable/disable rate limiter

#### Kafka Configuration
- `KAFKA_BOOTSTRAP_SERVERS` - Kafka bootstrap servers
- `KAFKA_CLIENT_ID` - Kafka client ID
- `KAFKA_GROUP_ID` - Kafka consumer group ID
- `PAYMENT_REQUEST_TOPIC` - Payment request topic name
- `PAYMENT_RESPONSE_TOPIC` - Payment response topic name
- `RESTAURANT_APPROVAL_REQUEST_TOPIC` - Restaurant approval request topic
- `RESTAURANT_APPROVAL_RESPONSE_TOPIC` - Restaurant approval response topic

#### CORS Configuration
- `CORS_TRUSTED_ORIGINS` - Comma-separated list of trusted origins

### Command Line Flags

All configuration can be overridden with command line flags:

```bash
./order-service \
  --port=4000 \
  --env=production \
  --db-dsn="postgres://user:pass@localhost/db" \
  --kafka-bootstrap-servers="kafka:9092" \
  --payment-request-topic="payment-requests"
```

## Default Values

The system provides sensible defaults for all configuration options:

```go
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
```

## Service-Specific Configuration

Each service can extend the base configuration:

```go
type OrderServiceConfig struct {
    config.Config
    OrderSpecificField string
}

func LoadOrderServiceConfig() *OrderServiceConfig {
    baseConfig := config.LoadConfig("order-service")
    return &OrderServiceConfig{
        Config: baseConfig,
        OrderSpecificField: "order-specific-value",
    }
}
```

## Integration with DI Container

The configuration system integrates seamlessly with the shared DI container:

```go
import (
    "github.com/leninner/shared/config"
    "github.com/leninner/shared/di"
)

// Load configuration
cfg := config.LoadConfig("order-service")

// Create shared container
container := di.NewSharedContainer(cfg, logger)

// Use in service-specific container
orderContainer := &OrderContainer{
    SharedContainer: container,
    // ... service-specific dependencies
}
```

## Best Practices

1. **Use Environment Variables for Production**: Store sensitive configuration in environment variables
2. **Use Command Line Flags for Development**: Override configuration during development
3. **Service-Specific Defaults**: Each service should provide sensible defaults
4. **Configuration Validation**: Validate configuration on startup
5. **Centralized Configuration**: Use the shared configuration for common settings

## Migration Guide

### From Service-Specific Configuration

1. Replace service-specific config with shared config
2. Update environment variable names
3. Update command line flag names
4. Test configuration loading

### Example Migration

**Before:**
```go
type AppConfig struct {
    Database DatabaseConfig
    Kafka    KafkaConfig
    Topics   TopicConfig
    Server   ServerConfig
}
```

**After:**
```go
type AppConfig struct {
    config.Config
}
```

This reduces code duplication and ensures consistency across all microservices. 