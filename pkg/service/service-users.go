package service

import (
	"sync"

	"gorm.io/gorm"
)

type Users struct {
	db    *gorm.DB
	mutex sync.Mutex
}

func NewUsersService(database *gorm.DB) *Users {
	return &Users{
		db: database,
	}
}
