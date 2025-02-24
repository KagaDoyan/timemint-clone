package repositories

import (
	"fmt"
	"go-fiber/domain/entities"
	"time"

	"gorm.io/gorm"
)

type EventRepository interface {
	FindAll(page, limit int) ([]entities.Event, int64, error)
	Create(Event entities.Event) (*entities.Event, error)
	Update(id uint, Event entities.Event) (*entities.Event, error)
	Delete(id uint) error
	IsEvent(date time.Time) (bool, error)
	CalendarEvent(month, year int) ([]entities.Event, error)
}

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) EventRepository {
	db.AutoMigrate(&entities.Event{})
	return &eventRepository{db: db}
}

func (r eventRepository) FindAll(page, limit int) ([]entities.Event, int64, error) {
	var Events []entities.Event
	var count int64
	err := r.db.Find(&Events).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Apply pagination
	err = r.db.Limit(limit).Offset(offset).Order("created_at desc").Order("created_at desc").Preload("Invites").Find(&Events).Error
	if err != nil {
		return nil, 0, err
	}
	return Events, count, nil
}

func (r eventRepository) Create(Event entities.Event) (*entities.Event, error) {
	err := r.db.Create(&Event).Error
	if err != nil {
		return nil, err
	}
	var result entities.Event
	err = r.db.First(&result, Event.ID).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r eventRepository) Update(id uint, updatedEvent entities.Event) (*entities.Event, error) {
	// Start a transaction to ensure data consistency
	tx := r.db.Begin()

	// Update the basic event details
	err := tx.Model(&entities.Event{}).Where("id = ?", id).Updates(map[string]interface{}{
		"name":        updatedEvent.Name,
		"description": updatedEvent.Description,
		"start":       updatedEvent.Start,
		"end":         updatedEvent.End,
		"date":        updatedEvent.Date,
		"created_by":  updatedEvent.CreatedBy,
	}).Error
	if err != nil {
		tx.Rollback() // Rollback in case of error
		return nil, err
	}

	// Update the invites if provided
	if len(updatedEvent.Invites) > 0 {
		// Clear existing invites
		err = tx.Table("event_invites").Where("event_id = ?", id).Delete(nil).Error
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		// Add new invites
		for _, invite := range updatedEvent.Invites {
			err = tx.Table("event_invites").Create(&map[string]interface{}{"event_id": id, "employee_id": invite.ID}).Error
			if err != nil {
				tx.Rollback()
				return nil, fmt.Errorf("error creating invite: %v", err)
			}
		}
	}

	// Fetch and return the updated event
	var result entities.Event
	err = tx.Preload("Invites").First(&result, id).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit() // Commit the transaction
	return &result, nil
}

func (r eventRepository) Delete(id uint) error {
	err := r.db.Unscoped().Delete(&entities.Event{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r eventRepository) IsEvent(date time.Time) (bool, error) {
	var Event entities.Event
	//check if date between start and end date
	err := r.db.Where("start_date <= ? AND end_date >= ?", date, date).First(&Event).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r eventRepository) CalendarEvent(month, year int) ([]entities.Event, error) {
	var events []entities.Event
	err := r.db.Preload("Invites").Where("MONTH(date) = ? AND YEAR(date) = ?", month, year).Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}
