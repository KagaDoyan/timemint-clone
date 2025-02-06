package models

type ShiftAssignment struct {
	ID            uint     `json:"id"`
	EmployeeID    uint     `json:"employee_id"`
	Employee      Employee `json:"employee"`
	ShiftID       uint     `json:"shift_id"`
	Shift         Shift    `json:"shift"`
	Date          string   `json:"date"`
	CreatedBy     uint     `json:"created_by"`
	CreatedByUser Employee `json:"created_by_user"`
}
