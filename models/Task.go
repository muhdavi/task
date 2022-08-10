package models

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	ID        uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Content   string         `json:"content"`
	Person    string         `json:"person"`
	DueDate   time.Time      `json:"due_date"`
	IsDone    bool           `json:"is_done"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
