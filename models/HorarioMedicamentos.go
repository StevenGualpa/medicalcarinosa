package models

import (
	"gorm.io/gorm"
	"time"
)

type HorarioMedicamento struct {
	gorm.Model
	PacienteID     uint      `gorm:"not null"`
	MedicamentoID  uint      `gorm:"not null"`
	HoraInicial    time.Time `gorm:"type:timestamp;"`
	HoraProxima    time.Time `gorm:"type:timestamp;"`
	DosisRestantes int
	// Relaciones
	Paciente    Paciente    `gorm:"foreignKey:PacienteID"`
	Medicamento Medicamento `gorm:"foreignKey:MedicamentoID"`
}
