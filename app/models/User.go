package models

import (
	"time"
)

// User = user
type User struct {
	ID        uint      `gorm:"primary_key;auto_increment" json:"id"`
	Fristname string    `gorm:"type:varchar(255);" json:"fristname"`
	Lastname  string    `gorm:"type:varchar(255);" json:"lastname"`
	Nickname  string    `gorm:"type:varchar(255);" json:"nickname"`
	Email     string    `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password  string    `gorm:"type:varchar(20);" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
