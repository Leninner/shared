package logger

func NewDevelopmentLogger(serviceName string) (*Logger, error) {
	config := LoggerConfig{
		Level:       "debug",
		Environment: "development",
		ServiceName: serviceName,
		Encoding:    "console",
	}
	return NewLogger(config)
}

func NewProductionLogger(serviceName string) (*Logger, error) {
	config := LoggerConfig{
		Level:       "info",
		Environment: "production",
		ServiceName: serviceName,
		Encoding:    "json",
	}
	return NewLogger(config)
}

func NewTestLogger(serviceName string) (*Logger, error) {
	config := LoggerConfig{
		Level:       "debug",
		Environment: "test",
		ServiceName: serviceName,
		Encoding:    "console",
	}
	return NewLogger(config)
} 