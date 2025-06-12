package di

import (
	"context"
	"log/slog"
	"os"

	"github.com-test/internal/handler"
	"github.com-test/internal/repository"
	"github.com-test/internal/services"
	"github.com/labstack/echo/v4"
)

func Init(_ context.Context) (*echo.Echo, error) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	userStore := repository.NewUserMemory()
	userService := services.NewUserService(userStore, *logger)

	server := echo.New()

	server.HTTPErrorHandler = handler.CustomHTTPErrorHandler(server.DefaultHTTPErrorHandler)

	handler := handler.NewUserHandler(userService)

	group := server.Group("api/v1/users")

	group.POST("/create", handler.PostUser)
	group.GET("/get-all", handler.GetALLUsers)

	return server, nil
}
