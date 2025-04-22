package config

import "os"

type Config struct {
	Env          string
	RouterConfig RouterConfig
}

type RouterConfig struct {
	Port         string
	ReadTimeout  int
	WriteTimeout int
}

func LoadConfig() *Config {
	appConfig := Config{
		Env: os.Getenv("ENV"),
		RouterConfig: RouterConfig{
			Port: os.Getenv("PORT"),
		},
	}

	return &appConfig
}
