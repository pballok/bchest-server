package character

type Character struct {
	Name        string `json:"name"`
	Player      string `json:"player"`
	Description string `json:"description"`
}

func NewCharacter(playerName string, name string, description *string) *Character {
	// TODO: Make sure Player exists
	newCharacter := Character{
		Name:   name,
		Player: playerName,
	}

	if description == nil {
		newCharacter.Description = ""
	} else {
		newCharacter.Description = *description
	}

	return &newCharacter
}
