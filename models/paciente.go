package models

import "gorm.io/gorm"

type Paciente struct {
	gorm.Model
	UserID           uint
	NumeroEmergencia string
}
