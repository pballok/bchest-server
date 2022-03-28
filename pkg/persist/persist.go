package persist

import (
	"errors"

	"github.com/pballok/bchest-server/internal/player"
)

var Players = players{
	allPlayers: playerList{},
}

type playerList map[string]player.Player

type players struct {
	allPlayers playerList
}

func (players *players) AddNew(name string, password string) error {
	_, playerAlreadyExists := players.allPlayers[name]
	if playerAlreadyExists {
		return errors.New("Cannot create new Player, Player Name already exists.")
	}

	hashedPassword, err := player.HashPassword(password)
	if err != nil {
		return err
	}

	players.allPlayers[name] = player.Player{
		Name:           name,
		HashedPassword: hashedPassword,
	}

	return nil
}
