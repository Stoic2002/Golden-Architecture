package entity

import (
	"time"
)

// User represents a user entity
type User struct {
	ID        uint      `gorm:"primaryKey"`
	Email     string    `gorm:"size:255;uniqueIndex;not null"`
	Password  string    `gorm:"size:255;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// TableName specifies the table name for User
func (User) TableName() string {
	return "users"
}
