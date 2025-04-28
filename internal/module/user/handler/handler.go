package handler

import (
	"context"
	"github.com/aidenismee/go-lambda-sample/internal/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type IUserService interface {
	GetDetailUser(ctx context.Context, id int) (*model.User, error)
}
type UserHandler struct {
	UserService IUserService
}

func NewUserHandler(userService IUserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (h *UserHandler) GetDetailUser(ctx echo.Context) error {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	user, err := h.UserService.GetDetailUser(ctx.Request().Context(), id)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, user)

}
