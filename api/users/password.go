package users

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPass(password string) (string, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("Unable to hash password : " + err.Error())
	}
	return string(pass), nil
}
