package models

import "gorm.io/gorm"

type Paciente struct {
	gorm.Model
	UserID           uint   `gorm:"unique"`
	NumeroEmergencia string `json:"numeroEmergencia"`
}

func (Paciente) TableName() string {
	return "pacientes"
}
