package users

import (
	"errors"
	"fmt"
	"time"
)

var (
	usersDb = make(map[uint64]*User)
)

func GetAll() ([]*User, error) {
	return nil, nil
}

func (u *User) Get() error {
	res, ok := usersDb[u.ID]

	if !ok {
		return errors.New(fmt.Sprintf("user %d not found", u.ID))
	}
	u.FirstName = res.FirstName
	u.LastName = res.LastName
	u.Email = res.Email
	u.DateCreated = res.DateCreated
	return nil
}

func (u *User) Save() error {
	u.ID = uint64(len(usersDb) + 1)
	_, ok := usersDb[u.ID]
	if ok {
		return errors.New(fmt.Sprintf("user %v alreade exist", u.Email))
	}
	u.DateCreated = time.Now().String()
	usersDb[u.ID] = u
	return nil
}