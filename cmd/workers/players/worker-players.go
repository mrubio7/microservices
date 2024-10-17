package main

import (
	"flag"
	"fmt"
	"ibercs/cmd/workers/players/find_players"
	"ibercs/cmd/workers/players/update_players"
	"ibercs/pkg/logger"
)

const (
	UPDATE = "update"
	FIND   = "find"
)

func main() {
	logger.Initialize()
	mode := flag.String("flag", "", fmt.Sprintf("[%s] (Update all the players stats)\n[%s] (Finds new players in ladderboard)", UPDATE, FIND))
	number := flag.Int("size", 5000, "Number of players to find when use find option")
	flag.Parse()

	switch *mode {
	case UPDATE:
		update_players.Start()
	case FIND:
		find_players.Start(*number)
	default:
		logger.Debug("Unknown flag:", *mode)
	}
}
