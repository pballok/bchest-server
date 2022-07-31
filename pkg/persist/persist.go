package persist

import (
	"golang.org/x/exp/constraints"
)

type genericStorageTable[KeyType constraints.Ordered, ItemType any] interface {
	AddNew(key KeyType, item *ItemType) error
	Find(key KeyType) (ItemType, error)
	Count() int
}

type playerStorage interface {
	genericStorageTable[string, PlayerData]
}

type characterStorage interface {
	genericStorageTable[string, CharacterData]
}

type storage interface {
	Init() bool
	Players() playerStorage
	Characters() characterStorage
}

var Storage storage = &inMemoryStorage
