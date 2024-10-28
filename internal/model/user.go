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

type UserSessionModel struct {
	UserID    int    `gorm:"primaryKey; not null"`
	SessionID string `gorm:"not null"`
}
