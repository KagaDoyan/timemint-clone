package models

import "time"

type Employee struct {
	ID         uint      `json:"id"`
	EmployeeID string    `json:"employee_id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Address    string    `json:"address"`
	Position   string    `json:"position"`
	RoleID     uint      `json:"role_id"`
	Role       Role      `json:"role"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
