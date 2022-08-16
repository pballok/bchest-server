package inmemory

import "github.com/pballok/bchest-server/internal/persist/datatypes"

type playerStorage struct {
	table[string, datatypes.PlayerData]
}
