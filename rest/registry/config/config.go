package config

const (
	KeyString = "key"
	ValueString = "value"
	ConstKey   = "ping"
	ConstValue = "pong"
	Endpoint = "/syncmap"
)

type Config struct {
	Environment string
	LogLevel    string
	HttpPort    string
	Constant    map[string]string
}
