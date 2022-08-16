package interfaces

import (
	"github.com/pballok/bchest-server/internal/persist/datatypes"
	"golang.org/x/exp/constraints"
)

type genericStorageTable[KeyType constraints.Ordered, ItemType any] interface {
	AddNew(key KeyType, item *ItemType) error
	Find(key KeyType) (ItemType, error)
	Count() int
}

type PlayerStorage interface {
	genericStorageTable[string, datatypes.PlayerData]
}

type CharacterStorage interface {
	genericStorageTable[string, datatypes.CharacterData]
	ListByPlayer(playerName string) []datatypes.CharacterData
}

type Storage interface {
	Init() bool
	Players() PlayerStorage
	Characters() CharacterStorage
}
