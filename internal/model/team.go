package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type TeamModel struct {
	Id        int32           `gorm:"primaryKey;autoIncrement"`
	FaceitId  string          `gorm:"unique;index"`
	Name      string          `gorm:"not null"`
	Nickname  string          `gorm:"not null"`
	Avatar    string          `gorm:"null"`
	Active    bool            `gorm:"not null; default:true"`
	PlayersId JSONStringArray `gorm:"type:json;not null"`
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
