package faceit

type TournamentResponse struct {
	Payload payload `json:"payload"`
}

type payload struct {
	SeasonID string   `json:"season_id"`
	Regions  []region `json:"regions"`
}

type region struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Divisions []division `json:"divisions"`
}

type division struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Stages []stage `json:"stages"`
}

type stage struct {
	ID               string       `json:"id"`
	Name             string       `json:"name"`
	Conferences      []conference `json:"conferences"`
	RegistrationMode string       `json:"registration_mode"`
}

type conference struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	ChampionshipID string `json:"championship_id"`
}
