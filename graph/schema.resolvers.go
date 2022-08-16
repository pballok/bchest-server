package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/pballok/bchest-server/graph/generated"
	"github.com/pballok/bchest-server/graph/model"
	"github.com/pballok/bchest-server/internal/auth"
	"github.com/pballok/bchest-server/internal/character"
	"github.com/pballok/bchest-server/internal/persist"
	"github.com/pballok/bchest-server/internal/player"
)

func (r *mutationResolver) CreatePlayer(ctx context.Context, input model.PlayerInput) (*model.Player, error) {
	player, err := player.NewPlayer(input.Name, input.Password)
	if err != nil {
		return nil, err
	}
	err = persist.Storage.Players().AddNew(input.Name, &player.PlayerData)
	if err != nil {
		return nil, err
	}

	return player.GetModel(), nil
}

func (r *mutationResolver) CreateCharacter(ctx context.Context, input model.CharacterInput) (*model.Character, error) {
	currentPlayer, ok := auth.GetPlayerFromContext(ctx)
	if !ok || currentPlayer.Name == "" {
		return nil, fmt.Errorf("Unauthorized request! %v", currentPlayer)
	}
	newCharacter, err := character.NewCharacter(input.Name, &currentPlayer.Name, input.Description)
	if err != nil {
		return nil, err
	}
	err = persist.Storage.Characters().AddNew(input.Name, &newCharacter.CharacterData)
	if err != nil {
		return nil, err
	}
	return newCharacter.GetModel(), nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetCharacter(ctx context.Context, name string) (*model.Character, error) {
	characterData, err := persist.Storage.Characters().Find(name)
	if err != nil {
		return nil, fmt.Errorf("Character not found.")
	}
	return character.FromData(&characterData).GetModel(), nil
}

func (r *queryResolver) ListCharacters(ctx context.Context, player string) ([]*model.Character, error) {
	characterModels := []*model.Character{}
	characters := persist.Storage.Characters().ListByPlayer(player)
	for _, c := range characters {
		characterModels = append(characterModels, character.FromData(&c).GetModel())
	}
	return characterModels, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
