package inmemory

import (
	"testing"

	"github.com/pballok/bchest-server/internal/persist/datatypes"
)

var testStorage = characterStorage{
	table: newInMemoryTable[string, datatypes.CharacterData]("Characters"),
}

func setupTestCharacterStorage() {
	testStorage.AddNew("Character1", &datatypes.CharacterData{
		Name:       "Character1",
		PlayerName: "Player1",
	})
	testStorage.AddNew("Character2", &datatypes.CharacterData{
		Name:       "Character2",
		PlayerName: "Player2",
	})
	testStorage.AddNew("Character3", &datatypes.CharacterData{
		Name:       "Character3",
		PlayerName: "",
	})
	testStorage.AddNew("Character4", &datatypes.CharacterData{
		Name:       "Character4",
		PlayerName: "Player1",
	})
	testStorage.AddNew("Character5", &datatypes.CharacterData{
		Name:       "Character5",
		PlayerName: "",
	})
}

func TestInMemoryCharacterStorage_ListByPlayerShouldReturnTheCharacters(t *testing.T) {
	setupTestCharacterStorage()

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
