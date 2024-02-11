package models

import "gorm.io/gorm"

type Cuidador struct {
	gorm.Model
	UserID   uint
	Relacion string
}
