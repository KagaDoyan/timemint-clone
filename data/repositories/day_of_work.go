package repositories

import (
	"go-fiber/domain/entities"
	"time"

	"gorm.io/gorm"
)

type DayOfWorkRepository interface {
	FindAll(page, limit int) ([]entities.DayOfWork, int64, error)
	FindByID(id uint) (*entities.DayOfWork, error)
	Create(dayOfWork entities.DayOfWork) (*entities.DayOfWork, error)
	Update(id uint, dayOfWork entities.DayOfWork) (*entities.DayOfWork, error)
	Delete(id uint) error
	FindByDay(day time.Time) (*entities.DayOfWork, error)
}

type dayOfWorkRepository struct {
	db *gorm.DB
}

func NewDayOfWorkRepository(db *gorm.DB) DayOfWorkRepository {
	db.AutoMigrate(&entities.DayOfWork{})
	return &dayOfWorkRepository{db: db}
}

func (s dayOfWorkRepository) FindByDay(day time.Time) (*entities.DayOfWork, error) {
	var dayOfWork entities.DayOfWork
	err := s.db.Where("day = ?", day.Weekday().String()).First(&dayOfWork).Error
	if err != nil {
		return nil, err
	}
	return &dayOfWork, nil
}

func (s dayOfWorkRepository) FindAll(page, limit int) ([]entities.DayOfWork, int64, error) {
	var dayOfWorks []entities.DayOfWork
	var total int64

	// Calculate offset
	offset := (page - 1) * limit

	// Count total records
	if err := s.db.Model(&entities.DayOfWork{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Find paginated records
	err := s.db.Offset(offset).Limit(limit).Find(&dayOfWorks).Error
	if err != nil {
		return nil, 0, err
	}
	return dayOfWorks, total, nil
}

func (s dayOfWorkRepository) FindByID(id uint) (*entities.DayOfWork, error) {
	var dayOfWork entities.DayOfWork
	err := s.db.First(&dayOfWork, id).Error
	if err != nil {
		return nil, err
	}
	return &dayOfWork, nil
}

func (s dayOfWorkRepository) Create(dayOfWork entities.DayOfWork) (*entities.DayOfWork, error) {
	err := s.db.Create(&dayOfWork).Error
	if err != nil {
		return nil, err
	}
	var result entities.DayOfWork
	err = s.db.First(&result, dayOfWork.ID).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s dayOfWorkRepository) Update(id uint, dayOfWork entities.DayOfWork) (*entities.DayOfWork, error) {
	err := s.db.Model(&entities.DayOfWork{}).Select("day", "start_time", "end_time", "is_work_day").Where("id = ?", id).Updates(&dayOfWork).Error
	if err != nil {
		return nil, err
	}
	var result entities.DayOfWork
	err = s.db.First(&result, id).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s dayOfWorkRepository) Delete(id uint) error {
	err := s.db.Unscoped().Delete(&entities.DayOfWork{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
