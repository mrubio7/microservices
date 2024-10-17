package model

type PlayerModel struct {
	ID       int32            `gorm:"primaryKey;autoIncrement"`
	Nickname string           `gorm:"not null"`
	FaceitId string           `gorm:"unique;index"`
	SteamId  string           `gorm:"null"`
	Stats    PlayerStatsModel `gorm:"foreignKey:ID;references:ID"`
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
