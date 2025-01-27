package services

import (
	"go-fiber/data/repositories"
	"go-fiber/domain/entities"
	"go-fiber/domain/models"
)

type activeLocationService struct {
	activeLocationRepo repositories.ActiveLocationRepository
}

type ActiveLocationService interface {
	FindAll(page, limit int) ([]models.ActiveLocation, int64, error)
	FindByID(id uint) (*models.ActiveLocation, error)
	Create(activeLocation models.ActiveLocation) (*models.ActiveLocation, error)
	Update(id uint, activeLocation models.ActiveLocation) (*models.ActiveLocation, error)
	Delete(id uint) error
}

func NewActiveLocationService(activeLocationRepo repositories.ActiveLocationRepository) ActiveLocationService {
	return &activeLocationService{activeLocationRepo: activeLocationRepo}
}

func (s activeLocationService) FindAll(page, limit int) ([]models.ActiveLocation, int64, error) {
	data, total, err := s.activeLocationRepo.FindAll(page, limit)
	if err != nil {
		return nil, 0, err
	}
	var result []models.ActiveLocation
	for _, activeLocation := range data {
		result = append(result, models.ActiveLocation{
			ID:          activeLocation.ID,
			Name:        activeLocation.Name,
			Description: activeLocation.Description,
			Latitude:    activeLocation.Latitude,
			Longitude:   activeLocation.Longitude,
		})
	}
	return result, total, nil
}

func (s activeLocationService) Create(activeLocation models.ActiveLocation) (*models.ActiveLocation, error) {
	locationEntity := entities.ActiveLocation{
		Name:        activeLocation.Name,
		Description: activeLocation.Description,
		Latitude:    activeLocation.Latitude,
		Longitude:   activeLocation.Longitude,
	}
	data, err := s.activeLocationRepo.Create(locationEntity)
	if err != nil {
		return nil, err
	}
	return &models.ActiveLocation{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		Latitude:    data.Latitude,
		Longitude:   data.Longitude,
	}, nil
}

func (s activeLocationService) Update(id uint, activeLocation models.ActiveLocation) (*models.ActiveLocation, error) {
	locationEntity := entities.ActiveLocation{
		Name:        activeLocation.Name,
		Description: activeLocation.Description,
		Latitude:    activeLocation.Latitude,
		Longitude:   activeLocation.Longitude,
	}
	data, err := s.activeLocationRepo.Update(id, locationEntity)
	if err != nil {
		return nil, err
	}
	return &models.ActiveLocation{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		Latitude:    data.Latitude,
		Longitude:   data.Longitude,
	}, nil
}

func (s activeLocationService) Delete(id uint) error {
	return s.activeLocationRepo.Delete(id)
}

func (s activeLocationService) FindByID(id uint) (*models.ActiveLocation, error) {
	data, err := s.activeLocationRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return &models.ActiveLocation{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		Latitude:    data.Latitude,
		Longitude:   data.Longitude,
	}, nil
}
