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
	Hora               string           `gorm:"type:time"` // Asegúrate de que este tipo sea compatible con tu SGBD.
	Estado             string           `gorm:"type:varchar(100);default:'pendiente'"`
	PacienteCuidador   PacienteCuidador `gorm:"foreignKey:PacienteCuidadorID"`
}
