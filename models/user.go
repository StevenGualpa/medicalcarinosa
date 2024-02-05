package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func (User) TableName() string {
	return "users"
}
