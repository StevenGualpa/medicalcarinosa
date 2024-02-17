package repository

import (
	"GolandProyectos/models"
	"gorm.io/gorm"
)

// PacienteCuidadorRepository define la interfaz para interactuar con la relación PacienteCuidador.
type PacienteCuidadorRepository interface {
	Create(pc models.PacienteCuidador) (models.PacienteCuidador, error)
	Update(pc models.PacienteCuidador) (models.PacienteCuidador, error)
	Delete(id uint) error
	GetAll() ([]models.PacienteCuidador, error)
	GetByPaciente(pacienteID uint) ([]models.Cuidador, error)
	GetByCuidador(cuidadorID uint) ([]models.Paciente, error)
}

type pacienteCuidadorRepository struct {
	db *gorm.DB
}

func NewPacienteCuidadorRepository(db *gorm.DB) PacienteCuidadorRepository {
	return &pacienteCuidadorRepository{db: db}
}

// Create inserta una nueva relación PacienteCuidador en la base de datos.
func (r *pacienteCuidadorRepository) Create(pc models.PacienteCuidador) (models.PacienteCuidador, error) {
	result := r.db.Create(&pc)
	if result.Error != nil {
		return models.PacienteCuidador{}, result.Error
	}
	return pc, nil
}

// Update modifica una relación PacienteCuidador existente en la base de datos.
func (r *pacienteCuidadorRepository) Update(pc models.PacienteCuidador) (models.PacienteCuidador, error) {
	result := r.db.Save(&pc)
	if result.Error != nil {
		return models.PacienteCuidador{}, result.Error
	}
	return pc, nil
}

// Delete elimina una relación PacienteCuidador de la base de datos.
func (r *pacienteCuidadorRepository) Delete(id uint) error {
	result := r.db.Delete(&models.PacienteCuidador{}, id)
	return result.Error
}

// GetAll devuelve todas las relaciones PacienteCuidador de la base de datos.
func (r *pacienteCuidadorRepository) GetAll() ([]models.PacienteCuidador, error) {
	var pcs []models.PacienteCuidador
	result := r.db.Find(&pcs)
	if result.Error != nil {
		return nil, result.Error
	}
	return pcs, nil
}

// GetByPaciente devuelve todos los cuidadores asignados a un paciente específico.
func (r *pacienteCuidadorRepository) GetByPaciente(pacienteID uint) ([]models.Cuidador, error) {
	var cuidadores []models.Cuidador
	result := r.db.Joins("JOIN paciente_cuidadors on paciente_cuidadors.cuidador_id = cuidadors.id").Where("paciente_cuidadors.paciente_id = ?", pacienteID).Find(&cuidadores)
	if result.Error != nil {
		return nil, result.Error
	}
	return cuidadores, nil
}

// GetByCuidador devuelve todos los pacientes asignados a un cuidador específico.
func (r *pacienteCuidadorRepository) GetByCuidador(cuidadorID uint) ([]models.Paciente, error) {
	var pacientes []models.Paciente
	result := r.db.Joins("JOIN paciente_cuidadors on paciente_cuidadors.paciente_id = pacientes.id").Where("paciente_cuidadors.cuidador_id = ?", cuidadorID).Find(&pacientes)
	if result.Error != nil {
		return nil, result.Error
	}
	return pacientes, nil
}
