package repository

import (
	"sync"

	"github.com-test/internal/services/models"
)

type UserMemory interface {
	Create(name string, age int) (models.User, error)
	Delete(userId int) error
	GetAll() models.Users
	Update(id int, name string, age int) error
}

type userMemory struct {
	mx    sync.Mutex
	users models.Users
}

func NewUserMemory() UserMemory {
	return &userMemory{
		users: make(models.Users, 0, 100_000),
	}
}

func (u *userMemory) Create(name string, age int) (models.User, error) {
	u.mx.Lock()
	defer u.mx.Unlock()

	userId := len(u.users) + 1
	user := models.User{
		Id:   userId,
		Name: name,
		Age:  age,
	}

	u.users = append(u.users, user)
	return user, nil
}

func (u *userMemory) Delete(userId int) error {
	return nil
}

func (u *userMemory) GetAll() models.Users {
	u.mx.Lock()
	defer u.mx.Unlock()
	users := make(models.Users, 0, len(u.users))

	for _, user := range u.users {
		letter := rune(user.Name[0])

		switch letter {
		case 'A', 'E', 'I', 'O', 'U':
			users = append(users, user)
		}
	}

	return users
}

func (u *userMemory) Update(id int, name string, age int) error {
	u.mx.Lock()
	defer u.mx.Unlock()

	for i, user := range u.users {
		if user.Id == id {
			u.users[i].Id = id
			u.users[i].Name = name
			u.users[i].Age = age
		}
	}

	return nil
}
