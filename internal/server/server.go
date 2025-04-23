package server

import (
	"github.com/aidenismee/go-lambda-sample/db"
	"github.com/aidenismee/go-lambda-sample/pkg/static"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"

	internalRouter "github.com/aidenismee/go-lambda-sample/internal/router"
	"github.com/aidenismee/go-lambda-sample/pkg/config"
	routerPkg "github.com/aidenismee/go-lambda-sample/pkg/router"
)

var router routerPkg.Router

func initServer() {
	config := config.LoadConfig()

	database := db.NewDB(config.DatabaseConfig)

	routerCtx := routerPkg.NewRouterCtx().
		SetDependencies(static.DBCtxKey, database.GetClient())

	router = routerPkg.New(config.RouterConfig,
		internalRouter.SetupRouter(routerCtx))
}

func InitLambdaHandler() *echoadapter.EchoLambda {
	initServer()

	return echoadapter.New(router.Engine())
}

func StartServer() {
	initServer()

	router.Start()
}
