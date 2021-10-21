package model

import "time"

type AppConfig struct {
	AppInfo           AppInfo           `mapstructure:"app"`
	UpstreamConfig    UpstreamConfig    `mapstructure:"upstream"`
	MongoConfig       MongoConfig       `mapstructure:"mongo"`
	FiberConfig       FiberConfig       `mapstructure:"fiber"`
	QueueConfig       QueueConfig       `mapstructure:"queue"`
	LoggingConfig     LoggingConfig     `mapstructure:"logging"`
	HealthCheckConfig HealthCheckConfig `mapstructure:"health-check"`
}

type AppInfo struct {
	Name   string `mapstructure:"name"`
	Banner bool   `mapstructure:"banner"`
}

type UpstreamConfig struct {
	BaseURL string `mapstructure:"base-url"`
	HMACKey string `mapstructure:"hmac-key"`
}

type MongoConfig struct {
	URI         string `mapstructure:"uri"`
	Database    string `mapstructure:"database"`
	MinPoolSize uint64 `mapstructure:"min-pool-size"`
	MaxPoolSize uint64 `mapstructure:"max-pool-size"`
	MaxIdleTime uint64 `mapstructure:"max-idle-time"`
}

type FiberConfig struct {
	Address      string `mapstructure:"address"`
	ReadTimeout  uint64 `mapstructure:"read-timeout"`
	WriteTimeout uint64 `mapstructure:"write-timeout"`
	IdleTimeout  uint64 `mapstructure:"idle-timeout"`
}

type QueueConfig struct {
	MaxParallel int `mapstructure:"max-parallel"`
}

type LoggingConfig struct {
	LogLevel string `mapstructure:"level"`
}

type HealthCheckConfig struct {
	CacheDuration   time.Duration `mapstructure:"cache-duration"`
	RefreshInterval time.Duration `mapstructure:"refresh-interval"`
	InitialDelay    time.Duration `mapstructure:"initial-delay"`
}
