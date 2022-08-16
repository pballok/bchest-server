package auth

import (
	"context"
	"net/http"

	"github.com/pballok/bchest-server/internal/persist"
	"github.com/pballok/bchest-server/internal/persist/datatypes"
	"github.com/pballok/bchest-server/internal/player"
)

type contextKeyType string

const contextPlayerKey contextKeyType = "player"

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		playerName := "pballok" // TODO: get this from JWT

		player, err := persist.Storage.Players().Find(playerName)
		if err == nil {
			ctx := context.WithValue(r.Context(), contextPlayerKey, player)
			r = r.WithContext(ctx)
		}
		next.ServeHTTP(w, r)
	})
}

func GetPlayerFromContext(ctx context.Context) (player.Player, bool) {
	playerData, ok := ctx.Value(contextPlayerKey).(datatypes.PlayerData)
	if !ok {
		return player.Player{}, ok
	}
	return *player.FromData(&playerData), true
}
