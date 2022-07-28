package auth

import (
	"context"
	"net/http"

	"github.com/pballok/bchest-server/pkg/persist"
	"github.com/pballok/bchest-server/pkg/player"
)

type contextKeyType string

const contextPlayerKey contextKeyType = "player"

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// There is no real authentication for now, just hardcode a logged in Player
		player, _ := persist.Storage.Players().Find("pballok")
		ctx := context.WithValue(r.Context(), contextPlayerKey, player)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func GetPlayerFromContext(ctx context.Context) (player.Player, bool) {
	pl, ok := ctx.Value(contextPlayerKey).(player.Player)
	if !ok {
		return player.Player{}, ok
	}
	return pl, true
}
