package main

import (
	"ibercs/cmd/api_gateway/api"
)

func main() {
	app := api.New()

	app.Start()
}
