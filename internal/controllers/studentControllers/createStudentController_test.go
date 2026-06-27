package controllers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/AbanoubGirges/Go-School-System/internal/config"
	studentModels "github.com/AbanoubGirges/Go-School-System/internal/models/student"
	"github.com/AbanoubGirges/Go-School-System/internal/repo/studentRepo"
	"github.com/gin-gonic/gin"
	sqliteDriver "github.com/glebarez/sqlite"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func newTestStudentRepo(t *testing.T) *studentRepo.StudentRepo {
	t.Helper()

	db, err := gorm.Open(sqliteDriver.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open sqlite: %v", err)
	}
	if err := db.AutoMigrate(&studentModels.Student{}, &studentModels.Attendance{}); err != nil {
		t.Fatalf("failed to migrate schema: %v", err)
	}

	cfgRepo := &config.Repo{DB: db}
	return studentRepo.NewStudentRepo(cfgRepo)
}

func TestCreateAttendanceController_PersistsAttendance(t *testing.T) {
	gin.SetMode(gin.TestMode)
	repo := newTestStudentRepo(t)
	controller := NewStudentController(repo, "test-secret")

	studentID := uuid.New()
	payload := `{"studentId":"` + studentID.String() + `","date":"2026-06-27T00:00:00Z","present":true}`

	req := httptest.NewRequest(http.MethodPost, "/students/attendance", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	controller.CreateAttendanceController(c)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, w.Code)
	}

	attendance, err := repo.GetAttendanceByStudent(studentID, context.Background())
	if err != nil {
		t.Fatalf("expected attendance lookup to succeed: %v", err)
	}
	if len(attendance) != 1 {
		t.Fatalf("expected exactly one attendance row, got %d", len(attendance))
	}
	if attendance[0].StudentID != studentID {
		t.Fatalf("expected attendance for student %s, got %s", studentID, attendance[0].StudentID)
	}
	if !attendance[0].Present {
		t.Fatalf("expected attendance to be marked present")
	}
}

func TestGetAttendanceController_ReturnsStudentAttendance(t *testing.T) {
	gin.SetMode(gin.TestMode)
	repo := newTestStudentRepo(t)
	controller := NewStudentController(repo, "test-secret")

	studentID := uuid.New()
	attendance := &studentModels.Attendance{ID: uuid.New(), StudentID: studentID, Date: time.Now().UTC(), Present: true}
	if err := repo.CreateAttendance(attendance, context.Background()); err != nil {
		t.Fatalf("failed to seed attendance: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/students/"+studentID.String()+"/attendance", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{{Key: "id", Value: studentID.String()}}

	controller.GetAttendanceController(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, w.Code)
	}
}
