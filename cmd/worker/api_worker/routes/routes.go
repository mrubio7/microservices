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
	r.router.POST("/update/players", handlers.UpdatePlayers)
}
