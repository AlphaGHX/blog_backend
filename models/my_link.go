package models

type MyLink struct {
	SiteName string `gorm:"primarykey;not null;unique" json:"site-name"`
	SiteLink string `gorm:"not null" json:"site-link"`
}
