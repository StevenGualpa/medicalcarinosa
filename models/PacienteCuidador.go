package models

import "gorm.io/gorm"

// PacienteCuidador representa la relación entre un paciente y su cuidador.
type PacienteCuidador struct {
	gorm.Model
	PacienteID uint     `gorm:"not null"`
	CuidadorID uint     `gorm:"not null"`
	Paciente   Paciente `gorm:"foreignKey:PacienteID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Cuidador   Cuidador `gorm:"foreignKey:CuidadorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
