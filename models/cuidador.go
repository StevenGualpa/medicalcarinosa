package models

import "gorm.io/gorm"

type Cuidador struct {
	gorm.Model
	UserID   uint   `gorm:"unique"`
	Relacion string `json:"relacion"` // Ej: "hijo", "sobrino", etc.
}

func (Cuidador) TableName() string {
	return "cuidadores"
}
