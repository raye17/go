package encrypt

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(password string) ([]byte, error) {
	return encrypt(password)
}
func encrypt(password string) ([]byte, error) {
	pwd := []byte(password)
	hashedPwd, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hashedPwd, err
}
func Decrypt(hashPwd, password string) error {
	hashedPwd := []byte(hashPwd)
	pwd := []byte(password)
	return decrypt(hashedPwd, pwd)
}
func decrypt(hashed, pwd []byte) error {
	err := bcrypt.CompareHashAndPassword(hashed, pwd)
	if err != nil {
		return err
	}
	fmt.Println("pwd is true")
	return nil
}
