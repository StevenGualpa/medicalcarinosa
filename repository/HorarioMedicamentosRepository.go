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

	// Recolecta los IDs de Paciente que necesitan la carga manual de User
	pacienteIDsNecesitanUser := make([]uint, 0)
	for _, horario := range horariosMedicamentos {
		if horario.Paciente.User.ID == 0 && horario.Paciente.UserID != 0 {
			pacienteIDsNecesitanUser = append(pacienteIDsNecesitanUser, horario.Paciente.UserID)
		}
	}

	// Elimina duplicados de pacienteIDsNecesitanUser
	uniquePacienteIDs := make(map[uint]bool)
	for _, id := range pacienteIDsNecesitanUser {
		uniquePacienteIDs[id] = true
	}

	// Carga todos los Users necesarios en una sola consulta si hay IDs para buscar
	if len(uniquePacienteIDs) > 0 {
		var users []models.User
		if err := repo.db.Where("id IN ?", pacienteIDsNecesitanUser).Find(&users).Error; err != nil {
			return nil, err
		}

		// Mapa para acceso rápido a los Users por ID
		userMap := make(map[uint]models.User)
		for _, user := range users {
			userMap[user.ID] = user
		}

		// Asigna manualmente el User a cada Paciente necesario
		for i, horario := range horariosMedicamentos {
			if user, ok := userMap[horario.Paciente.UserID]; ok {
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
