// HorarioMedicamentosRepository.go
package repository

import (
	"GolandProyectos/models"
	"gorm.io/gorm"
	"time"
)

type HorarioMedicamentosRepository interface {
	Insert(pacienteID, medicamentoID uint, dosisInicial int) error
	InsertFinal(pacienteID, medicamentoID uint, dosisInicial, frecuencia int) error

	GetAll() ([]models.HorarioMedicamentoDetalle, error)
	GetAll2() ([]models.HorarioMedicineDetalle, int, error)

	Update(horarioMedicamento models.HorarioMedicamento) (models.HorarioMedicamento, error)
	Delete(id uint) error
	GetByID(id uint) (models.HorarioMedicamento, error)
	Delete2(pacienteID, medicamentoID uint) error
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
			HoraProxima:    horaActual.Add(time.Duration(5) * time.Hour),
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

func (repo *horarioMedicamentosRepository) GetAll2() ([]models.HorarioMedicineDetalle, int, error) {
	var detalles []models.HorarioMedicineDetalle
	var count int64

	// Primero, cuenta la cantidad de elementos que cumplen con la condición
	repo.db.Model(&models.HorarioMedicine{}).Where("deleted_at IS NULL").Count(&count)

	// Luego, ejecuta la consulta para obtener los detalles
	err := repo.db.Raw(`
        SELECT hm.id, hm.paciente_id, us.first_name, us.last_name,
               hm.medicamento_id, md.nombre, md.descripcion,
               hm.frecuencia, hm.dosis_restantes, hm.hora_inicial, hm.hora_proxima
        FROM horario_medicines AS hm
        JOIN medicines AS md ON hm.medicamento_id = md.id
        JOIN pacientes AS pc ON hm.paciente_id = pc.id
        JOIN users AS us ON pc.user_id = us.id
        WHERE hm.deleted_at IS NULL
    `).Scan(&detalles).Error

	if err != nil {
		return nil, 0, err
	}

	return detalles, int(count), nil
}

func (repo *horarioMedicamentosRepository) Update(horarioMedicamento models.HorarioMedicamento) (models.HorarioMedicamento, error) {
	if err := repo.db.Save(&horarioMedicamento).Error; err != nil {
		return models.HorarioMedicamento{}, err
	}
	return horarioMedicamento, nil
}

func (repo *horarioMedicamentosRepository) Delete(id uint) error {
	if err := repo.db.Delete(&models.HorarioMedicine{}, id).Error; err != nil {
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

// Metodos para el baboso de jordy
func (repo *horarioMedicamentosRepository) InsertFinal(pacienteID, medicamentoID uint, dosisInicial, frecuencia int) error {
	horaActual := time.Now()
	for i := 0; i < dosisInicial; i++ {
		horario := models.HorarioMedicine{
			PacienteID:     pacienteID,
			MedicamentoID:  medicamentoID,
			HoraInicial:    horaActual,
			HoraProxima:    horaActual.Add(time.Duration(frecuencia) * time.Hour),
			DosisRestantes: dosisInicial - i,
			Frecuencia:     frecuencia, // Asumiendo que agregaste este campo según los nuevos cambios
		}
		if err := repo.db.Create(&horario).Error; err != nil {
			return err
		}
		horaActual = horario.HoraProxima
	}
	return nil
}

func (repo *horarioMedicamentosRepository) Delete2(pacienteID, medicamentoID uint) error {
	result := repo.db.Where("paciente_id = ? AND medicamento_id = ?", pacienteID, medicamentoID).Delete(&models.HorarioMedicine{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
