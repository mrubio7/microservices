package faceit

import (
	"fmt"
	"ibercs/internal/model"
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
