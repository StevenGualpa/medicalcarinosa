package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Email       string `json:"email" gorm:"unique"`
	Password    string `json:"password"`
	BirthDate   string `json:"birthdate"` // Tipo cambiado a string
	Gender      string `json:"gender"`
	PhoneNumber string `json:"phone"`
}

func (User) TableName() string {
	return "users"
}
