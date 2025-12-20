package config

type Config struct {
	DefaultBalance float64
}

func NewConfig() *Config {
	return &Config{1000}
}
