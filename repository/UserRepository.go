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
	query := r.db.Model(&models.User{})

	if role != "" {
		query = query.Where("roles = ?", role)
	}

	// Preload de relaciones basado en el rol
	if role == "cuidador" {
		query = query.Preload("Cuidador")
	} else if role == "paciente" {
		query = query.Preload("Paciente")
	}

	result := query.Find(&users)
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
	// Inicia una transacción
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Verificar si el correo electrónico ya está registrado
	var count int64
	tx.Model(&models.User{}).Where("email = ?", user.Email).Count(&count)
	if count > 0 {
		tx.Rollback()
		return models.User{}, errors.New("el email ya está en uso")
	}

	// Intenta crear el usuario
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return models.User{}, err
	}

	// Lógica para manejar roles específicos
	switch user.Roles {
	case "cuidador":
		cuidador, ok := roleData.(*models.Cuidador) // Asegúrate de pasar un puntero al tipo correcto
		if !ok || cuidador == nil {
			tx.Rollback()
			return models.User{}, errors.New("datos de rol de cuidador inválidos")
		}
		cuidador.UserID = user.ID
		if err := tx.Create(cuidador).Error; err != nil {
			tx.Rollback()
			return models.User{}, err
		}
	case "paciente":
		paciente, ok := roleData.(*models.Paciente) // Asegúrate de pasar un puntero al tipo correcto
		if !ok || paciente == nil {
			tx.Rollback()
			return models.User{}, errors.New("datos de rol de paciente inválidos")
		}
		paciente.UserID = user.ID
		if err := tx.Create(paciente).Error; err != nil {
			tx.Rollback()
			return models.User{}, err
		}
	}

	// Si todo fue exitoso, finaliza la transacción
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return models.User{}, err
	}

	return user, nil
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
