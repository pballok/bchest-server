package player

import (
	"fmt"

	"github.com/pballok/bchest-server/graph/model"
	"github.com/pballok/bchest-server/pkg/persist"
	"golang.org/x/crypto/bcrypt"
)

type Player struct {
	persist.PlayerData
}

func NewPlayer(name string, password string) (*Player, error) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("Cannot create new Player: %w", err)
	}
	return &Player{
		PlayerData: persist.PlayerData{
			Name:           name,
			HashedPassword: hashedPassword,
		},
	}, nil
}

func WithData(data *persist.PlayerData) *Player {
	return &Player{
		PlayerData: *data,
	}
}

func (p *Player) GetModel() *model.Player {
	return &model.Player{
		Name: p.Name,
	}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
