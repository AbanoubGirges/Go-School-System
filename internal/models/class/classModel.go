package models

import (
	"github.com/google/uuid"
)
type Class struct {
	ID          uuid.UUID	`gorm:"primaryKey"`
	Name    string
	Grade   string
}
type Class_User struct {
	ID          uuid.UUID	`gorm:"primaryKey"`
	ClassID     uuid.UUID	`gorm:"foreignKey:ClassID"`
	UserID      uuid.UUID	`gorm:"foreignKey:UserID"`
}