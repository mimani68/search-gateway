package dto

type Article struct {
	Id              string `gorm:"type:uuid;primary_key;unique_index;"`
	Title           string `gorm:"null"`
	SubTitle        string `gorm:"null"`
	Language        string `gorm:"null"`
	Source          string `gorm:"null"`
	UniqeIdInSource string `gorm:"null"`
	TypeOfRecord    string `gorm:"null"`
	Section         string `gorm:"null"`
	Summery         string `gorm:"null"`
	Image           string `gorm:"null"`
	Attachments     string `gorm:"null"`
	ReleaseDate     string `gorm:"null"`
	CaptureDate     string `gorm:"null"`
	Link            string `gorm:"null"`
	PermanentLink   string `gorm:"null"`
	Author          string `gorm:"null"`
	Body            string `gorm:"null"`
	Meta            string `gorm:"null"`
}
