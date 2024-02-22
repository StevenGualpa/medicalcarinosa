// HorarioMedicamentosRepository.go
package repository

import (
	"GolandProyectos/models"
	"gorm.io/gorm"
	"time"
)

type HorarioMedicamentosRepository interface {
	Insert(pacienteID, medicamentoID uint, dosisInicial int) error
	GetAll() ([]models.HorarioMedicamentoDetalle, error)
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

func (repo *horarioMedicamentosRepository) GetAll() ([]models.HorarioMedicamentoDetalle, error) {
	var detalles []models.HorarioMedicamentoDetalle
	// Define una consulta que une las tablas y selecciona los campos necesarios.
	// Reemplaza "horario_medicamentos", "pacientes", "users", "medicamentos" con los nombres reales de tus tablas.
	err := repo.db.Table("horario_medicamentos").Select(
		"horario_medicamentos.id, horario_medicamentos.paciente_id, users.first_name, users.last_name, " +
			"horario_medicamentos.medicamento_id, medicamentos.nombre, medicamentos.descripcion, " +
			"medicamentos.numero_dosis, medicamentos.frecuencia, horario_medicamentos.hora_inicial, " +
			"horario_medicamentos.hora_proxima, horario_medicamentos.dosis_restantes").
		Joins("left join pacientes on pacientes.id = horario_medicamentos.paciente_id").
		Joins("left join users on users.id = pacientes.user_id").
		Joins("left join medicamentos on medicamentos.id = horario_medicamentos.medicamento_id").
		Scan(&detalles).Error

	if err != nil {
		return nil, err
	}

	return detalles, nil
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
