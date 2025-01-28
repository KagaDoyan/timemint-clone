package repositories

import (
	"go-fiber/domain/entities"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ShiftRepository interface {
	Create(shift entities.Shift) (*entities.Shift, error)
	FindAll(page, limit int) ([]entities.Shift, int64, error)
	FindById(id uint) (*entities.Shift, error)
	Update(id uint, shift entities.Shift) (*entities.Shift, error)
	Delete(id uint) error
}

type shiftRepositoryImpl struct {
	db *gorm.DB
}

func NewShiftRepository(db *gorm.DB) ShiftRepository {
	db.AutoMigrate(&entities.Shift{})
	return &shiftRepositoryImpl{db: db}
}

func (r shiftRepositoryImpl) Create(shift entities.Shift) (*entities.Shift, error) {
	err := r.db.Create(&shift).Error
	if err != nil {
		return nil, err
	}
	var result entities.Shift
	err = r.db.Preload(clause.Associations).First(&result, shift.ID).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r shiftRepositoryImpl) Update(id uint, shift entities.Shift) (*entities.Shift, error) {
	err := r.db.Model(&entities.Shift{}).Where("id = ?", id).Updates(&shift).Error
	if err != nil {
		return nil, err
	}
	var result entities.Shift
	err = r.db.Preload(clause.Associations).First(&result, id).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r shiftRepositoryImpl) Delete(id uint) error {
	err := r.db.Unscoped().Delete(&entities.Shift{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r shiftRepositoryImpl) FindAll(page, limit int) ([]entities.Shift, int64, error) {
	var shifts []entities.Shift
	var total int64
	err := r.db.Model(&entities.Shift{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	offset := (page - 1) * limit
	err = r.db.Offset(offset).Limit(limit).Preload(clause.Associations).Find(&shifts).Error
	if err != nil {
		return nil, 0, err
	}
	return shifts, total, nil
}

func (r shiftRepositoryImpl) FindById(id uint) (*entities.Shift, error) {
	var shift entities.Shift
	err := r.db.Preload(clause.Associations).First(&shift, id).Error
	if err != nil {
		return nil, err
	}
	return &shift, nil
}
