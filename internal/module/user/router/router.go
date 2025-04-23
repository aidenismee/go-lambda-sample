package router

import (
	"github.com/aidenismee/go-lambda-sample/internal/module/user/handler"
	"github.com/aidenismee/go-lambda-sample/internal/module/user/service"
	"github.com/aidenismee/go-lambda-sample/internal/repository"
	"github.com/aidenismee/go-lambda-sample/pkg/router"
	"github.com/aidenismee/go-lambda-sample/pkg/static"
	"github.com/labstack/echo/v4"
)

func SetupUserRouter(routerCtx router.RouterCtx, routerGroup *echo.Group) error {
	db := routerCtx.GetDependencies(static.DBCtxKey)

	userGroup := routerGroup.Group("/users")

	userRepo := repository.NewUserRepository(db)

	userService := service.NewUserService(userRepo)

	userHandler := handler.NewUserHandler(userService)

	userGroup.GET("/:id", userHandler.GetDetailUser)

	return nil
}
