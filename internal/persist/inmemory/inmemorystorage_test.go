package inmemory

import (
	"testing"

	"github.com/pballok/bchest-server/internal/persist/datatypes"
)

func newTestStorage() *storageType {
	return &storageType{
		players: playerStorage{
			table: newTable[string, datatypes.PlayerData]("Players"),
		},
		characters: characterStorage{
			table: newTable[string, datatypes.CharacterData]("Characters"),
		},
	}
}

func TestInMemoryStorage_InitReturnsWithTrue(t *testing.T) {
	testStorage := newTestStorage()
	got := testStorage.Init()
	if !got {
		t.Fatalf(`InMemoryStorage Initialization failed.`)
	}
}

func TestInMemoryStorage_GettingThePlayerStorageIsSuccessful(t *testing.T) {
	playerStorage := newTestStorage().Players()
	want := 0
	got := playerStorage.Count()
	if want != got {
		t.Fatalf(`Failed to get PlayerStorage out of InMemoryStorage`)
	}
}

func TestInMemoryStorage_GettingTheCharacterStorageIsSuccessful(t *testing.T) {
	characterStorage := newTestStorage().Characters()
	want := 0
	got := characterStorage.Count()
	if want != got {
		t.Fatalf(`Failed to get CharacterStorage out of InMemoryStorage`)
	}
}
