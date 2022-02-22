package models

type User struct {
	Username string `gorm:"primarykey;not null;unique" json:"username"`
	Password string `gorm:"not null" json:"password"`
}
