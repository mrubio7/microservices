package main

import (
	"ibercs/cmd/api_gateway/api_app/routes"
	"ibercs/internal/api"
	"ibercs/internal/api/middlewares"
)

func main() {
	app := api.New(true)

	routes := routes.New(app.Router)

	app.Router.Use(middlewares.CORSMiddleware())
	routes.RegisterUserRoutes()
	routes.RegisterPlayerRoutes()
	routes.RegisterTeamRoutes()
	routes.RegisterTournamentRoutes()
	routes.RegisterMatchRoutes()

	app.Start()
}
