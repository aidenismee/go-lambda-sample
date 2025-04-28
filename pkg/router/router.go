package router

import (
	"context"
	"fmt"
	"github.com/aidenismee/go-lambda-sample/pkg/config"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Router interface {
	Start()
	Engine() *echo.Echo
	Close()
}

type router struct {
	config config.RouterConfig
	engine *echo.Echo
}

type RouteFunc func(router *echo.Echo) error

func New(config config.RouterConfig, registerRoute RouteFunc) Router {
	engine := setUpRouter(&config)

	if err := registerRoute(engine); err != nil {
		log.Fatal(err)
	}

	return &router{
		engine: engine,
		config: config,
	}
}

func setUpRouter(config *config.RouterConfig) *echo.Echo {
	engine := echo.New()

	if config == nil {
		config = loadDefaultConfig()
	}

	engine.Server.Addr = fmt.Sprintf(":%d", config.Port)
	engine.Server.ReadTimeout = time.Duration(config.ReadTimeout) * time.Minute
	engine.Server.WriteTimeout = time.Duration(config.WriteTimeout) * time.Minute

	return engine
}

func (s *router) Start() {
	go s.Close()

	log.Printf("Server is listening at %v", s.config.Port)
	if err := s.engine.StartServer(s.engine.Server); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

func (s *router) Engine() *echo.Echo {
	return s.engine
}

func (s *router) Close() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.engine.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
