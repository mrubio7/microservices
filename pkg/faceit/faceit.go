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
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/mconnat/go-faceit/pkg/client"
	"github.com/mconnat/go-faceit/pkg/models"
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

func (c *FaceitClient) GetPlayerDetails(playerId string) *models.Player {
	p, err := c.client.GetPlayerByID(playerId, nil)
	if err != nil {
		return nil
	}

	return &p
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
		Web:       team.Website,
		Twitter:   team.Twitter,
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
		Type:     model.ORGANIZER_TYPE_ORGANIZER,
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
		Avatar:          champ.Avatar,
		Status:          champ.Avatar,
		JoinPolicy:      champ.JoinChecks.JoinPolicy,
		GeoCountries:    champ.JoinChecks.WhitelistGeoCountries,
		MinLevel:        champ.JoinChecks.MinSkillLevel,
		MaxLevel:        champ.JoinChecks.MaxSkillLevel,
	}
}

func (c *FaceitClient) GetAllChampionshipFromOrganizer(organizerId string) []model.TournamentModel {
	var tournaments []model.TournamentModel
	var offset int = 0
	var limit int = 40

	for {
		params := map[string]interface{}{
			"offset":  strconv.Itoa(offset),
			"limit":   strconv.Itoa(limit),
			"country": "es",
		}

		champs, err := c.client.GetOrganizerChampionships(organizerId, params)
		if err != nil {
			return nil
		}

		if len(champs.Items) == 0 {
			break
		}

		for _, c := range champs.Items {
			tournaments = append(tournaments, model.TournamentModel{
				FaceitId:        c.ChampionshipId,
				OrganizerId:     c.OrganizerId,
				BackgroundImage: c.BackgroundImage,
				Name:            c.Name,
				CoverImage:      c.CoverImage,
				RegisterDate:    time.UnixMilli(int64(c.SubscriptionEnd)),
				StartDate:       time.UnixMilli(int64(c.ChampionshipStart)),
				Avatar:          c.Avatar,
				Status:          c.Status,
				JoinPolicy:      c.JoinChecks.JoinPolicy,
				GeoCountries:    c.JoinChecks.WhitelistGeoCountries,
				MinLevel:        c.JoinChecks.MinSkillLevel,
				MaxLevel:        c.JoinChecks.MaxSkillLevel,
			})
		}

		offset += limit
	}

	return tournaments
}

func (c *FaceitClient) GetTeamsInTournament(tournamentId string) []model.TeamModel {
	var res []model.TeamModel

	var offset int = 0
	var limit int = 10

	for {
		params := map[string]interface{}{
			"offset": strconv.Itoa(offset),
			"limit":  strconv.Itoa(limit),
		}

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
		if len(teams.Items) == 0 {
			break
		}
	}

	return res
}

func (c *FaceitClient) GetESEASeasons_PRODUCTION() []model.EseaLeagueModel {
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

	tournaments := make([]model.EseaLeagueModel, 2)

	payload := rb["payload"].(map[string]any)

	currentSeason := payload["current_season"].(map[string]any)
	tournaments[0] = *convertToEseaModel(currentSeason)

	nextSeason := payload["next_season"].(map[string]any)
	tournaments[1] = *convertToEseaModel(nextSeason)

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
					isPlayoffs := strings.Contains(stage.Name, "Playoffs")
					eseaDivisions = append(eseaDivisions, model.EseaDivisionModel{
						ConferenceId:       stage.Conferences[0].ID,
						TournamentId:       stage.Conferences[0].ChampionshipID,
						EseaLeagueFaceitId: seasonId,
						DivisionName:       division.Name,
						StageName:          stage.Name,
						Playoffs:           isPlayoffs,
					})
				}
			}
		} else {
			continue
		}
	}

	return eseaDivisions
}

func (c *FaceitClient) GetESEADivisionStanding_PRODUCTION(seasonId string) []model.EseaStandingModel {
	var standings []model.EseaStandingModel
	var offset int = 0
	var limit int = 100

	for {
		res, err := http.Get(fmt.Sprintf("https://www.faceit.com/api/team-leagues/v2/standings?entityId=%s&entityType=conference&offset=%d&limit=%d", seasonId, offset, limit))
		if err != nil {
			logger.Error("Error getting ESEA Seasons: %s", err.Error())
			return nil
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			logger.Error("Error reading response body: %s", err.Error())
			return nil
		}

		var rb map[string]interface{}
		if err := json.Unmarshal(body, &rb); err != nil {
			logger.Error("Error unmarshaling response body: %s", err.Error())
			return nil
		}

		// Acceder al payload
		payload, ok := rb["payload"].(map[string]interface{})
		if !ok {
			logger.Error("Error payload is not a map")
			return nil
		}

		// Acceder a standings como un arreglo
		standingsData, ok := payload["standings"].([]interface{})
		if !ok {
			logger.Error("Error standings is not an array")
			return nil
		}

		for _, s := range standingsData {
			standing, ok := s.(map[string]interface{})
			if !ok {
				logger.Error("Error standing is not a map")
				continue
			}

			// Manejar posibles valores nulos o tipos incorrectos
			teamFaceitID, _ := standing["premade_team_id"].(string)
			points, _ := standing["points"].(float64)
			matchesPlayed, _ := standing["matches"].(float64)
			matchesWon, _ := standing["won"].(float64)
			matchesLost, _ := standing["lost"].(float64)
			matchesTied, _ := standing["tied"].(float64)
			buchholzScore, _ := standing["buchholz_score"].(float64)
			rankStart, _ := standing["rank_start"].(float64)
			rankEnd, _ := standing["rank_end"].(float64)
			isDisqualified, _ := standing["is_disqualified"].(bool)

			standings = append(standings, model.EseaStandingModel{
				TeamFaceitId:   teamFaceitID,
				Points:         int(points),
				MatchesPlayed:  int(matchesPlayed),
				MatchesWon:     int(matchesWon),
				MatchesLost:    int(matchesLost),
				MatchesTied:    int(matchesTied),
				BuchholzScore:  int(buchholzScore),
				DivisionId:     seasonId,
				RankStart:      int(rankStart),
				RankEnd:        int(rankEnd),
				IsDisqualified: isDisqualified,
			})
		}

		offset += limit
		if len(standingsData) == 0 {
			break
		}
	}

	return standings
}

func (c *FaceitClient) GetMatchesFromTournamentId(faceitId string, tournamentName string) []model.MatchModel {
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
			var datetime time.Time

			if m.Status == "CANCELLED" || m.Status == "FINISHED" {
				datetime = time.Unix(int64(m.FinishedAt), 0)
			} else if strings.Contains(tournamentName, "ESEA") {
				datetime = time.Unix(int64(m.ScheduledAt), 0)
			} else {
				datetime = time.Unix(int64(m.ScheduledAt), 0)
			}

			var maps []string
			if len(m.Voting.Map.Pick) > 3 {
				maps = []string{"undefined"}
			} else {
				maps = m.Voting.Map.Pick
			}

			res = append(res, model.MatchModel{
				FaceitId:           m.MatchId,
				TournamentFaceitId: faceitId,
				TournamentName:     tournamentName,
				TeamAFaceitId:      m.Teams["faction1"].FactionId,
				TeamAName:          m.Teams["faction1"].Name,
				TeamBFaceitId:      m.Teams["faction2"].FactionId,
				TeamBName:          m.Teams["faction2"].Name,
				BestOf:             int32(m.BestOf),
				Timestamp:          datetime,
				Map:                maps,
				Demos:              m.DemoUrl,
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

func (c *FaceitClient) GetMatchDetails(faceitId string) *model.MatchModel {
	m, err := c.client.GetMatchByID(faceitId, nil)
	if err != nil {
		return nil
	}

	var datetime time.Time

	if strings.Contains(m.CompetitionName, "ESEA") {
		datetime = time.Unix(int64(m.ScheduledAt), 0)
	} else {
		datetime = time.Unix(int64(m.StartedAt), 0)
	}

	var maps []string
	if len(m.Voting.Map.Pick) > 3 {
		maps = []string{"undefined"}
	} else {
		maps = m.Voting.Map.Pick
	}

	return &model.MatchModel{
		FaceitId:           m.MatchId,
		TournamentFaceitId: faceitId,
		TournamentName:     m.CompetitionName,
		TeamAFaceitId:      m.Teams["faction1"].FactionId,
		TeamAName:          m.Teams["faction1"].Name,
		TeamBFaceitId:      m.Teams["faction2"].FactionId,
		TeamBName:          m.Teams["faction2"].Name,
		BestOf:             int32(m.BestOf),
		Timestamp:          datetime,
		Map:                maps,
		Demos:              m.DemoUrl,
		Status:             m.Status,
		ScoreTeamA:         int32(m.Results.Score["faction1"]),
		ScoreTeamB:         int32(m.Results.Score["faction2"]),
	}
}

func (c *FaceitClient) GetMatchDetailsComplete(faceitId string) *models.Match {
	m, err := c.client.GetMatchByID(faceitId, nil)
	if err != nil {
		return nil
	}

	return &m
}

func convertToEseaModel(season map[string]any) *model.EseaLeagueModel {
	startDate, err := time.Parse("2006-01-02T15:04:05Z", season["time_start"].(string))
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
	case time.Now().Before(endDate) && time.Now().After(startDate):
		status = "live"
	case time.Now().Before(endDate):
		status = "join"
	}

	re := regexp.MustCompile(`\d+`)
	matches := re.FindStringSubmatch(season["season_name"].(string))
	if len(matches) == 0 {
		return nil
	}

	seasonNumber, err := strconv.Atoi(matches[0])
	if err != nil {
		return nil
	}

	return &model.EseaLeagueModel{
		FaceitId:  season["season_id"].(string),
		Season:    int32(seasonNumber),
		Name:      season["season_name"].(string),
		StartDate: startDate,
		Status:    status,
	}
}
