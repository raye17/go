package hash

import (
	"bcrypt/errorcheck"
	"golang.org/x/crypto/bcrypt"
)

func PasswordHash(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), 14)
	errorcheck.Check(err)
	return string(bytes), err
}
