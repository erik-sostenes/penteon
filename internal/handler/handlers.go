package handler

import (
	"net/http"

	"github.com-test/internal/services"
	"github.com/labstack/echo/v4"
)

type UserRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) UserHandler {
	return UserHandler{
		service: service,
	}
}

func (u UserHandler) PostUser(e echo.Context) error {
	userRequest := &UserRequest{}

	if err := e.Bind(userRequest); err != nil {
		return err
	}

	user, err := u.service.Create(userRequest.Name, userRequest.Age)
	if err != nil {
		return err
	}

	_ = e.JSON(http.StatusCreated, UserRequest(user))
	return nil
}

func (u UserHandler) GetALLUsers(e echo.Context) error {
	userRequest := &UserRequest{}

	if err := e.Bind(userRequest); err != nil {
		return err
	}

	users := u.service.GetAll()

	usersRequest := make([]UserRequest, 0)

	for _, user := range users {
		usersRequest = append(usersRequest, UserRequest(user))
	}

	_ = e.JSON(http.StatusCreated, usersRequest)
	return nil
}
