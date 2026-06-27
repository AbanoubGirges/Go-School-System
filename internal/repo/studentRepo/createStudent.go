package studentRepo

import (
	"context"

	models "github.com/AbanoubGirges/Go-School-System/internal/models/student"
	"github.com/google/uuid"
)

func (s *StudentRepo) CreateStudent(student *models.Student, ctx context.Context) error {
	return s.repo.DB.WithContext(ctx).Create(student).Error
}

func (s *StudentRepo) GetStudentsByClass(classID uuid.UUID, ctx context.Context) ([]models.Student, error) {
	var students []models.Student
	err := s.repo.DB.WithContext(ctx).Where("class = ?", classID).Find(&students).Error
	return students, err
}

func (s *StudentRepo) GetStudentByID(studentID uuid.UUID, ctx context.Context) (*models.Student, error) {
	var student models.Student
	err := s.repo.DB.WithContext(ctx).First(&student, "id = ?", studentID).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (s *StudentRepo) CreateAttendance(attendance *models.Attendance, ctx context.Context) error {
	return s.repo.DB.WithContext(ctx).Create(attendance).Error
}

func (s *StudentRepo) GetAttendanceByStudent(studentID uuid.UUID, ctx context.Context) ([]models.Attendance, error) {
	var attendance []models.Attendance
	err := s.repo.DB.WithContext(ctx).Where("student_id = ?", studentID).Order("date desc").Find(&attendance).Error
	return attendance, err
}