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

type HorarioMedicamentoDetalle struct {
	ID             uint      `json:"id"`
	PacienteID     uint      `json:"paciente_id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	MedicamentoID  uint      `json:"medicamento_id"`
	Nombre         string    `json:"nombre"`
	Descripcion    string    `json:"descripcion"`
	NumeroDosis    int       `json:"numero_dosis"`
	Frecuencia     int       `json:"frecuencia"`
	HoraInicial    time.Time `json:"hora_inicial"`
	HoraProxima    time.Time `json:"hora_proxima"`
	DosisRestantes int       `json:"dosis_restantes"`
}
