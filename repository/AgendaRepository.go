package repository

import (
	"GolandProyectos/models"
	"gorm.io/gorm"
)

type AgendaRepository interface {
	Create(agenda models.Agenda) (models.Agenda, error)
	Update(agenda models.Agenda) (models.Agenda, error)
	Delete(id uint) error
	GetAll() ([]models.Agenda, error)
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

func (r *agendaRepository) GetAll() ([]models.Agenda, error) {
	var agendas []models.Agenda
	if err := r.db.Preload("PacienteCuidador").Preload("PacienteCuidador.Paciente").Preload("PacienteCuidador.Cuidador").Preload("PacienteCuidador.Paciente.User").Preload("PacienteCuidador.Cuidador.User").Find(&agendas).Error; err != nil {
		return nil, err
	}
	return agendas, nil
}

func (r *agendaRepository) GetById(id uint) (models.Agenda, error) {
	var agenda models.Agenda
	if err := r.db.Preload("PacienteCuidador").Preload("PacienteCuidador.Paciente").Preload("PacienteCuidador.Cuidador").Preload("PacienteCuidador.Paciente.User").Preload("PacienteCuidador.Cuidador.User").First(&agenda, id).Error; err != nil {
		return models.Agenda{}, err
	}
	return agenda, nil
}
