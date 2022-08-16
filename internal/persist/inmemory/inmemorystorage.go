package inmemory

import (
	"github.com/pballok/bchest-server/internal/persist/datatypes"
	"github.com/pballok/bchest-server/internal/persist/interfaces"
)

type storageType struct {
	players    playerStorage
	characters characterStorage
}

func (s *storageType) Init() bool {
	// Temporary init code below
	tempPlayer := datatypes.PlayerData{
		Name:           "pballok",
		HashedPassword: "hash",
	}
	s.players.AddNew(tempPlayer.Name, &tempPlayer)

	return true
}

func (s *storageType) Players() interfaces.PlayerStorage {
	return s.players
}

func (s *storageType) Characters() interfaces.CharacterStorage {
	return s.characters
}

var InMemoryStorage = storageType{
	players: playerStorage{
		table: table[string, datatypes.PlayerData]{
			name: "Players",
		},
	},
	characters: characterStorage{
		table: table[string, datatypes.CharacterData]{
			name: "Characters",
		},
	},
}
