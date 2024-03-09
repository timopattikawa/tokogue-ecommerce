package util

import "golang.org/x/crypto/bcrypt"

type PasswordGenerator interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type PasswordGeneratorImpl struct {
}

func NewPasswordGenerator() PasswordGenerator {
	return &PasswordGeneratorImpl{}
}

func (p PasswordGeneratorImpl) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (p PasswordGeneratorImpl) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
