package router

import "github.com/aidenismee/go-lambda-sample/pkg/config"

const (
	defaultPort         = "8080"
	defaultReadTimeout  = 5
	defaultWriteTimeout = 3
)

func loadDefaultConfig() *config.RouterConfig {
	return &config.RouterConfig{
		Port:         defaultPort,
		ReadTimeout:  defaultReadTimeout,
		WriteTimeout: defaultWriteTimeout,
	}
}
