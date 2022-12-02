package hash

import (
	"golang.org/x/crypto/bcrypt"
	"test/Bcrypt/bcrypt01/errorcheck"
)

func PasswordHash(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), 14)
	errorcheck.Check(err)
	return string(bytes), err
}
