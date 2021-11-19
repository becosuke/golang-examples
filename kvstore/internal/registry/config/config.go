package config

import (
	"os"
)

const (
	KeyString   = "key"
	ValueString = "value"
	Endpoint    = "/node"
)

type Config struct {
	constConfig
	envConfig
}

func NewConfig() *Config {
	return &Config{
		constConfig: newConstConfig(),
		envConfig:   newEnvConfig(),
	}
}

type constConfig struct {
	KeyString   string
	ValueString string
	Endpoint    string
}

type envConfig struct {
	Environment string
	LogLevel    string
	HttpPort    string
}

func newConstConfig() constConfig {
	return constConfig{
		KeyString:   KeyString,
		ValueString: ValueString,
		Endpoint:    Endpoint,
	}
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
