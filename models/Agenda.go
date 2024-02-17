package models

import (
	"gorm.io/gorm"
	"time"
)

// Agenda representa la programación de citas o eventos.
type Agenda struct {
	gorm.Model
	PacienteCuidadorID uint             `gorm:"not null"`
	Fecha              time.Time        `gorm:"type:date"`
	Hora               time.Time        `gorm:"type:timestamptz"` // Cambiado a time.Time
	Estado             string           `gorm:"type:varchar(100);default:'pendiente'"`
	PacienteCuidador   PacienteCuidador `gorm:"foreignKey:PacienteCuidadorID"`
}
