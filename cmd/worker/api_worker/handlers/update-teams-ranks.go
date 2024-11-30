package handlers

import (
	"ibercs/internal/model"
	"ibercs/pkg/config"
	"ibercs/pkg/database"
	"ibercs/pkg/faceit"
	"ibercs/pkg/logger"
	"ibercs/pkg/managers"
	"ibercs/pkg/response"
	"math"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mconnat/go-faceit/pkg/models"
)

func UpdateTeamsRanks(c *gin.Context) {
	cfg, err := config.LoadWorker()
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, response.BuildError("Unable to load config"))
		return
	}

	faceitClient := faceit.New(cfg.ThirdPartyApiTokens.FaceitApiToken)

	stateDatabase := database.NewDatabase(cfg.StateDb)
	stateManager := managers.NewStateManager(stateDatabase.GetDB())

	teamDatabase := database.NewDatabase(cfg.TeamsDb)
	teamManager := managers.NewTeamManager(teamDatabase.GetDB())

	MatchesDatabase := database.NewDatabase(cfg.MatchesDb)
	matchManager := managers.NewMatchManager(MatchesDatabase.GetDB())

	teamsMap := buildTeamsMap(cfg.TeamsDb)

	err = workerTeamsRanksUpdate(teamManager, matchManager, faceitClient, teamsMap)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, response.BuildError("Error while update teams ranks"))
		return
	}

	if err := stateManager.Update_TeamsRanksLastUpdate(); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, response.BuildError("Error updating TeamsRanksLastUpdate state"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("Teams updated", nil))
}

func workerTeamsRanksUpdate(teamManager *managers.TeamManager, matchManager *managers.MatchManager, faceitClient *faceit.FaceitClient, teamsMap map[string]bool) error {
	teams, err := teamManager.GetActiveTeams()
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	for _, team := range teams {
		var teamRank *model.TeamRankModel
		teamRank, err = teamManager.GetTeamRankByFaceitId(team.FaceitId)
		if err != nil || teamRank == nil {
			// Si no existe, inicializamos el ranking
			teamRank = &model.TeamRankModel{
				FaceitId:     team.FaceitId,
				ActualPoints: 100,
				OldPoints:    100,
				Matches:      0,
			}
			teamRank, err = teamManager.CreateTeamRank(teamRank)
			if err != nil {
				logger.Error(err.Error())
				return err
			}
		} else {
			// Si existe, actualizamos los puntos antiguos
			teamRank.OldPoints = teamRank.ActualPoints
		}

		// Obtener los partidos del equipo
		matches, err := matchManager.GetMatchesByTeamId(team.FaceitId)
		if err != nil || len(matches) == 0 {
			logger.Warning("Team %s has no matches", team.Name)
			continue
		}

		teamRank.Matches = 0

		for _, match := range matches {
			if match.PointsMatchTeamA == 0 && match.PointsMatchTeamB == 0 { // Si aún no se calcularon los puntos
				matchDetails := faceitClient.GetMatchDetailsComplete(match.FaceitId)
				if matchDetails == nil {
					logger.Warning("Match details not found for match ID %s", match.FaceitId)
					continue
				}

				// Extraer la liga y calcular el multiplicador
				league := extractLeague(matchDetails.CompetitionName)
				multiplier := leagueMultipliers[league]
				if multiplier == 0 {
					multiplier = 1.0
				}

				if teamRank.LeaguePoints < 100*multiplier {
					teamRank.LeaguePoints = 100 * multiplier
				}

				// Calcular el Elo promedio de ambos equipos
				rosterA := mapToPlayers(matchDetails.Teams["faction1"].Roster)
				rosterB := mapToPlayers(matchDetails.Teams["faction2"].Roster)
				if len(rosterA) == 0 || len(rosterB) == 0 {
					continue
				}
				averageEloTeamA := calculateAverageElo(rosterA, faceitClient)
				averageEloTeamB := calculateAverageElo(rosterB, faceitClient)
				winnerIsTeamA := match.ScoreTeamA > match.ScoreTeamB

				// Calcular la diferencia de Elo
				eloDifference := float32(math.Abs(float64(averageEloTeamA - averageEloTeamB)))
				pointAdjustment := eloDifference / 100.0 // 1 punto extra/menos por cada 100 de diferencia

				// Puntos base
				basePointsWin := float32(10)
				basePointsLoss := float32(10)

				if winnerIsTeamA {
					if averageEloTeamA > averageEloTeamB {
						// Equipo A tiene más Elo y gana
						match.PointsMatchTeamA = basePointsWin - pointAdjustment
						match.PointsMatchTeamB = -(basePointsLoss - pointAdjustment)
					} else {
						// Equipo A tiene menos Elo y gana
						match.PointsMatchTeamA = basePointsWin + pointAdjustment
						match.PointsMatchTeamB = -(basePointsLoss + pointAdjustment)
					}
				} else {
					if averageEloTeamB > averageEloTeamA {
						// Equipo B tiene más Elo y gana
						match.PointsMatchTeamA = -(basePointsLoss - pointAdjustment)
						match.PointsMatchTeamB = basePointsWin - pointAdjustment
					} else {
						// Equipo B tiene menos Elo y gana
						match.PointsMatchTeamA = -(basePointsLoss + pointAdjustment)
						match.PointsMatchTeamB = basePointsWin + pointAdjustment
					}
				}

				// Limitar valores para evitar inconsistencias
				if match.PointsMatchTeamA < -10 {
					match.PointsMatchTeamA = -10
				}
				if match.PointsMatchTeamB < -10 {
					match.PointsMatchTeamB = -10
				}

				err = matchManager.Update(&match)
				if err != nil {
					logger.Error("Cannot update match points for match %s: %v", match.FaceitId, err)
					return err
				}
			}

			// Actualizar puntos del ranking
			if match.TeamAFaceitId == team.FaceitId {
				teamRank.ActualPoints += match.PointsMatchTeamA
			} else if match.TeamBFaceitId == team.FaceitId {
				teamRank.ActualPoints += match.PointsMatchTeamB
			}

			teamRank.Matches++
			logger.Info("Match %s processed. TotalPoints: %.2f, Matches: %d", match.FaceitId, teamRank.ActualPoints, teamRank.Matches)

			// Limitar a los últimos 12 partidos
			if teamRank.Matches == 12 {
				break
			}
		}

		// Guardar el ranking actualizado
		logger.Info("TeamRank: %0.2f", teamRank.LeaguePoints+teamRank.ActualPoints)
		if err := teamManager.UpdateTeamRank(teamRank); err != nil {
			logger.Error("Failed to update rank for team %s: %s", team.FaceitId, err.Error())
			continue
		}
	}

	return nil
}

var leagueMultipliers = map[string]float32{
	"Open":         0.6,
	"Open10":       1.2,
	"Intermediate": 1.8,
	"Main":         2.4,
	"Advanced":     3.2,
}

func extractLeague(competitionName string) string {
	parts := strings.Fields(competitionName) // Divide el nombre por espacios
	if len(parts) > 1 {
		league := parts[3] // Asume que la liga es la tercera última palabra
		return league
	}
	return ""
}

// calculateAverageElo calcula el Elo promedio de un equipo.
func calculateAverageElo(players []models.Player, faceitClient *faceit.FaceitClient) int {
	totalElo := 0
	for _, player := range players {
		details := faceitClient.GetPlayerDetails(player.PlayerId)
		totalElo += details.Games["cs2"].FaceitElo
	}
	return totalElo / len(players)
}

func mapToPlayers(roster []struct {
	AnticheatRequired bool   `json:"anticheat_required"`
	Avatar            string `json:"avatar"`
	GamePlayerId      string `json:"game_player_id"`
	GamePlayerName    string `json:"game_player_name"`
	GameSkillLevel    int    `json:"game_skill_level"`
	Membership        string `json:"membership"`
	Nickname          string `json:"nickname"`
	PlayerId          string `json:"player_id"`
}) []models.Player {
	var players []models.Player
	for _, r := range roster {
		players = append(players, models.Player{
			PlayerId: r.PlayerId,
			Nickname: r.Nickname,
			Avatar:   r.Avatar,
		})
	}
	return players
}
