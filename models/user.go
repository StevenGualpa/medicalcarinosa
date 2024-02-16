package models

import "gorm.io/gorm"

// User representa un usuario general con información básica.
type User struct {
	gorm.Model
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Email       string `json:"email" gorm:"unique"`
	Password    string `json:"password"`
	BirthDate   string `json:"birthdate"`
	Gender      string `json:"gender"`
	PhoneNumber string `json:"phone"`
	Roles       string `json:"roles"` // Define el rol del usuario ("cuidador", "paciente", "admin", etc.).
	// Relaciones
	Cuidador *Cuidador `gorm:"foreignKey:UserID"`
	Paciente *Paciente `gorm:"foreignKey:UserID"`
}

// Cuidador contiene información específica para usuarios con el rol de cuidador.
type Cuidador struct {
	gorm.Model
	UserID   uint   `gorm:"uniqueIndex"` // Clave foránea que apunta a User.
	Relacion string `json:"relacion"`
	User     User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// Paciente contiene información específica para usuarios con el rol de paciente.
type Paciente struct {
	gorm.Model
	UserID           uint   `gorm:"uniqueIndex"` // Clave foránea que apunta a User.
	NumeroEmergencia string `json:"numeroEmergencia"`
	User             User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
