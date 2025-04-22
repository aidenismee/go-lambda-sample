package server

import (
	"context"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"

	"github.com/aidenismee/go-lambda-sample/pkg/config"
	"github.com/aidenismee/go-lambda-sample/pkg/router"

	internalRouter "github.com/aidenismee/go-lambda-sample/internal/router"
)

func InitLambdaFunctionHandler() *echoadapter.EchoLambda {
	config := config.LoadConfig()

	ctx := context.Background()
	router.New(config.RouterConfig, internalRouter.SetupRouter(ctx))

	router := router.New(config.RouterConfig, internalRouter.SetupRouter(ctx))

	return echoadapter.New(router.Engine())
}

func StartServer() {
	config := config.LoadConfig()

	ctx := context.Background()
	router.New(config.RouterConfig, internalRouter.SetupRouter(ctx))

	router.New(config.RouterConfig, internalRouter.SetupRouter(ctx)).Start()
}
