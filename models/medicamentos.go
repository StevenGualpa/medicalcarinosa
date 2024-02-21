// medicamentos.go
package models

import (
	"gorm.io/gorm"
)

type Medicamento struct {
	gorm.Model
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	NumeroDosis int    `json:"numeroDosis"`
	Frecuencia  int    `json:"frecuencia"`
}
