package repositories

import (
	"go-fiber/domain/entities"

	"gorm.io/gorm"
)

type ActiveLocationRepository interface {
	FindAll(page, limit int) ([]entities.ActiveLocation, int64, error)
	FindByID(id uint) (*entities.ActiveLocation, error)
	Create(activeLocation entities.ActiveLocation) (*entities.ActiveLocation, error)
	Update(id uint, activeLocation entities.ActiveLocation) (*entities.ActiveLocation, error)
	Delete(id uint) error
}

type activeLocationRepository struct {
	db *gorm.DB
}

func NewActiveLocationRepository(db *gorm.DB) ActiveLocationRepository {
	db.AutoMigrate(&entities.ActiveLocation{})
	return &activeLocationRepository{db: db}
}

func (a activeLocationRepository) FindAll(page, limit int) ([]entities.ActiveLocation, int64, error) {
	var activeLocations []entities.ActiveLocation
	var total int64

	// Calculate offset
	offset := (page - 1) * limit

	// Count total records
	if err := a.db.Model(&entities.ActiveLocation{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Find paginated records
	err := a.db.Offset(offset).Limit(limit).Find(&activeLocations).Error
	if err != nil {
		return nil, 0, err
	}
	return activeLocations, total, nil
}

func (a activeLocationRepository) Create(activeLocation entities.ActiveLocation) (*entities.ActiveLocation, error) {
	err := a.db.Create(&activeLocation).Error
	if err != nil {
		return nil, err
	}
	var result entities.ActiveLocation
	err = a.db.First(&result, activeLocation.ID).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (a activeLocationRepository) Update(id uint, activeLocation entities.ActiveLocation) (*entities.ActiveLocation, error) {
	err := a.db.Model(&entities.ActiveLocation{}).Select("name", "description", "latitude", "longitude").Where("id = ?", id).Updates(&activeLocation).Error
	if err != nil {
		return nil, err
	}
	var result entities.ActiveLocation
	err = a.db.First(&result, id).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (a activeLocationRepository) Delete(id uint) error {
	err := a.db.Unscoped().Delete(&entities.ActiveLocation{}, "id = ?", id).Error
	if err != nil {
		return err
	}
	return nil
}

func (a activeLocationRepository) FindByID(id uint) (*entities.ActiveLocation, error) {
	var activeLocation entities.ActiveLocation
	err := a.db.First(&activeLocation, id).Error
	if err != nil {
		return nil, err
	}
	return &activeLocation, nil
}
