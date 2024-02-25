package model

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
	Username  string     `gorm:"uniqueIndex;not null"`
	Password  string     `gorm:"not null"`
	Email     string     `gorm:"uniqueIndex;not null"`
	Nickname  string
	HeadImg   string
	Mobile    string `gorm:"uniqueIndex;not null"`
}
