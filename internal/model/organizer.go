package model

type OrganizerModel struct {
	ID       int32  `gorm:"primaryKey; autoIncrement"`
	FaceitId string `gorm:"not null"`
	Name     string `gorm:"not null"`
	Website  string `gorm:"null"`
	Twitter  string `gorm:"null"`
	Twitch   string `gorm:"null"`
	Avatar   string `gorm:"null"`
}
