package repository

import (
	"GolandProyectos/models"
	"errors"
	"gorm.io/gorm"
	"strconv"
)

type UserRepository interface {
	GetUserByID(id uint) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(id uint) error
	GetAllUsers() ([]models.User, int, error)
	GetAllUsersWithRoleFilter(role string) ([]models.User, int, error) // Nuevo método agregado
	Login(email, password string) (models.User, string, error)
	Login2(email, password string) (models.User, string, error)

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

	// Aplica el filtro de rol, si se proporciona
	if role != "" {
		query = query.Where("roles = ?", role)
	}

	// Preload de las relaciones basado en el rol
	// Asegúrate de que los nombres usados en Preload coincidan con los de tus definiciones de modelo
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

	// Verifica la cédula ecuatoriana antes de crear el usuario
	if !isValidEcuadorianID(user.Cedula) {
		return models.User{}, errors.New("la cédula proporcionada no es válida")
	}

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
		cuidador, ok := roleData.(*models.Cuidador)
		if !ok || cuidador == nil {
			tx.Rollback()
			return models.User{}, errors.New("datos de rol de cuidador inválidos")
		}
		cuidador.UserID = user.ID
		if err := tx.Create(&cuidador).Error; err != nil {
			tx.Rollback()
			return models.User{}, err
		}
	case "paciente":
		paciente, ok := roleData.(*models.Paciente)
		if !ok || paciente == nil {
			tx.Rollback()
			return models.User{}, errors.New("datos de rol de paciente inválidos")
		}
		paciente.UserID = user.ID
		if err := tx.Create(&paciente).Error; err != nil {
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

// Asegúrate de que la función isValidEcuadorianID esté accesible para este paquete o esté definida dentro del mismo.

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

func isValidEcuadorianID(id string) bool {
	if len(id) != 10 {
		return false
	}

	coefficients := []int{2, 1, 2, 1, 2, 1, 2, 1, 2}
	sum := 0

	for i := 0; i < 9; i++ {
		digit, _ := strconv.Atoi(string(id[i]))
		result := coefficients[i] * digit

		if result >= 10 {
			result -= 9
		}
		sum += result
	}

	lastDigit, _ := strconv.Atoi(string(id[9]))
	remainder := sum % 10
	checkDigit := 0

	if remainder != 0 {
		checkDigit = 10 - remainder
	}

	return checkDigit == lastDigit
}

func (r *userRepository) Login2(email, password string) (models.User, string, error) {
	var user models.User
	// Primero verifica si el usuario existe y si la contraseña es correcta
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return models.User{}, "Correo no existe", err
	}

	// Aquí deberías incluir la lógica para verificar la contraseña hasheada
	if user.Password != password { // Simplificación; usa bcrypt en producción
		return models.User{}, "Clave incorrecta", errors.New("clave incorrecta")
	}

	// Según el rol del usuario, realiza la consulta adicional necesaria
	if user.Roles == "admin" {
		// Lógica para el rol de admin
		// Realiza la consulta específica de admin y actualiza el usuario si es necesario
	} else if user.Roles == "paciente" {
		// Lógica para el rol de paciente
		// Realiza la consulta específica de paciente y actualiza el usuario si es necesario
		var paciente models.Paciente
		if err := r.db.Where("user_id = ?", user.ID).First(&paciente).Error; err != nil {
			return models.User{}, "Paciente no encontrado", err
		}
		user.Paciente = &paciente
	} else if user.Roles == "cuidador" {
		// Lógica para el rol de cuidador
		// Realiza la consulta específica de cuidador y actualiza el usuario si es necesario
		var cuidador models.Cuidador
		if err := r.db.Where("user_id = ?", user.ID).First(&cuidador).Error; err != nil {
			return models.User{}, "Cuidador no encontrado", err
		}
		user.Cuidador = &cuidador
	}

	return user, "Éxito", nil
}
