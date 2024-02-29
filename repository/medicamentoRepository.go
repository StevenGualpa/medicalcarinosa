// medicamentoRepository.go
package repository

import (
	"GolandProyectos/models"
	"gorm.io/gorm"
)

type MedicineRepository interface {
	Create(medicine models.Medicine) (models.Medicine, error)
	Update(medicine models.Medicine) (models.Medicine, error)
	Delete(id uint) error
	GetAll() ([]models.Medicine, error)
	GetById(id uint) (models.Medicine, error)
}

type medicineRepository struct {
	db *gorm.DB
}

func NewMedicineRepository(db *gorm.DB) MedicineRepository {
	return &medicineRepository{db}
}

func (r *medicineRepository) Create(medicine models.Medicine) (models.Medicine, error) {
	result := r.db.Create(&medicine)
	return medicine, result.Error
}

func (r *medicineRepository) Update(medicine models.Medicine) (models.Medicine, error) {
	result := r.db.Save(&medicine)
	return medicine, result.Error
}

func (r *medicineRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Medicine{}, id)
	return result.Error
}

func (r *medicineRepository) GetAll() ([]models.Medicine, error) {
	var medicines []models.Medicine
	result := r.db.Find(&medicines)
	return medicines, result.Error
}

func (r *medicineRepository) GetById(id uint) (models.Medicine, error) {
	var medicine models.Medicine
	result := r.db.First(&medicine, id)
	return medicine, result.Error
}
