package services

import (
	"context"
	"fmt"
	"log/slog"

	"github.com-test/internal/repository"
	"github.com-test/internal/services/models"
)

type UserService interface {
	Create(name string, age int) (models.User, error)
	Delete(userId int) error
	GetAll() models.Users
	Update(name string, age int) error
}

type userService struct {
	store  repository.UserMemory
	logger slog.Logger
}

func NewUserService(store repository.UserMemory, logger slog.Logger) UserService {
	return &userService{
		store:  store,
		logger: logger,
	}
}

func (u userService) Create(name string, age int) (models.User, error) {
	if age <= 0 {
		u.logger.ErrorContext(context.Background(), fmt.Sprintf("invalid_user_%s", name))

		return models.User{}, Error{
			Msg:  "should greater to 0",
			Code: "inter-error",
		}
	}

	return u.store.Create(name, age)
}

func (u *userService) Delete(userId int) error {
	return u.store.Delete(userId)
}

func (u *userService) GetAll() models.Users {
	return u.store.GetAll()
}

func (u *userService) Update(name string, age int) error {
	return nil
}
