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
	if err := repo.db.Preload("Paciente").Preload("Paciente.User").Preload("Medicamento").Find(&horariosMedicamentos).Error; err != nil {
		return nil, err
	}

	// Una vez que tienes los registros, asegúrate de que el ID del usuario no es cero
	// y si es así, puedes intentar cargar los datos del usuario manualmente
	for i, horario := range horariosMedicamentos {
		if horario.Paciente.User.ID == 0 && horario.Paciente.UserID != 0 {
			var user models.User
			if err := repo.db.First(&user, horario.Paciente.UserID).Error; err == nil {
				horariosMedicamentos[i].Paciente.User = user
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
