package repository

import (
	"GolandProyectos/models"
	"gorm.io/gorm"
)

type AgendaRepository interface {
	Create(agenda models.Agenda) (models.Agenda, error)
	Update(agenda models.Agenda) (models.Agenda, error)
	Delete(id uint) error
	GetAll() ([]models.AgendaDetalle, error)
	GetById(id uint) (models.Agenda, error)
}

type agendaRepository struct {
	db *gorm.DB
}

func NewAgendaRepository(db *gorm.DB) AgendaRepository {
	return &agendaRepository{db}
}

func (r *agendaRepository) Create(agenda models.Agenda) (models.Agenda, error) {
	if err := r.db.Create(&agenda).Error; err != nil {
		return models.Agenda{}, err
	}
	return agenda, nil
}

func (r *agendaRepository) Update(agenda models.Agenda) (models.Agenda, error) {
	if err := r.db.Save(&agenda).Error; err != nil {
		return models.Agenda{}, err
	}
	return agenda, nil
}

func (r *agendaRepository) Delete(id uint) error {
	return r.db.Delete(&models.Agenda{}, id).Error
}

func (r *agendaRepository) GetAll() ([]models.AgendaDetalle, error) {
	var detalles []models.AgendaDetalle

	// Tu consulta SQL personalizada
	query := `
	SELECT 
	  ag.id, 
	  ag.nombre, 
	  ag.descripcion, 
	  ag.fecha, 
	  ag.hora, 
	  ag.estado,
	  pa.id AS paciente_id, 
	  us.first_name AS paciente_nombre, 
	  us.last_name AS paciente_apellido
	FROM 
	  agendas AS ag
	  INNER JOIN pacientes AS pa ON ag.paciente_id = pa.id
	  INNER JOIN users AS us ON pa.user_id = us.id
    `

	// Ejecutar la consulta
	if err := r.db.Raw(query).Scan(&detalles).Error; err != nil {
		return nil, err
	}

	return detalles, nil
}

func (r *agendaRepository) GetById(id uint) (models.Agenda, error) {
	var agenda models.Agenda
	if err := r.db.Preload("PacienteCuidador").Preload("PacienteCuidador.Paciente").Preload("PacienteCuidador.Cuidador").Preload("PacienteCuidador.Paciente.User").Preload("PacienteCuidador.Cuidador.User").First(&agenda, id).Error; err != nil {
		return models.Agenda{}, err
	}
	return agenda, nil
}
