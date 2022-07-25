package player

import (
	"golang.org/x/crypto/bcrypt"
)

type Player struct {
	Name           string `json:"name"`
	HashedPassword string `json:"password"`
}

func NewPlayer(name string, password string) (*Player, error) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}
	return &Player{
		Name:           name,
		HashedPassword: hashedPassword,
	}, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
