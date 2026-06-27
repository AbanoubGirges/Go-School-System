package models

import (
	"time"

	"github.com/google/uuid"
)

type Student struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	Name        string
	PhoneNumber string `gorm:"uniqueIndex"`
	Email       string `gorm:"uniqueIndex"`
	Class       uuid.UUID `gorm:"foreignKey:Class"`
	Location    string
	Coordinates string
	Age         uint
	Birthdate   string
}

type Attendance struct {
	ID         uuid.UUID `gorm:"primaryKey"`
	StudentID  uuid.UUID `gorm:"not null"`
	Student    Student   `gorm:"foreignKey:StudentID"`
	Date       time.Time `gorm:"not null"`
	Present    bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}