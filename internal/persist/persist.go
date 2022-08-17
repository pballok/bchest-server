package persist

import (
	"github.com/pballok/bchest-server/internal/persist/inmemory"
	"github.com/pballok/bchest-server/internal/persist/interfaces"
)

var Storage interfaces.Storage = &inmemory.InMemoryStorage
