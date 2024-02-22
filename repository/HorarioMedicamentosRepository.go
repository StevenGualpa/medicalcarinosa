// HorarioMedicamentosRepository.go
package repository

import (
	"GolandProyectos/models"
	"gorm.io/gorm"
	"time"
)

type HorarioMedicamentosRepository interface {
	Insert(pacienteID, medicamentoID uint, dosisInicial int) error
	GetAll() ([]models.HorarioMedicamento, error)
	Update(horarioMedicamento models.HorarioMedicamento) (models.HorarioMedicamento, error)
	Delete(id uint) error
	GetByID(id uint) (models.HorarioMedicamento, error)
}

type horarioMedicamentosRepository struct {
	db *gorm.DB
}

func NewHorarioMedicamentosRepository(db *gorm.DB) HorarioMedicamentosRepository {
	return &horarioMedicamentosRepository{db: db}
}

func (repo *horarioMedicamentosRepository) Insert(pacienteID, medicamentoID uint, dosisInicial int) error {
	var medicamento models.Medicamento
	if err := repo.db.First(&medicamento, medicamentoID).Error; err != nil {
		return err
	}

	horaActual := time.Now()
	for i := 0; i < dosisInicial; i++ {
		horario := models.HorarioMedicamento{
			PacienteID:     pacienteID,
			MedicamentoID:  medicamentoID,
			HoraInicial:    horaActual,
			HoraProxima:    horaActual.Add(time.Duration(medicamento.Frecuencia) * time.Hour),
			DosisRestantes: dosisInicial - i,
		}
		if err := repo.db.Create(&horario).Error; err != nil {
			return err
		}
		horaActual = horario.HoraProxima
	}
	return nil
}

func (repo *horarioMedicamentosRepository) GetAll() ([]models.HorarioMedicamento, error) {
	var horariosMedicamentos []models.HorarioMedicamento
	if err := repo.db.Preload("Paciente").Preload("Medicamento").Find(&horariosMedicamentos).Error; err != nil {
		return nil, err
	}

	// Itera sobre los horarios de medicamentos para asegurarte de que la información del usuario está cargada
	for i, horario := range horariosMedicamentos {
		if horario.Paciente.User.ID == 0 {
			var user models.User
			// Realiza una consulta adicional para obtener la información del usuario
			if err := repo.db.Where("id = ?", horario.Paciente.UserID).First(&user).Error; err == nil {
				// Asigna el usuario al paciente dentro del horario de medicamento
				horariosMedicamentos[i].Paciente.User = user
			} else {
				// Maneja el error si el usuario no se encuentra
				// Puedes decidir si quieres devolver un error, continuar sin el usuario, etc.
			}
		}
	}

	return horariosMedicamentos, nil
}

func (repo *horarioMedicamentosRepository) Update(horarioMedicamento models.HorarioMedicamento) (models.HorarioMedicamento, error) {
	if err := repo.db.Save(&horarioMedicamento).Error; err != nil {
		return models.HorarioMedicamento{}, err
	}
	return horarioMedicamento, nil
}

func (repo *horarioMedicamentosRepository) Delete(id uint) error {
	if err := repo.db.Delete(&models.HorarioMedicamento{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (repo *horarioMedicamentosRepository) GetByID(id uint) (models.HorarioMedicamento, error) {
	var horario models.HorarioMedicamento
	if err := repo.db.Preload("Paciente").Preload("Paciente.User").Preload("Medicamento").First(&horario, id).Error; err != nil {
		return models.HorarioMedicamento{}, err
	}
	return horario, nil
}
