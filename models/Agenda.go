package models

import (
	"gorm.io/gorm"
	"time"
)

// Agenda representa la programación de citas o eventos.
type Agenda struct {
	gorm.Model
	PacienteCuidadorID uint             `gorm:"not null"`                              // Clave foránea de la relación PacienteCuidador.
	Fecha              time.Time        `gorm:"type:date"`                             // La fecha del evento.
	Hora               string           `gorm:"type:time"`                             // La hora del evento.
	Estado             string           `gorm:"type:varchar(100);default:'pendiente'"` // Estado del evento (pendiente, cancelado).
	PacienteCuidador   PacienteCuidador `gorm:"foreignKey:PacienteCuidadorID"`         // Relación con PacienteCuidador.
}
