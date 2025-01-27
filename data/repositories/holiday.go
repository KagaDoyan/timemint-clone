package repositories

import (
	"go-fiber/domain/entities"
	"time"

	"gorm.io/gorm"
)

type HolidayRepository interface {
	FindAll(page, limit int) ([]entities.Holiday, int64, error)
	Create(holiday entities.Holiday) (*entities.Holiday, error)
	Update(id uint, holiday entities.Holiday) (*entities.Holiday, error)
	Delete(id uint) error
	IsHoliday(date time.Time) (bool, error)
}

type holidayRepository struct {
	db *gorm.DB
}

func NewHolidayRepository(db *gorm.DB) HolidayRepository {
	db.AutoMigrate(&entities.Holiday{})
	return &holidayRepository{db: db}
}

func (r holidayRepository) FindAll(page, limit int) ([]entities.Holiday, int64, error) {
	var holidays []entities.Holiday
	var count int64
	err := r.db.Find(&holidays).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Apply pagination
	err = r.db.Limit(limit).Offset(offset).Find(&holidays).Error
	if err != nil {
		return nil, 0, err
	}
	return holidays, count, nil
}

func (r holidayRepository) Create(holiday entities.Holiday) (*entities.Holiday, error) {
	err := r.db.Create(&holiday).Error
	if err != nil {
		return nil, err
	}
	var result entities.Holiday
	err = r.db.First(&result, holiday.ID).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r holidayRepository) Update(id uint, holiday entities.Holiday) (*entities.Holiday, error) {
	err := r.db.Model(&entities.Holiday{}).Where("id = ?", id).Updates(holiday).Error
	if err != nil {
		return nil, err
	}
	var result entities.Holiday
	err = r.db.First(&result, id).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r holidayRepository) Delete(id uint) error {
	err := r.db.Unscoped().Delete(&entities.Holiday{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r holidayRepository) IsHoliday(date time.Time) (bool, error) {
	var holiday entities.Holiday
	//check if date between start and end date
	err := r.db.Where("start_date <= ? AND end_date >= ?", date, date).First(&holiday).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
