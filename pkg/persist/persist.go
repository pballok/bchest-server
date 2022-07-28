package persist

import (
	"golang.org/x/exp/constraints"
)

type persistedTable[KeyType constraints.Ordered, ItemType any] interface {
	AddNew(key KeyType, item *ItemType) error
	Find(key KeyType) (ItemType, error)
	Purge()
	Count() int
}

type allPersistedTables struct {
	players    persistedTable[string, PlayerData]
	characters persistedTable[string, CharacterData]
}

func (s *allPersistedTables) Purge() {
	s.players.Purge()
}

func (s *allPersistedTables) Players() persistedTable[string, PlayerData] {
	return s.players
}

func (s *allPersistedTables) Characters() persistedTable[string, CharacterData] {
	return s.characters
}

type storage interface {
	Init() bool
	Purge()
	Players() persistedTable[string, PlayerData]
	Characters() persistedTable[string, CharacterData]
}

var Storage storage = &inMemoryStorage
