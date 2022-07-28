package persist

import (
	"github.com/pballok/bchest-server/pkg/character"
	"github.com/pballok/bchest-server/pkg/player"
	"golang.org/x/exp/constraints"
)

type persistedTable[KeyType constraints.Ordered, ItemType any] interface {
	AddNew(key KeyType, item *ItemType) error
	Find(key KeyType) (ItemType, error)
	Purge()
	Count() int
}

type allPersistedTables struct {
	players    persistedTable[string, player.Player]
	characters persistedTable[string, character.Character]
}

func (s *allPersistedTables) Purge() {
	s.players.Purge()
}

func (s *allPersistedTables) Players() persistedTable[string, player.Player] {
	return s.players
}

func (s *allPersistedTables) Characters() persistedTable[string, character.Character] {
	return s.characters
}

type storage interface {
	Init() bool
	Purge()
	Players() persistedTable[string, player.Player]
	Characters() persistedTable[string, character.Character]
}

var Storage storage = &inMemoryStorage
