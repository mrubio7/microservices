package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type TeamModel struct {
	Id          int32           `gorm:"primaryKey;autoIncrement"`
	FaceitId    string          `gorm:"unique;index"`
	Name        string          `gorm:"not null"`
	Nickname    string          `gorm:"not null"`
	Avatar      string          `gorm:"null"`
	Active      bool            `gorm:"not null; default:false"`
	Stats       TeamStatsModel  `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PlayersId   JSONStringArray `gorm:"type:jsonb;not null"`
	Twitter     string          `gorm:"null"`
	Instagram   string          `gorm:"null"`
	Web         string          `gorm:"null"`
	Tournaments JSONStringArray `gorm:"type:json;null"`
}

// Define a custom type
type JSONStringArray []string

func (j JSONStringArray) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSONStringArray) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSONStringArray value: %v", value)
	}
	return json.Unmarshal(bytes, j)
}

type TeamStatsModel struct {
	ID            int32          `gorm:"primaryKey;autoIncrement:false"`
	TotalMatches  int32          `gorm:"not null"`
	Wins          int32          `gorm:"not null"`
	Winrate       float32        `gorm:"not null"`
	RecentResults JSONInt32Slice `gorm:"type:json;not null"`
	MapStats      JSONMapStats   `gorm:"type:json;not null"`
}

type JSONInt32Slice []int32

func (j JSONInt32Slice) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSONInt32Slice) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSONInt32Slice value: %v", value)
	}
	return json.Unmarshal(bytes, j)
}

type JSONMapStats map[string]TeamMapStats

func (j JSONMapStats) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSONMapStats) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSONMapStats value: %v", value)
	}
	return json.Unmarshal(bytes, j)
}

type TeamMapStats struct {
	MapName string `json:"map_name"`
	WinRate int32  `json:"win_rate"`
	Matches int32  `json:"matches"`
}
