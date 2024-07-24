package passwordHasher

import "golang.org/x/crypto/bcrypt"

type BcryptHasher struct{}

func (b BcryptHasher) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (b BcryptHasher) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func NewBcryptHasher() PasswordHasher {
	return BcryptHasher{}
}
