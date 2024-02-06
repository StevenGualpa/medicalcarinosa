package repository

import (
	"GolandProyectos/models" // Asegúrate de ajustar esta importación a tu estructura de proyecto
	"errors"
	"gorm.io/gorm"
)

// UserRepository define los métodos para interactuar con la tabla de usuarios en la base de datos.
type UserRepository interface {
	GetUserByID(id uint) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(id uint) error
	GetAllUsers() ([]models.User, int, error)
	Login(email, password string) (models.User, string, error)
}

// userRepository implementa la interfaz UserRepository con una conexión a base de datos gorm.DB.
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository crea una nueva instancia de userRepository.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

// GetAllUsers retorna todos los usuarios de la base de datos y la cantidad.
func (r *userRepository) GetAllUsers() ([]models.User, int, error) {
	var users []models.User
	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return users, int(result.RowsAffected), nil
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

// Login verifica las credenciales del usuario y devuelve un mensaje junto con los datos del usuario.
func (r *userRepository) Login(email, password string) (models.User, string, error) {
	var user models.User
	// Busca el usuario por correo electrónico
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		// Si no se encuentra el usuario, devuelve un error
		return models.User{}, "Correo no existe", err
	}

	// TODO: Aquí deberías comparar la contraseña proporcionada con la almacenada de manera segura
	// Este ejemplo utiliza una comparación directa por simplicidad
	if user.Password != password {
		// Si las contraseñas no coinciden, devuelve un error
		return models.User{}, "Clave incorrecta", errors.New("clave incorrecta")
	}

	// Si la autenticación es exitosa, devuelve los datos del usuario y un mensaje de éxito
	return user, "Éxito", nil
}
