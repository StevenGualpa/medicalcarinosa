// medicamentoRepository.go
package repository

import (
	"GolandProyectos/models"
	"gorm.io/gorm"
)

type MedicamentoRepository interface {
	Create(medicamento models.Medicamento) (models.Medicamento, error)
	Update(medicamento models.Medicamento) (models.Medicamento, error)
	Delete(id uint) error
	GetAll() ([]models.Medicamento, error)
	GetById(id uint) (models.Medicamento, error)
}

type medicamentoRepository struct {
	db *gorm.DB
}

func NewMedicamentoRepository(db *gorm.DB) MedicamentoRepository {
	return &medicamentoRepository{db}
}

func (r *medicamentoRepository) Create(medicamento models.Medicamento) (models.Medicamento, error) {
	result := r.db.Create(&medicamento)
	return medicamento, result.Error
}

func (r *medicamentoRepository) Update(medicamento models.Medicamento) (models.Medicamento, error) {
	result := r.db.Save(&medicamento)
	return medicamento, result.Error
}

func (r *medicamentoRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Medicamento{}, id)
	return result.Error
}

func (r *medicamentoRepository) GetAll() ([]models.Medicamento, error) {
	var medicamentos []models.Medicamento
	result := r.db.Find(&medicamentos)
	return medicamentos, result.Error
}

func (r *medicamentoRepository) GetById(id uint) (models.Medicamento, error) {
	var medicamento models.Medicamento
	result := r.db.First(&medicamento, id)
	return medicamento, result.Error
}
