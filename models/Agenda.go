package models

import (
	"gorm.io/gorm"
	"time"
)

// Agenda representa la programación de citas o eventos.
type Agenda struct {
	gorm.Model
	PacienteID  uint      `gorm:"not null"` // Ahora apunta directamente al ID de paciente
	Fecha       time.Time `gorm:"type:date"`
	Hora        time.Time `gorm:"type:timestamptz"` // Asegúrate que este tipo sea compatible con tu base de datos
	Estado      string    `gorm:"type:varchar(100);default:'pendiente'"`
	Nombre      string    `gorm:"type:varchar(255)"`     // Agregado
	Descripcion string    `gorm:"type:text"`             // Agregado
	Paciente    Paciente  `gorm:"foreignKey:PacienteID"` // Actualizado para referenciar la clave foránea correcta
}

// Define un modelo de detalle para capturar los resultados de tu consulta
type AgendaDetalle struct {
	ID               uint      `json:"id"`
	Nombre           string    `json:"nombre"`
	Descripcion      string    `json:"descripcion"`
	Fecha            time.Time `json:"fecha"`
	Hora             time.Time `json:"hora"`
	Estado           string    `json:"estado"`
	PacienteID       uint      `json:"paciente_id"`
	PacienteNombre   string    `json:"paciente_nombre"`
	PacienteApellido string    `json:"paciente_apellido"`
	// Nota: CuidadorID, CuidadorNombre, y CuidadorApellido han sido removidos
}
