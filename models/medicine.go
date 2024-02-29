// medicamentos.go
package models

import (
	"gorm.io/gorm"
	"time"
)

type Medicine struct {
	gorm.Model
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
}

type HorarioMedicine struct {
	gorm.Model
	PacienteID     uint `gorm:"not null"`
	MedicamentoID  uint `gorm:"not null"`
	Frecuencia     int  `json:"frecuencia"`
	DosisRestantes int
	HoraInicial    time.Time `gorm:"type:timestamp;"`
	HoraProxima    time.Time `gorm:"type:timestamp;"`
	// Relaciones
	Paciente Paciente `gorm:"foreignKey:PacienteID"`
	Medicine Medicine `gorm:"foreignKey:MedicamentoID"`
}

type HorarioMedicineDetalle struct {
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
