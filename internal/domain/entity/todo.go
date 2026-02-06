package entity

import (
	"time"
)

// Todo represents a todo item entity
type Todo struct {
	ID          uint      `gorm:"primaryKey"`
	Title       string    `gorm:"size:255;not null"`
	Description string    `gorm:"type:text"`
	Completed   bool      `gorm:"default:false"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

// TableName specifies the table name for Todo
func (Todo) TableName() string {
	return "todos"
}
