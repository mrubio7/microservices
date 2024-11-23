package routes

import (
	"ibercs/cmd/worker/api_worker/handlers"
	"ibercs/internal/router"
)

type routes struct {
	router *router.Router
}

func New(router *router.Router) *routes {
	return &routes{
		router: router,
	}
}

func (r routes) RegisterUpdates() {
	r.router.POST("/find/esea", handlers.FindEseaLeague)
	r.router.POST("/find/tournaments", handlers.FindTournaments)
	r.router.POST("/find/matches", handlers.FindMatches)
	r.router.POST("/update/matches/nearby", handlers.UpdateNearbyMatches)
	r.router.POST("/update/players", handlers.UpdatePlayers)
	r.router.POST("/update/tournaments", handlers.UpdateTournaments)
	r.router.POST("/update/teams", handlers.UpdateTeams)
	r.router.POST("/update/esea", handlers.UpdateEsea)
}
