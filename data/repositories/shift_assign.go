package repositories

import (
	"go-fiber/domain/entities"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type shiftAssignRepository struct {
	db *gorm.DB
}

type ShiftAssignRepository interface {
	Create(shiftAssign entities.ShiftAssignment) (*entities.ShiftAssignment, error)
	BatchCreate(shiftAssigns []entities.ShiftAssignment) ([]entities.ShiftAssignment, error)
	Delete(id uint) error
	FindAll(page, limit int) ([]entities.ShiftAssignment, int64, error)
	FindById(id uint) (*entities.ShiftAssignment, error)
}

func NewShiftAssignRepository(db *gorm.DB) ShiftAssignRepository {
	return &shiftAssignRepository{
		db: db,
	}
}

func (r shiftAssignRepository) Create(shiftAssign entities.ShiftAssignment) (*entities.ShiftAssignment, error) {
	err := r.db.Create(&shiftAssign).Error
	if err != nil {
		return nil, err
	}
	var result entities.ShiftAssignment
	err = r.db.First(&result, shiftAssign.ID).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r shiftAssignRepository) BatchCreate(shiftAssigns []entities.ShiftAssignment) ([]entities.ShiftAssignment, error) {
	err := r.db.Create(&shiftAssigns).Error
	if err != nil {
		return nil, err
	}
	return shiftAssigns, nil
}

func (r shiftAssignRepository) Delete(id uint) error {
	err := r.db.Delete(&entities.ShiftAssignment{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r shiftAssignRepository) FindAll(page, limit int) ([]entities.ShiftAssignment, int64, error) {
	var shiftAssigns []entities.ShiftAssignment
	var total int64
	err := r.db.Model(&entities.ShiftAssignment{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	offset := (page - 1) * limit
	err = r.db.Offset(offset).Limit(limit).Preload(clause.Associations).Find(&shiftAssigns).Error
	if err != nil {
		return nil, 0, err
	}
	return shiftAssigns, total, nil
}

func (r shiftAssignRepository) FindById(id uint) (*entities.ShiftAssignment, error) {
	var shiftAssign entities.ShiftAssignment
	err := r.db.Preload(clause.Associations).First(&shiftAssign, id).Error
	if err != nil {
		return nil, err
	}
	return &shiftAssign, nil
}
