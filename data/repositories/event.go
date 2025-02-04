package repositories

import (
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
	err = r.db.Limit(limit).Offset(offset).Order("created_at desc").Order("created_at desc").Find(&Events).Error
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

func (r eventRepository) Update(id uint, Event entities.Event) (*entities.Event, error) {
	err := r.db.Model(&entities.Event{}).Where("id = ?", id).Updates(Event).Error
	if err != nil {
		return nil, err
	}
	var result entities.Event
	err = r.db.First(&result, id).Error
	if err != nil {
		return nil, err
	}
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
	err := r.db.Where("MONTH(date) = ? AND YEAR(date) = ?", month, year).Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}
