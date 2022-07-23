package persist

import (
	"github.com/pballok/bchest-server/internal/player"
	"golang.org/x/exp/constraints"
)

type storageItems[KeyType constraints.Ordered, ItemType any] interface {
	AddNewItem(key KeyType, item *ItemType) error
	FindItem(key KeyType) (ItemType, error)
	Purge()
	Count() int
}

type storage struct {
	Players storageItems[string, player.Player]
}

func (s *storage) Purge() {
	s.Players.Purge()
}

var InMemoryStorage = storage{
	Players: NewInMemoryItemList[string, player.Player]("Player"),
}
