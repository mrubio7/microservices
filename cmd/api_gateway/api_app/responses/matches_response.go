package responses

import "ibercs/internal/model"

type response_Match struct {
	ID           int32  `json:"id"`
	FaceitId     string `json:"faceit_id"`
	TeamAName    string `json:"team_a_name"`
	TeamBName    string `json:"team_b_name"`
	IsTeamAKnown bool   `json:"is_team_a_known"`
	IsTeamBKnown bool   `json:"is_team_b_known"`
	BestOf       int32  `json:"best_of"`
	Map          string `json:"map"`
	TeamA        string `json:"team_a"`
	TeamB        string `json:"team_b"`
	Timestamp    int64  `json:"timestamp"`
}

func Build_MatchResponse(match *model.MatchModel) response_Match {
	return response_Match{
		ID:           int32(match.ID),
		FaceitId:     match.FaceitId,
		TeamAName:    match.TeamAName,
		TeamBName:    match.TeamBName,
		IsTeamAKnown: match.IsTeamAKnown,
		IsTeamBKnown: match.IsTeamBKnown,
		BestOf:       int32(match.BestOf),
		Map:          match.Map[0],
		TeamA:        match.TeamAFaceitId,
		TeamB:        match.TeamBFaceitId,
		Timestamp:    match.Timestamp.Unix(),
	}
}
