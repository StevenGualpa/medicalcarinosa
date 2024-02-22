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
	Hora               time.Time        `gorm:"type:timestamptz"` // Asegúrate que este tipo sea compatible con tu base de datos
	Estado             string           `gorm:"type:varchar(100);default:'pendiente'"`
	Nombre             string           `gorm:"type:varchar(255)"` // Agregado
	Descripcion        string           `gorm:"type:text"`         // Agregado
	PacienteCuidador   PacienteCuidador `gorm:"foreignKey:PacienteCuidadorID"`
}
