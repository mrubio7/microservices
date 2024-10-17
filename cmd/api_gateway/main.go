package main

import (
	"ibercs/internal/api"
	"ibercs/pkg/logger"
)

func main() {
	logger.Initialize()

	app := api.New()

	app.Start()
}
