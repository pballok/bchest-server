package inmemory

import (
	"testing"

	"github.com/pballok/bchest-server/internal/persist/datatypes"
)

func newTestCharacterStorage() *characterStorage {
	storage := characterStorage{
		table: newTable[string, datatypes.CharacterData]("Characters"),
	}

	storage.AddNew("Character1", &datatypes.CharacterData{
		Name:       "Character1",
		PlayerName: "Player1",
	})
	storage.AddNew("Character2", &datatypes.CharacterData{
		Name:       "Character2",
		PlayerName: "Player2",
	})
	storage.AddNew("Character3", &datatypes.CharacterData{
		Name:       "Character3",
		PlayerName: "",
	})
	storage.AddNew("Character4", &datatypes.CharacterData{
		Name:       "Character4",
		PlayerName: "Player1",
	})
	storage.AddNew("Character5", &datatypes.CharacterData{
		Name:       "Character5",
		PlayerName: "",
	})

	return &storage
}

func TestInMemoryCharacterStorage_ListByPlayerShouldReturnTheCharacters(t *testing.T) {
	testStorage := newTestCharacterStorage()

	filteredCharacters := testStorage.ListByPlayer("Player1")
	want := 2
	got := len(filteredCharacters)
	if want != got {
		t.Fatalf(`Filtered Characters by a Player with %v characters returned a list of %v characters.`, want, got)
	}

	filteredCharacters = testStorage.ListByPlayer("Player2")
	want = 1
	got = len(filteredCharacters)
	if want != got {
		t.Fatalf(`Filtered Characters by a Player with %v characters returned a list of %v characters.`, want, got)
	}

	filteredCharacters = testStorage.ListByPlayer("Player9")
	want = 0
	got = len(filteredCharacters)
	if want != got {
		t.Fatalf(`Filtered Characters by a Player with %v characters returned a list of %v characters.`, want, got)
	}
}
