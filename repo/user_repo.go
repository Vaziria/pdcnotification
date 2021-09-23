package repo

import (
	"github.com/vaziria/pdcnotification/models"
)

type UserRepo struct{}

func (repo UserRepo) findByEmail(email string) models.User {

	user := models.User{
		Email: "email@gmail.com",
		Token: []string{"token"},
	}

	return user
}

func (repo UserRepo) createUser(email string, token []string) {

}

func (repo UserRepo) addToken(email string, token []string) {

}
