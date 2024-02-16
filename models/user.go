package models

import (
	"gorm.io/gorm"
)

// User representa la tabla de usuarios con información básica.
type User struct {
	gorm.Model
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Email       string `json:"email" gorm:"unique"`
	Password    string `json:"password"`
	BirthDate   string `json:"birthdate"`
	Gender      string `json:"gender"`
	PhoneNumber string `json:"phone"`
	Roles       string `json:"roles"` // Este campo determina si el usuario es un "cuidador", "paciente" o cualquier otro rol que se necesite.
}

// Cuidador representa la información específica de los cuidadores.
type Cuidador struct {
	gorm.Model
	UserID   uint   `gorm:"uniqueIndex"` // Relación uno a uno asegurando que cada cuidador esté vinculado a un usuario único
	Relacion string `json:"relacion"`
	User     User   `gorm:"foreignKey:UserID"` // Referencia a User para establecer la relación
}

// Paciente representa la información específica de los pacientes.
type Paciente struct {
	gorm.Model
	UserID           uint   `gorm:"uniqueIndex"` // Relación uno a uno asegurando que cada paciente esté vinculado a un usuario único
	NumeroEmergencia string `json:"numeroEmergencia"`
	User             User   `gorm:"foreignKey:UserID"` // Referencia a User para establecer la relación
}
