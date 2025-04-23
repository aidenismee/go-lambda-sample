package config

import "os"

type Config struct {
	Env            string
	RouterConfig   RouterConfig
	DatabaseConfig DatabaseConfig
}

type RouterConfig struct {
	Port         string
	ReadTimeout  int
	WriteTimeout int
}

type DatabaseConfig struct {
	DatabaseLog bool
	DatabasePsn string
}

func LoadConfig() *Config {
	appConfig := Config{
		Env: os.Getenv("ENV"),
		RouterConfig: RouterConfig{
			Port: os.Getenv("PORT"),
		},
		DatabaseConfig: DatabaseConfig{
			DatabaseLog: true,
			DatabasePsn: os.Getenv("DATABASE_PSN"),
		},
	}

	return &appConfig
}
