package models

import (
	"time"
)

type Role string

const (
	Employee Role = "employee"
	Employer Role = "employer"
)

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"-"`
	Role     Role   `json:"role"`
}

type Task struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `gorm:"default:Pending" json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	DueDate     time.Time `json:"due_date"`
	AssigneeID  uint      `json:"assignee_id"`
	CreatedBy   uint      `json:"created_by"`
}
