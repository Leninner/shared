# Shared Logger

A reusable logger implementation using Zap that can be initialized in each microservice with customization options.

## Features

- Structured logging with Zap
- Customizable log levels, encoding, and output
- Environment-specific configurations
- Service name and environment tagging
- Context-aware logging with fields

## Usage

### Basic Usage

```go
import "github.com/leninner/shared/logger"

// Create a development logger
devLogger, err := logger.NewDevelopmentLogger("order-service")
if err != nil {
    panic(err)
}

// Create a production logger
prodLogger, err := logger.NewProductionLogger("order-service")
if err != nil {
    panic(err)
}
```

### Custom Configuration

```go
config := logger.LoggerConfig{
    Level:       "debug",
    Environment: "staging",
    ServiceName: "order-service",
    OutputPath:  "/var/log/order-service.log",
    Encoding:    "json",
}

customLogger, err := logger.NewLogger(config)
if err != nil {
    panic(err)
}
```

### Using the Logger

```go
// Basic logging
logger.Info("Order created successfully")
logger.Error("Failed to process payment")

// With structured fields
logger.Info("Order validated",
    zap.String("orderId", orderID),
    zap.String("customerId", customerID),
    zap.Int("totalItems", itemCount))

// With context
ctx := map[string]interface{}{
    "requestId": "req-123",
    "userId":    "user-456",
}
contextLogger := logger.WithContext(ctx)
contextLogger.Info("Processing request")
```

### Integration with Domain Services

```go
type OrderDomainServiceImpl struct {
    logger *logger.Logger
}

func NewOrderDomainServiceImpl(logger *logger.Logger) *OrderDomainServiceImpl {
    return &OrderDomainServiceImpl{
        logger: logger,
    }
}

func (s *OrderDomainServiceImpl) ProcessOrder(order *entity.Order) error {
    s.logger.Info("Processing order",
        zap.String("orderId", order.GetID().GetValue().String()))
    
    // ... business logic
    
    s.logger.Info("Order processed successfully",
        zap.String("orderId", order.GetID().GetValue().String()))
    
    return nil
}
```

## Configuration Options

- **Level**: Log level (debug, info, warn, error, fatal, panic)
- **Environment**: Environment name (development, staging, production)
- **ServiceName**: Microservice name for identification
- **OutputPath**: File path for log output (empty for stdout)
- **Encoding**: Log format (json, console)

## Factory Functions

- `NewDevelopmentLogger(serviceName)`: Creates a development logger with console output and debug level
- `NewProductionLogger(serviceName)`: Creates a production logger with JSON output and info level
- `NewTestLogger(serviceName)`: Creates a test logger with console output and debug level 