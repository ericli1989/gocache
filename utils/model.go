package utils

var (
	GoCacheConfig *Config
)

type WebConfig struct {
	Ip   string
	Port int
}

type Config struct {
	deployMode 	string
	Web         *WebConfig
}
