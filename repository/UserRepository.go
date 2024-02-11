package repository

import (
	"GolandProyectos/models"
	"errors"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(id uint) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(id uint) error
	GetAllUsers() ([]models.User, int, error)
	GetAllUsersWithRoleFilter(role string) ([]models.User, int, error) // Nuevo método agregado
	Login(email, password string) (models.User, string, error)
	CreateUserWithRole(user models.User, roleData interface{}) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAllUsers() ([]models.User, int, error) {
	var users []models.User
	result := r.db.Find(&users)
	return users, int(result.RowsAffected), result.Error
}

func (r *userRepository) GetAllUsersWithRoleFilter(role string) ([]models.User, int, error) {
	var users []models.User
	var result *gorm.DB

	if role == "" {
		result = r.db.Preload("Cuidador").Preload("Paciente").Find(&users)
	} else {
		result = r.db.Where("roles = ?", role).Preload("Cuidador").Preload("Paciente").Find(&users)
	}

	if result.Error != nil {
		return nil, 0, result.Error
	}

	return users, int(result.RowsAffected), nil
}

func (r *userRepository) GetUserByID(id uint) (models.User, error) {
	var user models.User
	result := r.db.First(&user, id)
	return user, result.Error
}

func (r *userRepository) CreateUser(user models.User) (models.User, error) {
	result := r.db.Create(&user)
	return user, result.Error
}

// Metodo para crear si es usuario adicioal si es cuidado o paciente

func (r *userRepository) CreateUserWithRole(user models.User, roleData interface{}) (models.User, error) {
	// Verificar si el correo electrónico ya está registrado
	var count int64
	r.db.Model(&models.User{}).Where("email = ?", user.Email).Count(&count)
	if count > 0 {
		// El correo electrónico ya está en uso
		return models.User{}, errors.New("El email ya esta en uso")
	}

	// Inicia una transacción
	tx := r.db.Begin()

	// Intenta crear el usuario
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback() // En caso de error, revierte la transacción
		return models.User{}, err
	}

	// Lógica para manejar roles específicos
	switch user.Roles {
	case "cuidador":
		cuidador, ok := roleData.(models.Cuidador)
		if !ok {
			tx.Rollback() // Revertir si los datos del rol no son válidos
			return models.User{}, errors.New("invalid role data for cuidador")
		}
		cuidador.UserID = user.ID
		if err := tx.Create(&cuidador).Error; err != nil {
			tx.Rollback()
			return models.User{}, err
		}
	case "paciente":
		paciente, ok := roleData.(models.Paciente)
		if !ok {
			tx.Rollback() // Revertir si los datos del rol no son válidos
			return models.User{}, errors.New("invalid role data for paciente")
		}
		paciente.UserID = user.ID
		if err := tx.Create(&paciente).Error; err != nil {
			tx.Rollback()
			return models.User{}, err
		}
	case "admin":
		// No se requiere acción adicional para el rol 'admin'
		// La lógica específica del rol 'admin' puede ser implementada aquí si es necesario
	}

	// Si todo fue exitoso, hace commit de la transacción
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return models.User{}, err
	}

	return user, nil // Devuelve el usuario creado con éxito
}

func (r *userRepository) UpdateUser(user models.User) (models.User, error) {
	result := r.db.Save(&user)
	return user, result.Error
}

func (r *userRepository) DeleteUser(id uint) error {
	result := r.db.Delete(&models.User{}, id)
	return result.Error
}

func (r *userRepository) Login(email, password string) (models.User, string, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return models.User{}, "Correo no existe", err
	}

	// Aquí deberías incluir la lógica para verificar la contraseña hasheada
	if user.Password != password { // Simplificación; usa bcrypt en producción
		return models.User{}, "Clave incorrecta", errors.New("clave incorrecta")
	}

	return user, "Éxito", nil
}
