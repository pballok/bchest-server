package persist

import (
	"testing"
)

func TestPlayerAddNew_PlayerAddedSuccessfully(t *testing.T) {
	Players.Purge()

	playerName := "Player1"
	playerPassword := "Password1"
	err := Players.AddNew(playerName, playerPassword)

	if err != nil {
		t.Fatalf(`Failed to add a single Player: %v`, err)
	}
}

func TestPlayerAddNew_PlayerAddedWithExistingNameReturnsError(t *testing.T) {
	Players.Purge()

	playerName := "Player1"
	playerPassword := "Password1"
	err := Players.AddNew(playerName, playerPassword)

	if err != nil {
		t.Fatalf(`Failed to first single Player: %v`, err)
	}

	err = Players.AddNew(playerName, playerPassword)
	if err == nil {
		t.Fatal(`Adding second player with same name should have returned an error but it did not.`)
	}
}
