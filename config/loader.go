package config

import (
	"flag"
	"os"
	"strconv"
	"strings"
	"time"
)

type ConfigLoader struct {
	serviceName string
	config      Config
}

func NewConfigLoader(serviceName string) *ConfigLoader {
	return &ConfigLoader{
		serviceName: serviceName,
		config:      NewDefaultConfig(),
	}
}

func (l *ConfigLoader) LoadFromFlags() *ConfigLoader {
	flag.IntVar(&l.config.Port, "port", l.config.Port, "API server port")
	flag.StringVar(&l.config.Env, "env", l.config.Env, "Environment (development|staging|production)")
	
	flag.StringVar(&l.config.DB.DSN, "db-dsn", l.config.DB.DSN, "Database connection string")
	flag.IntVar(&l.config.DB.MaxOpenConns, "db-max-open-conns", l.config.DB.MaxOpenConns, "Maximum number of open connections to the database")
	flag.IntVar(&l.config.DB.MaxIdleConns, "db-max-idle-conns", l.config.DB.MaxIdleConns, "Maximum number of idle connections to the database")
	flag.DurationVar(&l.config.DB.MaxIdleTime, "db-max-idle-time", l.config.DB.MaxIdleTime, "Maximum amount of time a connection may be idle before being closed")
	
	flag.Float64Var(&l.config.Limiter.RPS, "limiter-rps", l.config.Limiter.RPS, "Rate limiter requests per second")
	flag.IntVar(&l.config.Limiter.Burst, "limiter-burst", l.config.Limiter.Burst, "Rate limiter burst")
	flag.BoolVar(&l.config.Limiter.Enabled, "limiter-enabled", l.config.Limiter.Enabled, "Enable rate limiter")
	
	flag.StringVar(&l.config.Kafka.BootstrapServers, "kafka-bootstrap-servers", l.config.Kafka.BootstrapServers, "Kafka bootstrap servers")
	flag.StringVar(&l.config.Kafka.ClientID, "kafka-client-id", l.serviceName, "Kafka client ID")
	flag.StringVar(&l.config.Kafka.GroupID, "kafka-group-id", l.serviceName+"-group", "Kafka consumer group ID")
	
	flag.StringVar(&l.config.Kafka.Topics.PaymentRequest, "payment-request-topic-name", l.config.Kafka.Topics.PaymentRequest, "The topic name for the payment request")
	flag.StringVar(&l.config.Kafka.Topics.PaymentResponse, "payment-response-topic-name", l.config.Kafka.Topics.PaymentResponse, "The topic name for the payment response")
	flag.StringVar(&l.config.Kafka.Topics.RestaurantApprovalRequest, "restaurant-approval-request-topic-name", l.config.Kafka.Topics.RestaurantApprovalRequest, "The topic name for the restaurant approval request")
	flag.StringVar(&l.config.Kafka.Topics.RestaurantApprovalResponse, "restaurant-approval-response-topic-name", l.config.Kafka.Topics.RestaurantApprovalResponse, "The topic name for the restaurant approval response")
	
	flag.Parse()
	return l
}

func (l *ConfigLoader) LoadFromEnv() *ConfigLoader {
	if port := os.Getenv("PORT"); port != "" {
		if p, err := strconv.Atoi(port); err == nil {
			l.config.Port = p
		}
	}
	
	if env := os.Getenv("ENV"); env != "" {
		l.config.Env = env
	}
	
	if dsn := os.Getenv("DB_DSN"); dsn != "" {
		l.config.DB.DSN = dsn
	}
	
	if maxOpenConns := os.Getenv("DB_MAX_OPEN_CONNS"); maxOpenConns != "" {
		if m, err := strconv.Atoi(maxOpenConns); err == nil {
			l.config.DB.MaxOpenConns = m
		}
	}
	
	if maxIdleConns := os.Getenv("DB_MAX_IDLE_CONNS"); maxIdleConns != "" {
		if m, err := strconv.Atoi(maxIdleConns); err == nil {
			l.config.DB.MaxIdleConns = m
		}
	}
	
	if maxIdleTime := os.Getenv("DB_MAX_IDLE_TIME"); maxIdleTime != "" {
		if d, err := time.ParseDuration(maxIdleTime); err == nil {
			l.config.DB.MaxIdleTime = d
		}
	}
	
	if rps := os.Getenv("LIMITER_RPS"); rps != "" {
		if r, err := strconv.ParseFloat(rps, 64); err == nil {
			l.config.Limiter.RPS = r
		}
	}
	
	if burst := os.Getenv("LIMITER_BURST"); burst != "" {
		if b, err := strconv.Atoi(burst); err == nil {
			l.config.Limiter.Burst = b
		}
	}
	
	if enabled := os.Getenv("LIMITER_ENABLED"); enabled != "" {
		l.config.Limiter.Enabled = strings.ToLower(enabled) == "true"
	}
	
	if bootstrapServers := os.Getenv("KAFKA_BOOTSTRAP_SERVERS"); bootstrapServers != "" {
		l.config.Kafka.BootstrapServers = bootstrapServers
	}
	
	if clientID := os.Getenv("KAFKA_CLIENT_ID"); clientID != "" {
		l.config.Kafka.ClientID = clientID
	}
	
	if groupID := os.Getenv("KAFKA_GROUP_ID"); groupID != "" {
		l.config.Kafka.GroupID = groupID
	}
	
	if paymentRequestTopic := os.Getenv("PAYMENT_REQUEST_TOPIC"); paymentRequestTopic != "" {
		l.config.Kafka.Topics.PaymentRequest = paymentRequestTopic
	}
	
	if paymentResponseTopic := os.Getenv("PAYMENT_RESPONSE_TOPIC"); paymentResponseTopic != "" {
		l.config.Kafka.Topics.PaymentResponse = paymentResponseTopic
	}
	
	if restaurantApprovalRequestTopic := os.Getenv("RESTAURANT_APPROVAL_REQUEST_TOPIC"); restaurantApprovalRequestTopic != "" {
		l.config.Kafka.Topics.RestaurantApprovalRequest = restaurantApprovalRequestTopic
	}
	
	if restaurantApprovalResponseTopic := os.Getenv("RESTAURANT_APPROVAL_RESPONSE_TOPIC"); restaurantApprovalResponseTopic != "" {
		l.config.Kafka.Topics.RestaurantApprovalResponse = restaurantApprovalResponseTopic
	}
	
	if trustedOrigins := os.Getenv("CORS_TRUSTED_ORIGINS"); trustedOrigins != "" {
		l.config.CORS.TrustedOrigins = strings.Split(trustedOrigins, ",")
	}
	
	return l
}

func (l *ConfigLoader) Build() Config {
	return l.config
}

func LoadConfig(serviceName string) Config {
	loader := NewConfigLoader(serviceName)
	return loader.LoadFromFlags().LoadFromEnv().Build()
} 