package paperdb

import (
	"jwt-auth/mocks"
	"jwt-auth/models"
	"jwt-auth/utils"
)

type Mypaperdb struct{}

func (m Mypaperdb) GetUser(username string) (bool, error) {
	_, ok := utils.Struct2Map()[username]
	return ok, nil
}

func (m Mypaperdb) AddUser(user models.User) error {
	mocks.Users = append(mocks.Users, user)
	return nil
}
