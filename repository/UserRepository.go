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
// Suponiendo que esta función es parte de tu UserRepository
// UserRepository.go
func (r *userRepository) CreateUserWithRole(user models.User, roleData interface{}) (models.User, error) {
	// Iniciar una transacción
	tx := r.db.Begin()

	// Intentar crear el usuario
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback() // Deshacer la transacción si falla la creación del usuario
		return models.User{}, err
	}

	// Manejar la creación de roles específicos
	switch user.Roles {
	case "cuidador":
		cuidador, ok := roleData.(models.Cuidador)
		if !ok {
			tx.Rollback() // Deshacer la transacción si los datos del rol son inválidos
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
			tx.Rollback()
			return models.User{}, errors.New("invalid role data for paciente")
		}
		paciente.UserID = user.ID
		if err := tx.Create(&paciente).Error; err != nil {
			tx.Rollback()
			return models.User{}, err
		}
	}

	// Si todo salió bien, hacer commit de la transacción
	tx.Commit()
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
