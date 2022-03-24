package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/pballok/bchest-server/graph/generated"
	"github.com/pballok/bchest-server/graph/model"
)

func (r *mutationResolver) CreatePlayer(ctx context.Context, input model.PlayerInput) (*model.Player, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateCharacter(ctx context.Context, input model.CharacterInput) (*model.Character, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetCharacter(ctx context.Context, name string) (*model.Character, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListCharacters(ctx context.Context, player string) ([]*model.Character, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
