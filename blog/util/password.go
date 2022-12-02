package util

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func PasswordHash(pwd string) (pwdhash string, err error) {
	passwordhash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		err = fmt.Errorf("password hash failed err:%v", err)
		if err != nil {
			return
		}
		return
	}
	pwdhash = string(passwordhash)
	return
}
