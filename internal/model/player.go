package model

type PlayerModel struct {
	ID       int32            `gorm:"primaryKey;autoIncrement"`
	Nickname string           `gorm:"not null"`
	FaceitId string           `gorm:"unique;index"`
	SteamId  string           `gorm:"null"`
	Avatar   string           `gorm:"null"`
	Stats    PlayerStatsModel `gorm:"foreignKey:ID;references:ID"`
}

func (PlayerModel) TableName() string {
	return "players.player"
}

type PlayerStatsModel struct {
	ID                     int32   `gorm:"primaryKey;autoIncrement:false"`
	KrRatio                float32 `gorm:"type:numeric(8,2)"`
	KdRatio                float32 `gorm:"type:numeric(8,2)"`
	KillsAverage           float32 `gorm:"type:numeric(8,2)"`
	DeathsAverage          float32 `gorm:"type:numeric(8,2)"`
	HeadshotPercentAverage float32 `gorm:"type:numeric(8,2)"`
	MVPAverage             float32 `gorm:"type:numeric(8,2)"`
	AssistAverage          float32 `gorm:"type:numeric(8,2)"`
	TripleKillsAverage     float32 `gorm:"type:numeric(8,2)"`
	QuadroKillsAverage     float32 `gorm:"type:numeric(8,2)"`
	PentaKillsAverage      float32 `gorm:"type:numeric(8,2)"`
	Elo                    int32
}

func (PlayerStatsModel) TableName() string {
	return "players.player_stats"
}

type PlayerProminentModel struct {
	ID              int32   `gorm:"primaryKey;autoIncrement"`
	Avatar          string  `gorm:"null"`
	Nickname        string  `gorm:"not null"`
	FaceitId        string  `gorm:"not null"`
	SteamId         string  `gorm:"not null"`
	Score           float32 `gorm:"not null"`
	ProminentWeekID int32   `gorm:"index"`
}

func (PlayerProminentModel) TableName() string {
	return "players.prominent_player"
}

type ProminentWeekModel struct {
	ID      int32                  `gorm:"primaryKey;autoIncrement"`
	Week    int16                  `gorm:"not null;uniqueIndex:unique_week_year"`
	Year    int16                  `gorm:"not null;uniqueIndex:unique_week_year"`
	Players []PlayerProminentModel `gorm:"foreignKey:ProminentWeekID;references:ID"`
}

func (ProminentWeekModel) TableName() string {
	return "players.prominent_week"
}
