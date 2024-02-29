// medicamentos.go
package models

import (
	"gorm.io/gorm"
)

type Medicamento struct {
	gorm.Model
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
}
