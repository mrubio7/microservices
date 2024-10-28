package model

import (
	"database/sql"
)

type StateModel struct {
	ID               int32        `gorm:"primaryKey"`
	LastPlayerUpdate sql.NullTime `gorm:"null"`
}
