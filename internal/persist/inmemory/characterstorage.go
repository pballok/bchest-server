package inmemory

import (
	"github.com/pballok/bchest-server/internal/persist/datatypes"
)

type characterStorage struct {
	table[string, datatypes.CharacterData]
}

func (s characterStorage) ListByPlayer(playerName string) []datatypes.CharacterData {
	characters := []datatypes.CharacterData{}
	for _, c := range s.table.items {
		if c.PlayerName == playerName {
			characters = append(characters, c)
		}
	}

	return characters
}
