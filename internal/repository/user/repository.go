package user

import (
	model "github.com/tadilbek11kz/ePharma-backend/pkg/user"
)

type Repository interface {
	CreateUser(data model.CreateUserRequest) (user model.User, err error)
	GetUser(key string, value interface{}) (user model.User, err error)
}
