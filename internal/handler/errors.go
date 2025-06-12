package handler

import (
	"errors"
	"net/http"

	"github.com-test/internal/services"
	"github.com/labstack/echo/v4"
)

func CustomHTTPErrorHandler(next echo.HTTPErrorHandler) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		asErr := &services.Error{}
		if !errors.As(err, asErr) {
			next(err, c)
			return
		}

		msg := echo.Map{
			"code": asErr.Code,
			"mgs":  asErr.Msg,
		}

		_ = c.JSON(http.StatusBadRequest, msg)
		return
	}
}
