package main

import (
	"ibercs/cmd/worker/api_worker/routes"
	"ibercs/internal/api"
	"ibercs/internal/api/middlewares"
)

func main() {
	worker := api.New(false)

	routes := routes.New(worker.Router)

	routes.RegisterUpdates()

	worker.Router.Use(middlewares.CORSMiddleware())

	worker.Start()
}
