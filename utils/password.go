package utils 

import (
    "golang.org/x/crypto/bcrypt"
)

type PasswordHash interface{
	HashPassword(string) (string, error)
	GenerateFromPassword(password, hash_password string) error 
}

func HashPassword(password string) (string, error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func GenerateFromPassword(password, hash_password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash_password), []byte(password))
	return err 
}