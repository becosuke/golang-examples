package config

import (
	"os"
)

const (
	KeyString   = "key"
	ValueString = "value"
	Endpoint    = "/syncmap"
	ConstKey    = "ping"
	ConstValue  = "pong"
)

type Config struct {
	Constant map[string]string
	envConfig
}

func NewConfig() *Config {
	return &Config{
		Constant:  newConstant(),
		envConfig: newEnvConfig(),
	}
}

func newConstant() map[string]string {
	return map[string]string{
		ConstKey: ConstValue,
	}
}

type envConfig struct {
	Environment string
	LogLevel    string
	HttpPort    string
}

func newEnvConfig() envConfig {
	environment, ok := os.LookupEnv("ENVIRONMENT")
	if !ok {
		environment = "development"
	}
	logLevel, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		logLevel = "info"
	}
	httpPort, ok := os.LookupEnv("HTTP_PORT")
	if !ok {
		httpPort = "8080"
	}

	return envConfig{
		Environment: environment,
		LogLevel:    logLevel,
		HttpPort:    httpPort,
	}
}
