package model

type UserModel struct {
	ID          int         `gorm:"primaryKey;autoIncrement"`
	FaceitID    string      `gorm:"unique;index"`
	Name        string      `gorm:"not null"`
	Description string      `gorm:"null"`
	Twitter     string      `gorm:"null"`
	Twitch      string      `gorm:"null"`
	Role        int         `gorm:"null"`
	Player      PlayerModel `gorm:"-"`
}
