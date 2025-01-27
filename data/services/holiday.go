package services

import (
	"go-fiber/data/repositories"
	"go-fiber/domain/entities"
	"go-fiber/domain/models"
	"time"
)

type HolidayService interface {
	FindAll(page, limit int) ([]models.Holiday, int64, error)
	Create(holiday models.Holiday) (*models.Holiday, error)
	Update(id uint, holiday models.Holiday) (*models.Holiday, error)
	Delete(id uint) error
	IsHoliday(date string) (bool, error)
}

type holidayService struct {
	repository repositories.HolidayRepository
}

func NewHolidayService(repository repositories.HolidayRepository) HolidayService {
	return &holidayService{repository}
}

func (s holidayService) IsHoliday(date string) (bool, error) {
	d, err := time.Parse("02-01-2006", date)
	if err != nil {
		return false, err
	}
	return s.repository.IsHoliday(d)
}

func (s holidayService) FindAll(page, limit int) ([]models.Holiday, int64, error) {
	data, count, err := s.repository.FindAll(page, limit)
	if err != nil {
		return nil, 0, err
	}
	var result []models.Holiday
	for _, holiday := range data {
		result = append(result, models.Holiday{
			ID:          holiday.ID,
			Name:        holiday.Name,
			Description: holiday.Description,
			StartDate:   holiday.StartDate.Format("02-01-2006"),
			EndDate:     holiday.EndDate.Format("02-01-2006"),
			CreatedAt:   holiday.CreatedAt.Format("02-01-2006"),
			UpdatedAt:   holiday.UpdatedAt.Format("02-01-2006"),
			CreatedBy:   holiday.CreatedBy,
		})
	}
	return result, count, nil
}

func (s holidayService) Create(holiday models.Holiday) (*models.Holiday, error) {
	startDate, err := time.Parse("02-01-2006", holiday.StartDate)
	if err != nil {
		return nil, err
	}
	endDate, err := time.Parse("02-01-2006", holiday.EndDate)
	if err != nil {
		return nil, err
	}
	result, err := s.repository.Create(entities.Holiday{
		Name:        holiday.Name,
		Description: holiday.Description,
		StartDate:   startDate,
		EndDate:     endDate,
		CreatedBy:   holiday.CreatedBy,
	})
	if err != nil {
		return nil, err
	}
	return &models.Holiday{
		ID:          result.ID,
		Name:        result.Name,
		Description: result.Description,
		StartDate:   result.StartDate.Format("02-01-2006"),
		EndDate:     result.EndDate.Format("02-01-2006"),
	}, nil
}

func (s holidayService) Update(id uint, holiday models.Holiday) (*models.Holiday, error) {
	startDate, err := time.Parse("02-01-2006", holiday.StartDate)
	if err != nil {
		return nil, err
	}
	endDate, err := time.Parse("02-01-2006", holiday.EndDate)
	if err != nil {
		return nil, err
	}
	result, err := s.repository.Update(id, entities.Holiday{
		Name:        holiday.Name,
		Description: holiday.Description,
		StartDate:   startDate,
		EndDate:     endDate,
	})
	if err != nil {
		return nil, err
	}
	return &models.Holiday{
		ID:          result.ID,
		Name:        result.Name,
		Description: result.Description,
		StartDate:   result.StartDate.Format("02-01-2006"),
		EndDate:     result.EndDate.Format("02-01-2006"),
	}, nil
}

func (s holidayService) Delete(id uint) error {
	return s.repository.Delete(id)
}
