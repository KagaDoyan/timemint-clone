package models

type DayOfWork struct {
	ID        uint   `json:"id"`
	Day       string `json:"day"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	IsWorkDay bool   `json:"is_work_day"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}
