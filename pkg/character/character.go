package character

import (
	"fmt"

	"github.com/pballok/bchest-server/graph/model"
	"github.com/pballok/bchest-server/pkg/persist"
	"github.com/pballok/bchest-server/pkg/player"
)

type Character struct {
	persist.CharacterData
}

func NewCharacter(name string, playerName *string, description *string) (*Character, error) {
	newCharacter := Character{
		CharacterData: persist.CharacterData{
			Name: name,
		},
	}

	if playerName == nil {
		newCharacter.PlayerName = ""
	} else {
		_, err := persist.Storage.Players().Find(*playerName)
		if err != nil {
			return &Character{}, fmt.Errorf("Cannot create Character: %w", err)
		}
		newCharacter.PlayerName = *playerName
	}

	if description == nil {
		newCharacter.Description = ""
	} else {
		newCharacter.Description = *description
	}

	return &newCharacter, nil
}

func WithData(data *persist.CharacterData) *Character {
	return &Character{
		CharacterData: *data,
	}
}

func (c *Character) GetModel() *model.Character {
	modelCharacter := model.Character{
		Name:        c.Name,
		Description: &c.Description,
	}
	if c.CharacterData.PlayerName == "" {
		modelCharacter.Player = nil
	} else {
		playerData, err := persist.Storage.Players().Find(c.PlayerName)
		if err != nil {
			return &model.Character{}
		}
		modelCharacter.Player = player.WithData(&playerData).GetModel()
	}
	return &modelCharacter
}
