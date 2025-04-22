package server

import (
	"context"

	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"

	internalRouter "github.com/aidenismee/go-lambda-sample/internal/router"
	"github.com/aidenismee/go-lambda-sample/pkg/config"
	routerPkg "github.com/aidenismee/go-lambda-sample/pkg/router"
)

var router routerPkg.Router

func initServer() {
	config := config.LoadConfig()

	ctx := context.Background()
	router = routerPkg.New(config.RouterConfig, internalRouter.SetupRouter(ctx))
}

func InitLambdaHandler() *echoadapter.EchoLambda {
	initServer()

	return echoadapter.New(router.Engine())
}

func StartServer() {
	initServer()

	router.Start()
}
