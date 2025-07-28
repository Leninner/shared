package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerConfig struct {
	Level      string
	Environment string
	ServiceName string
	OutputPath  string
	Encoding    string
}

type Logger struct {
	*zap.Logger
}

func NewLogger(config LoggerConfig) (*Logger, error) {
	level, err := zapcore.ParseLevel(config.Level)
	if err != nil {
		level = zapcore.InfoLevel
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	var encoder zapcore.Encoder
	if config.Encoding == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	var writeSyncer zapcore.WriteSyncer
	if config.OutputPath != "" {
		file, err := os.OpenFile(config.OutputPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, err
		}
		writeSyncer = zapcore.AddSync(file)
	} else {
		writeSyncer = zapcore.AddSync(os.Stdout)
	}

	core := zapcore.NewCore(encoder, writeSyncer, level)

	fields := []zap.Field{
		zap.String("service", config.ServiceName),
		zap.String("environment", config.Environment),
	}

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	logger = logger.With(fields...)

	return &Logger{Logger: logger}, nil
}

func (l *Logger) With(fields ...zap.Field) *Logger {
	return &Logger{Logger: l.Logger.With(fields...)}
}

func (l *Logger) WithContext(ctx map[string]interface{}) *Logger {
	fields := make([]zap.Field, 0, len(ctx))
	for k, v := range ctx {
		fields = append(fields, zap.Any(k, v))
	}
	return l.With(fields...)
} 