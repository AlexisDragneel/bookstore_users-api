package services

import (
	"errors"
	"github.com/alexisdragneel/bookstore_users-api/domain/users"
)

func FindUser(id uint64) (*users.User, error) {
	if id == 0 {
		return nil, errors.New("id couldn't be 0")
	}
	user := &users.User{
		ID: id,
	}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(user users.User) (*users.User, error) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}