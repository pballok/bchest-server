package persist

import "github.com/pballok/bchest-server/internal/player"

var Players = map[string]player.Player{}

func SavePlayers() error {
	return nil
}
