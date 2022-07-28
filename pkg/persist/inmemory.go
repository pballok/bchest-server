package persist

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func newInMemoryTable[KeyType constraints.Ordered, ItemType any](itemName string) *inMemoryTable[KeyType, ItemType] {
	return &inMemoryTable[KeyType, ItemType]{
		name:  itemName,
		items: map[KeyType]ItemType{},
	}
}

type inMemoryTable[KeyType constraints.Ordered, ItemType any] struct {
	name  string
	items map[KeyType]ItemType
}

func (i *inMemoryTable[KeyType, ItemType]) AddNew(key KeyType, item *ItemType) error {
	_, itemAlreadyExists := i.items[key]
	if itemAlreadyExists {
		return fmt.Errorf("Cannot store new %s, because same Key already exists: %v", i.name, key)
	}

	i.items[key] = *item

	return nil
}

func (i *inMemoryTable[KeyType, ItemType]) Find(key KeyType) (ItemType, error) {
	item, exists := i.items[key]
	if !exists {
		var emptyItem ItemType
		return emptyItem, fmt.Errorf("Cannot get %s, because the Key doesn't exists: %v", i.name, key)
	}

	return item, nil
}

func (i *inMemoryTable[KeyType, ItemType]) Purge() {
	i.items = make(map[KeyType]ItemType)
}

func (i *inMemoryTable[KeyType, ItemType]) Count() int {
	return len(i.items)
}

type inMemoryStorageType = allPersistedTables

func (s *inMemoryStorageType) Init() bool {
	// Temporary init code below
	tempPlayer := PlayerData{
		Name:           "pballok",
		HashedPassword: "hash",
	}
	s.players.AddNew(tempPlayer.Name, &tempPlayer)

	return true
}

var inMemoryStorage = inMemoryStorageType{
	players:    newInMemoryTable[string, PlayerData]("Player"),
	characters: newInMemoryTable[string, CharacterData]("Character"),
}
