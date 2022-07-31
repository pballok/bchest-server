package persist

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type inMemoryTable[KeyType constraints.Ordered, ItemType any] struct {
	name  string
	items map[KeyType]ItemType
}

func (i inMemoryTable[KeyType, ItemType]) AddNew(key KeyType, item *ItemType) error {
	_, itemAlreadyExists := i.items[key]
	if itemAlreadyExists {
		return fmt.Errorf("Cannot store new %s, because same Key already exists: %v", i.name, key)
	}

	i.items[key] = *item

	return nil
}

func (i inMemoryTable[KeyType, ItemType]) Find(key KeyType) (ItemType, error) {
	item, exists := i.items[key]
	if !exists {
		var emptyItem ItemType
		return emptyItem, fmt.Errorf("Cannot get %s, because the Key doesn't exists: %v", i.name, key)
	}

	return item, nil
}

func (i inMemoryTable[KeyType, ItemType]) Count() int {
	return len(i.items)
}
