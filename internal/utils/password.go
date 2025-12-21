package utils

import bcrypt2 "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bcrypt, err := bcrypt2.GenerateFromPassword([]byte(password), bcrypt2.DefaultCost)
	return string(bcrypt), err
}

func CheckPassword(hash string, password string) error {
	return bcrypt2.CompareHashAndPassword([]byte(hash), []byte(password))
}
