package faceit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"ibercs/internal/model"
	"ibercs/pkg/logger"
	"io"
	"log"
	"net/http"
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
		logger.Error("Error getting player details: %s", err.Error())
		return nil
	}

	params := map[string]interface{}{
		"limit": strconv.Itoa(matchesNumber),
	}
	pstats, err := c.client.GetPlayersLastMatchesStats(userId, "cs2", params)
	if err != nil {
		logger.Error("Error getting last matches: %s", err.Error())
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
		logger.Error("Error getting team stats in faceitservice: %s", err.Error())
		return nil
	}

	var players []string
	for _, t := range team.Members {
		players = append(players, t.UserId)
	}

	mapStats := make(map[string]model.TeamMapStats, len(teamStats.Segments))
	for _, m := range teamStats.Segments {
		mapStats[m.Label] = model.TeamMapStats{
			MapName: m.Label,
			WinRate: int32(m.Stats.WinRatePercent),
			Matches: int32(m.Stats.Matches),
		}
	}

	return &model.TeamModel{
		FaceitId:  team.TeamId,
		Name:      team.Name,
		Nickname:  team.Nickname,
		Avatar:    team.Avatar,
		PlayersId: players,
		Stats: model.TeamStatsModel{
			TotalMatches:  int32(teamStats.Lifetime.Matches),
			Wins:          int32(teamStats.Lifetime.Wins),
			Winrate:       float32(teamStats.Lifetime.WinRatePercent),
			RecentResults: []int32(teamStats.Lifetime.RecentResults),
			MapStats:      mapStats,
		},
	}
}

func (c *FaceitClient) GetOrganizerById(organizerId string) *model.OrganizerModel {
	organizer, err := c.client.GetOrganizerByID(organizerId, nil)
	if err != nil {
		return nil
	}

	return &model.OrganizerModel{
		FaceitId: organizer.OrganizerId,
		Name:     organizer.Name,
		Website:  organizer.Website,
		Twitter:  organizer.Twitter,
		Twitch:   organizer.Twitch,
		Avatar:   organizer.Avatar,
	}
}

func (c *FaceitClient) GetChampionshipById(championshipId string) *model.TournamentModel {
	champ, err := c.client.GetChampionshipByID(championshipId, nil)
	if err != nil {
		return nil
	}

	return &model.TournamentModel{
		FaceitId:        champ.ChampionshipId,
		OrganizerId:     champ.OrganizerId,
		BackgroundImage: champ.BackgroundImage,
		Name:            champ.Name,
		CoverImage:      champ.CoverImage,
		RegisterDate:    time.UnixMilli(int64(champ.SubscriptionEnd)),
		StartDate:       time.UnixMilli(int64(champ.ChampionshipStart)),
		CurrentTeams:    champ.CurrentSubscriptions,
		Slots:           champ.Slots,
		Avatar:          champ.Avatar,
		Status:          champ.Avatar,
		JoinPolicy:      champ.JoinChecks.JoinPolicy,
		GeoCountries:    champ.JoinChecks.WhitelistGeoCountries,
		MinLevel:        champ.JoinChecks.MinSkillLevel,
		MaxLevel:        champ.JoinChecks.MaxSkillLevel,
	}
}

func (c *FaceitClient) GetAllChampionshipFromOrganizer(organizerId string, offset, limit int) []model.TournamentModel {
	params := map[string]interface{}{
		"offset":  strconv.Itoa(offset),
		"limit":   strconv.Itoa(limit),
		"country": "es",
	}

	champs, err := c.client.GetOrganizerChampionships(organizerId, params)
	if err != nil {
		return nil
	}

	var tournaments []model.TournamentModel
	for _, c := range champs.Items {
		tournaments = append(tournaments, model.TournamentModel{
			FaceitId:        c.ChampionshipId,
			OrganizerId:     c.OrganizerId,
			BackgroundImage: c.BackgroundImage,
			Name:            c.Name,
			CoverImage:      c.CoverImage,
			RegisterDate:    time.UnixMilli(int64(c.SubscriptionEnd)),
			StartDate:       time.UnixMilli(int64(c.ChampionshipStart)),
			CurrentTeams:    c.CurrentSubscriptions,
			Slots:           c.Slots,
			Avatar:          c.Avatar,
			Status:          c.Status,
			JoinPolicy:      c.JoinChecks.JoinPolicy,
			GeoCountries:    c.JoinChecks.WhitelistGeoCountries,
			MinLevel:        c.JoinChecks.MinSkillLevel,
			MaxLevel:        c.JoinChecks.MaxSkillLevel,
		})
	}

	return tournaments
}

func (c *FaceitClient) GetTeamsInTournament(tournamentId string, size int) []model.TeamModel {
	var res []model.TeamModel

	var offset int = 0
	var limit int = 50

	params := map[string]interface{}{
		"offset":  strconv.Itoa(offset),
		"limit":   strconv.Itoa(limit),
		"country": "es",
	}

	for offset < size {
		teams, err := c.client.GetSubscriptionsByChampionshipID(tournamentId, params)
		if err != nil {
			return nil
		}

		for _, t := range teams.Items {
			team := model.TeamModel{
				FaceitId:  t.Team.TeamId,
				Name:      t.Team.Name,
				Nickname:  t.Team.Nickname,
				Avatar:    t.Team.Avatar,
				PlayersId: t.Roster,
				Stats:     model.TeamStatsModel{},
			}

			res = append(res, team)
		}

		offset += limit
	}

	return res
}

func (c *FaceitClient) GetESEASeasons_PRODUCTION() []model.TournamentModel {
	// Realiza la solicitud HTTP
	res, err := http.Get("https://www.faceit.com/api/team-leagues/v1/get_league_seasons?league_id=a14b8616-45b9-4581-8637-4dfd0b5f6af8")
	if err != nil {
		log.Printf("Error getting ESEA Seasons: %s", err.Error())
		return nil
	}
	defer res.Body.Close() // Asegura que el cuerpo de la respuesta se cierre al final

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error reading response body: %s", err.Error())
		return nil
	}

	var rb map[string]any
	if err := json.Unmarshal(body, &rb); err != nil {
		log.Printf("Error unmarshaling response body: %s", err.Error())
		return nil
	}

	tournaments := make([]model.TournamentModel, 2)

	payload := rb["payload"].(map[string]any)

	currentSeason := payload["current_season"].(map[string]any)
	tournaments[0] = *convertToTournamentModel(currentSeason)

	nextSeason := payload["next_season"].(map[string]any)
	tournaments[1] = *convertToTournamentModel(nextSeason)

	return tournaments
}

func (c *FaceitClient) GetESEADivisionBySeasonId_PRODUCTION(seasonId string, name string) []model.EseaDivisionModel {
	payloadBody := map[string]string{"seasonId": seasonId}
	jsonData, err := json.Marshal(payloadBody)
	if err != nil {
		log.Printf("Error marshaling JSON: %s", err.Error())
		return nil
	}

	res, err := http.Post("https://www.faceit.com/api/team-leagues/v1/get_filters", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Error getting ESEA Seasons: %s", err.Error())
		return nil
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error reading response body: %s", err.Error())
		return nil
	}

	var rb TournamentResponse
	if err := json.Unmarshal(body, &rb); err != nil {
		log.Printf("Error unmarshaling response body: %s", err.Error())
		return nil
	}

	var eseaDivisions []model.EseaDivisionModel

	for _, region := range rb.Payload.Regions {
		if region.Name == "Europe" {
			for _, division := range region.Divisions {
				for _, stage := range division.Stages {
					eseaDivisions = append(eseaDivisions, model.EseaDivisionModel{
						FaceitId:     stage.Conferences[0].ChampionshipID,
						TournamentId: seasonId,
						Name:         fmt.Sprintf("%s %s - %s", name, division.Name, stage.Name),
					})
				}
			}
		} else {
			continue
		}
	}

	return eseaDivisions
}

func (c *FaceitClient) GetMatchesFromTournamentId(faceitId string) []model.MatchModel {
	var offset int = 0
	var limit int = 100

	var res []model.MatchModel
	for {
		params := map[string]interface{}{
			"offset": strconv.Itoa(offset),
			"limit":  strconv.Itoa(limit),
		}

		matches, err := c.client.GetMatchesByChampionshipID(faceitId, params)
		if err != nil {
			return nil
		}

		for _, m := range matches.Items {
			res = append(res, model.MatchModel{
				FaceitId:           m.MatchId,
				TournamentFaceitId: m.CompetitionId,
				TeamAFaceitId:      m.Teams["faction1"].FactionId,
				TeamAName:          m.Teams["faction1"].Name,
				TeamBFaceitId:      m.Teams["faction2"].FactionId,
				TeamBName:          m.Teams["faction2"].Name,
				BestOf:             int32(m.BestOf),
				Timestamp:          time.UnixMilli(int64(m.StartedAt)),
				Map:                nil,
				Status:             m.Status,
				ScoreTeamA:         int32(m.Results.Score["faction1"]),
				ScoreTeamB:         int32(m.Results.Score["faction2"]),
			})
		}

		offset += limit

		if len(matches.Items) == 0 {
			return res
		}
	}

}

func convertToTournamentModel(season map[string]any) *model.TournamentModel {
	registerDate, err := time.Parse("2006-01-02T15:04:05Z", season["time_start"].(string))
	if err != nil {
		logger.Error("Error parsing register date")
		return nil
	}

	endDate, err := time.Parse("2006-01-02T15:04:05Z", season["time_end"].(string))
	if err != nil {
		logger.Error("Error parsing register date")
		return nil
	}

	var status string
	switch {
	case time.Now().After(endDate):
		status = "finished"
	case time.Now().Before(endDate) && time.Now().After(registerDate):
		status = "live"
	case time.Now().Before(endDate):
		status = "join"
	}

	return &model.TournamentModel{
		OrganizerId:     "a14b8616-45b9-4581-8637-4dfd0b5f6af8",
		FaceitId:        season["season_id"].(string),
		Name:            fmt.Sprintf("ESEA %s", season["season_name"]),
		BackgroundImage: season["header_image_url"].(string),
		CoverImage:      season["thumbnail_url"].(string),
		RegisterDate:    registerDate,
		StartDate:       registerDate,
		MaxLevel:        10,
		MinLevel:        1,
		Status:          status,
		JoinPolicy:      "esea pass",
		GeoCountries:    []string{},
	}
}
