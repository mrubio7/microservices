package faceit

import (
	"fmt"
	"ibercs/internal/model"
	"ibercs/pkg/logger"
	"strconv"
	"time"

	"github.com/mconnat/go-faceit/pkg/client"
)

type FaceitClient struct {
	client client.FaceITClient
}

func New(token string) *FaceitClient {
	faceit, err := client.New(token)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &FaceitClient{
		client: faceit,
	}
}

func (c *FaceitClient) GetAllPlayers(chPlayers chan model.PlayerModel, size int) {
	defer close(chPlayers)

	var offset int = 0
	var limit int = 100

	for offset < size {
		params := map[string]interface{}{
			"offset":  strconv.Itoa(offset),
			"limit":   strconv.Itoa(limit),
			"country": "es",
		}

		playersList, err := c.client.GetGameRankingByRegion("cs2", "EU", params)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, p := range playersList.Items {
			chPlayers <- model.PlayerModel{
				FaceitId: p.PlayerId,
				Nickname: p.Nickname,
				Stats: model.PlayerStatsModel{
					Elo: int32(p.FaceitElo),
				},
			}
		}

		offset += limit
		time.Sleep(200 * time.Millisecond)
	}
}

func (c *FaceitClient) GetPlayerAverageDetails(userId string, matchesNumber int) *model.PlayerModel {
	p, err := c.client.GetPlayerByID(userId, nil)
	if err != nil {
		return nil
	}

	params := map[string]interface{}{
		"limit": strconv.Itoa(matchesNumber),
	}
	pstats, err := c.client.GetPlayersLastMatchesStats(userId, "cs2", params)
	if err != nil {
		return nil
	}

	var stats model.PlayerStatsModel
	stats.Elo = int32(p.Games["cs2"].FaceitElo)

	for _, s := range pstats.Items {
		kdRatio, _ := strconv.ParseFloat(s.Stats.KDRatio, 32)
		stats.KdRatio += float32(kdRatio)

		krRatio, _ := strconv.ParseFloat(s.Stats.KRRatio, 32)
		stats.KrRatio += float32(krRatio)

		hsAvg, _ := strconv.ParseFloat(s.Stats.HeadshotsPercent, 32)
		stats.HeadshotPercentAverage += float32(hsAvg)

		killsAverage, _ := strconv.ParseFloat(s.Stats.Kills, 32)
		stats.KillsAverage += float32(killsAverage)

		deathsAverage, _ := strconv.ParseFloat(s.Stats.Deaths, 32)
		stats.DeathsAverage += float32(deathsAverage)

		assistAverage, _ := strconv.ParseFloat(s.Stats.Assists, 32)
		stats.AssistAverage += float32(assistAverage)

		mvpAverage, _ := strconv.ParseFloat(s.Stats.MVPs, 32)
		stats.MVPAverage += float32(mvpAverage)

		tripleKillsAvg, _ := strconv.ParseFloat(s.Stats.TripleKills, 32)
		stats.TripleKillsAverage += float32(tripleKillsAvg)

		quadKillsAvg, _ := strconv.ParseFloat(s.Stats.QuadroKills, 32)
		stats.QuadroKillsAverage += float32(quadKillsAvg)

		pentKillsAvg, _ := strconv.ParseFloat(s.Stats.PentaKills, 32)
		stats.PentaKillsAverage += float32(pentKillsAvg)
	}

	return &model.PlayerModel{
		FaceitId: p.PlayerId,
		SteamId:  p.Platforms["steam"],
		Nickname: p.Nickname,
		Avatar:   p.Avatar,
		Stats: model.PlayerStatsModel{
			Elo:                    stats.Elo,
			KdRatio:                stats.KdRatio / float32(matchesNumber),
			KrRatio:                stats.KrRatio / float32(matchesNumber),
			KillsAverage:           stats.KillsAverage / float32(matchesNumber),
			DeathsAverage:          stats.DeathsAverage / float32(matchesNumber),
			AssistAverage:          stats.AssistAverage / float32(matchesNumber),
			HeadshotPercentAverage: stats.HeadshotPercentAverage / float32(matchesNumber),
			MVPAverage:             stats.MVPAverage / float32(matchesNumber),
			TripleKillsAverage:     stats.TripleKillsAverage / float32(matchesNumber),
			QuadroKillsAverage:     stats.QuadroKillsAverage / float32(matchesNumber),
			PentaKillsAverage:      stats.PentaKillsAverage / float32(matchesNumber),
		},
	}
}

func (c *FaceitClient) GetTeamById(teamId string) *model.TeamModel {
	team, err := c.client.GetTeamByID(teamId, nil)
	if err != nil {
		logger.Error("Error getting team in faceitservice: %s", err.Error())
		return nil
	}

	teamStats, err := c.client.GetTeamStats(teamId, "cs2", nil)
	if err != nil {
		logger.Error("Error getting team in faceitservice: %s", err.Error())
		return nil
	}
	var players []string

	for _, t := range team.Members {
		players = append(players, t.UserId)
	}

	mapStats := make(map[string]model.TeamMapStats, len(teamStats["Segments"].([]interface{})))

	for _, m := range teamStats["Segments"].([]interface{}) {
		// Aseguramos que 'm' sea un map[string]interface{}
		mapItem := m.(map[string]interface{})

		// Accedemos al Label y Stats de manera dinámica
		mapLabel := mapItem["Label"].(string)
		mapStatsItem := mapItem["Stats"].(map[string]interface{})

		// Convertimos los valores correspondientes
		winRate := int32(mapStatsItem["WinRatePercent"].(float64)) // asumiendo que es float64
		matches := int32(mapStatsItem["Matches"].(float64))        // asumiendo que es float64

		mapStats[mapLabel] = model.TeamMapStats{
			MapName: mapLabel,
			WinRate: winRate,
			Matches: matches,
		}
	}

	convertInterfaceSliceToInt32Slice := func(slice []interface{}) []int32 {
		result := make([]int32, len(slice))
		for i, v := range slice {
			result[i] = int32(v.(float64)) // Asumiendo que los valores son float64
		}
		return result
	}

	return &model.TeamModel{
		FaceitId:  team.Nickname,
		Name:      team.Name,
		Nickname:  team.Nickname,
		Avatar:    team.Avatar,
		PlayersId: players,
		Stats: model.TeamStatsModel{
			TotalMatches:  int32(teamStats["Lifetime"].(map[string]interface{})["Matches"].(float64)),
			Wins:          int32(teamStats["Lifetime"].(map[string]interface{})["Wins"].(float64)),
			Winrate:       float32(teamStats["Lifetime"].(map[string]interface{})["WinRatePercent"].(float64)),
			RecentResults: convertInterfaceSliceToInt32Slice(teamStats["Lifetime"].(map[string]interface{})["RecentResults"].([]interface{})),
			MapStats:      mapStats,
		},
	}

	// mapStats := make(map[string]model.TeamMapStats, len(teamStats.Segments))
	// for _, m := range teamStats.Segments {
	// 	mapStats[m.Label] = model.TeamMapStats{
	// 		MapName: m.Label,
	// 		WinRate: int32(m.Stats.WinRatePercent),
	// 		Matches: int32(m.Stats.Matches),
	// 	}
	// }

	// return &model.TeamModel{
	// 	FaceitId:  team.TeamId,
	// 	Name:      team.Name,
	// 	Nickname:  team.Nickname,
	// 	Avatar:    team.Avatar,
	// 	PlayersId: players,
	// 	Stats: model.TeamStatsModel{
	// 		TotalMatches:  int32(teamStats.Lifetime.Matches),
	// 		Wins:          int32(teamStats.Lifetime.Wins),
	// 		Winrate:       float32(teamStats.Lifetime.WinRatePercent),
	// 		RecentResults: []int32(teamStats.Lifetime.RecentResults),
	// 		MapStats:      mapStats,
	// 	},
	// }
}
