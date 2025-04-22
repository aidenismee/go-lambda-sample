package router

import (
	"context"

	"github.com/labstack/echo/v4"

	"github.com/aidenismee/go-lambda-sample/pkg/router"
)

func SetupRouter(ctx context.Context) router.RouteFunc {
	return func(router *echo.Echo) error {
		apiGroup := router.Group("/api")
		{
			_ = apiGroup.Group("v1")
			{

			}
		}

		return nil
	}
}
