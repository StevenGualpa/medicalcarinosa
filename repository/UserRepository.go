// repository/user_repository.go
package repository

import (
	"GolandProyectos/models" // Cambia esto por la ruta correcta al paquete models en tu proyecto

	"gorm.io/gorm"
)

// UserRepository es la interfaz que define los métodos para interactuar con la tabla de usuarios en la base de datos.
type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uint) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(id uint) error
}

// userRepository implementa la interfaz UserRepository con una conexión a base de datos gorm.DB.
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository crea una nueva instancia de userRepository.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

// GetAllUsers retorna todos los usuarios de la base de datos.
func (r *userRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

// GetUserByID retorna un usuario por su ID.
func (r *userRepository) GetUserByID(id uint) (models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return user, err
}

// CreateUser inserta un nuevo usuario en la base de datos.
func (r *userRepository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

// UpdateUser actualiza un usuario existente en la base de datos.
func (r *userRepository) UpdateUser(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error
	return user, err
}

// DeleteUser elimina un usuario de la base de datos por su ID.
func (r *userRepository) DeleteUser(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}
