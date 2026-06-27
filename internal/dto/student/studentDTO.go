package dto

import (
	"time"

	models "github.com/AbanoubGirges/Go-School-System/internal/models/student"
	"github.com/google/uuid"
)

type CreateStudentRequest struct {
	Name        string    `json:"name" binding:"required"`
	Age         int       `json:"age" binding:"required"`
	ClassID     uuid.UUID `json:"class_id" binding:"required"`
	PhoneNumber string    `json:"phone_number" binding:"required"`
	Email       string    `json:"email" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	Coordinates string    `json:"coordinates" binding:"required"`
	Birthdate   string    `json:"birthdate" binding:"required"`
}

func (r *CreateStudentRequest) BindToModel() *models.Student {
	id := uuid.New()
	return &models.Student{
		ID:          id,
		Name:        r.Name,
		Age:         uint(r.Age),
		Class:       r.ClassID,
		PhoneNumber: r.PhoneNumber,
		Email:       r.Email,
		Location:    r.Location,
		Coordinates: r.Coordinates,
		Birthdate:   r.Birthdate,
	}
}

type CreateAttendanceRequest struct {
	StudentID string    `json:"studentId" binding:"required"`
	Date      time.Time `json:"date" binding:"required"`
	Present   bool      `json:"present"`
}