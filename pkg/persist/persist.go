package persist

import (
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
	players persistedTable[string, player.Player]
}

func (s *allPersistedTables) Players() persistedTable[string, player.Player] {
	return s.players
}

func (s *allPersistedTables) Purge() {
	s.players.Purge()
}

type storage interface {
	Init() bool
	Purge()
	Players() persistedTable[string, player.Player]
}

var Storage storage = &inMemoryStorage
