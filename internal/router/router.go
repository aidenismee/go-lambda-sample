package router

import (
	"github.com/labstack/echo/v4"

	userRouter "github.com/aidenismee/go-lambda-sample/internal/module/user/router"
	"github.com/aidenismee/go-lambda-sample/pkg/router"
)

func SetupRouter(ctx router.RouterCtx) router.RouteFunc {
	return func(router *echo.Echo) error {
		apiGroup := router.Group("/api")
		{
			v1Group := apiGroup.Group("v1")
			{
				if err := userRouter.SetupUserRouter(ctx, v1Group); err != nil {
					return err
				}
			}
		}

		return nil
	}
}
