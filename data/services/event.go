package services

import (
	"go-fiber/core/logs"
	"go-fiber/data/repositories"
	"go-fiber/domain/entities"
	"go-fiber/domain/models"
	"time"

	"gorm.io/gorm"
)

type EventService interface {
	FindAll(page, limit int) ([]models.Event, int64, error)
	Create(Event models.Event) (int, error)
	Update(id uint, Event models.Event) (*models.Event, error)
	Delete(id uint) error
	IsEvent(date string) (bool, error)
	CalendarEvent(userID, month, year int) ([]models.Event, error)
}

type eventService struct {
	repository repositories.EventRepository
}

func NewEventService(repository repositories.EventRepository) EventService {
	return &eventService{repository}
}

func (s eventService) IsEvent(date string) (bool, error) {
	d, err := time.Parse("02-01-2006", date)
	if err != nil {
		return false, err
	}
	return s.repository.IsEvent(d)
}

func (s eventService) FindAll(page, limit int) ([]models.Event, int64, error) {
	data, count, err := s.repository.FindAll(page, limit)
	if err != nil {
		return nil, 0, err
	}
	var result []models.Event
	for _, Event := range data {
		var invites []models.Employee
		for _, invite := range Event.Invites {
			invites = append(invites, models.Employee{
				ID:   invite.ID,
				Name: invite.Name,
			})
		}
		result = append(result, models.Event{
			ID:          Event.ID,
			Name:        Event.Name,
			EventType:   Event.EventType,
			Description: Event.Description,
			Start:       Event.Start,
			End:         Event.End,
			Date:        Event.Date.Format("02-01-2006"),
			CreatedAt:   Event.CreatedAt.Format("02-01-2006"),
			UpdatedAt:   Event.UpdatedAt.Format("02-01-2006"),
			CreatedBy:   Event.CreatedBy,
			Invites:     invites,
		})
	}
	return result, count, nil
}

func (s eventService) Create(Event models.Event) (int, error) {

	startDate, err := time.Parse("02-01-2006", Event.StartDate)
	if err != nil {
		logs.Error(err)
		return 0, err
	}
	endDate, err := time.Parse("02-01-2006", Event.EndDate)
	if err != nil {
		logs.Error(err)
		return 0, err
	}
	if startDate.After(endDate) {
		endDate = startDate
	}
	var invites []entities.Employee
	for _, invite := range Event.Invites {
		invites = append(invites, entities.Employee{Model: gorm.Model{ID: invite.ID}}) // Assuming ID is provided for invites
	}
	// loop the date from start to end
	totalInserted := 0
	for date := startDate; date.Before(endDate.AddDate(0, 0, 1)); date = date.AddDate(0, 0, 1) {
		logs.Info(date.String())
		_, err := s.repository.Create(entities.Event{
			Name:        Event.Name,
			Description: Event.Description,
			EventType:   Event.EventType,
			Start:       Event.Start,
			End:         Event.End,
			Date:        date,
			CreatedBy:   Event.CreatedBy,
			Invites:     invites,
		})
		if err != nil {
			logs.Error(err)
			continue
		}
		totalInserted++
	}
	return totalInserted, nil
}

func (s eventService) Update(id uint, event models.Event) (*models.Event, error) {
	// Parse the date
	date, err := time.Parse("02-01-2006", event.Date)
	if err != nil {
		return nil, err
	}

	// Convert the invite list from models to entities
	var invites []entities.Employee
	for _, invite := range event.Invites {
		invites = append(invites, entities.Employee{Model: gorm.Model{ID: invite.ID}}) // Assuming ID is provided for invites
	}
	// Call the repository's update method
	result, err := s.repository.Update(id, entities.Event{
		Name:        event.Name,
		Description: event.Description,
		Start:       event.Start,
		End:         event.End,
		Date:        date,
		Invites:     invites,
	})
	if err != nil {
		return nil, err
	}

	// Map the updated event from entities to models
	var updatedInvites []models.Employee
	for _, invite := range result.Invites {
		updatedInvites = append(updatedInvites, models.Employee{
			ID:       invite.ID,
			Name:     invite.Name,
			Email:    invite.Email,
			Phone:    invite.Phone,
			Position: invite.Position,
		})
	}

	return &models.Event{
		ID:          result.ID,
		Name:        result.Name,
		Description: result.Description,
		Start:       result.Start,
		End:         result.End,
		Date:        result.Date.Format("02-01-2006"),
		Invites:     updatedInvites,
	}, nil
}

func (s eventService) CalendarEvent(userID, month, year int) ([]models.Event, error) {
	data, err := s.repository.CalendarEvent(month, year)
	if err != nil {
		return nil, err
	}
	var result []models.Event
	for _, Event := range data {
		if len(Event.Invites) > 0 {
			//check is user invited
			for _, invite := range Event.Invites {
				if invite.ID == uint(userID) {
					result = append(result, models.Event{
						ID:          Event.ID,
						Name:        Event.Name,
						EventType:   Event.EventType,
						Description: Event.Description,
						Start:       Event.Start,
						End:         Event.End,
						Date:        Event.Date.Format("02-01-2006"),
						CreatedAt:   Event.CreatedAt.Format("02-01-2006"),
						UpdatedAt:   Event.UpdatedAt.Format("02-01-2006"),
						CreatedBy:   Event.CreatedBy,
					})
				}
			}
		} else {
			result = append(result, models.Event{
				ID:          Event.ID,
				Name:        Event.Name,
				EventType:   Event.EventType,
				Description: Event.Description,
				Start:       Event.Start,
				End:         Event.End,
				Date:        Event.Date.Format("02-01-2006"),
				CreatedAt:   Event.CreatedAt.Format("02-01-2006"),
				UpdatedAt:   Event.UpdatedAt.Format("02-01-2006"),
				CreatedBy:   Event.CreatedBy,
			})
		}
	}
	return result, nil
}

func (s eventService) Delete(id uint) error {
	return s.repository.Delete(id)
}
