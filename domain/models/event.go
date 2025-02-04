package models

type Event struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	EventType   string `json:"event_type"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Start       string `json:"start"`
	End         string `json:"end"`
	Date        string `json:"date"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	CreatedBy   uint   `json:"created_by"`
}
