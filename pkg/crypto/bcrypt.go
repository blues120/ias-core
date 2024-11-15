package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

// BcryptMake is a function to generate a bcrypt hash from a password.
func BcryptMake(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// BcryptMakeCheck is a function to check a bcrypt hash from a password.
func BcryptMakeCheck(pwd []byte, hashedPwd string) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, pwd)
	if err != nil {
		return false
	}
	return true
}
