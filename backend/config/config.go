package config

type Config struct {
	Port int `json:"port"`
}

func NewDefaultConfig() *Config {
	return &Config{
		Port: 8080,
	}
}
