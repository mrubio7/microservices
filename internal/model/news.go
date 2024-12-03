package model

import "time"

type NewsModel struct {
	Id        int32     `gorm:"primaryKey; autoIncrement"`
	Title     string    `gorm:"not null"`
	Image     string    `gorm:"not null"`
	Content   string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	Published bool      `gorm:"default:false"`
	User      UserModel `gorm:"foreignKey:Id;references:Id"`
}

func (NewsModel) TableName() string {
	return "users.news"
}
