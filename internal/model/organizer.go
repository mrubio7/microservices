package model

type OrganizerModel struct {
	Id       int32  `gorm:"primaryKey; autoIncrement"`
	FaceitId string `gorm:"not null"`
	Name     string `gorm:"not null"`
	Website  string `gorm:"null"`
	Twitter  string `gorm:"null"`
	Twitch   string `gorm:"null"`
	Avatar   string `gorm:"null"`
	Type     string `gorm:"not null"`
}

const (
	ORGANIZER_TYPE_ESEA      = "ESEA"
	ORGANIZER_TYPE_ORGANIZER = "ORGANIZER"
)
