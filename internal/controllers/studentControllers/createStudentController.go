package controllers

import (
	"context"
	"net/http"
	"time"

	dto "github.com/AbanoubGirges/Go-School-System/internal/dto/student"
	models "github.com/AbanoubGirges/Go-School-System/internal/models/student"
	"github.com/AbanoubGirges/Go-School-System/internal/repo/studentRepo"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type StudentController struct {
	studentRepo *studentRepo.StudentRepo
	jwtSecret   string
}

func NewStudentController(studentRepo *studentRepo.StudentRepo, jwtSecret string) *StudentController {
	return &StudentController{
		studentRepo: studentRepo,
		jwtSecret:   jwtSecret,
	}
}

func (sc *StudentController) CreateStudentController(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 60*time.Second)
	defer cancel()

	var studentDTO dto.CreateStudentRequest
	err := c.ShouldBindJSON(&studentDTO)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, struct{ Error string }{Error: "FAILED_TO_PARSE_TO_JSON"})
		return
	}

	student := studentDTO.BindToModel()
	err = sc.studentRepo.CreateStudent(student, ctx)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, struct{ Error string }{Error: "FAILED_TO_CREATE_STUDENT"})
		return
	}

	c.IndentedJSON(http.StatusCreated, struct{ Message string }{Message: "STUDENT_CREATED_SUCCESSFULLY"})
}

func (sc *StudentController) GetAllStudentsController(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 60*time.Second)
	defer cancel()

	classIDParam := c.Query("classId")
	if classIDParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "classId query parameter is required"})
		return
	}

	classID, err := uuid.Parse(classIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid classId"})
		return
	}

	students, err := sc.studentRepo.GetStudentsByClass(classID, ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "FAILED_TO_FETCH_STUDENTS"})
		return
	}

	c.JSON(http.StatusOK, students)
}

func (sc *StudentController) GetStudentByIdController(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 60*time.Second)
	defer cancel()

	studentIDParam := c.Param("id")
	studentID, err := uuid.Parse(studentIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid student id"})
		return
	}

	student, err := sc.studentRepo.GetStudentByID(studentID, ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "STUDENT_NOT_FOUND"})
		return
	}

	c.JSON(http.StatusOK, student)
}

func (sc *StudentController) CreateAttendanceController(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 60*time.Second)
	defer cancel()

	var attendanceDTO dto.CreateAttendanceRequest
	if err := c.ShouldBindJSON(&attendanceDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "FAILED_TO_PARSE_TO_JSON"})
		return
	}

	studentID, err := uuid.Parse(attendanceDTO.StudentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid studentId"})
		return
	}

	attendance := &models.Attendance{
		ID:        uuid.New(),
		StudentID: studentID,
		Date:      attendanceDTO.Date,
		Present:   attendanceDTO.Present,
	}

	if err := sc.studentRepo.CreateAttendance(attendance, ctx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "FAILED_TO_CREATE_ATTENDANCE"})
		return
	}

	c.JSON(http.StatusCreated, attendance)
}

func (sc *StudentController) GetAttendanceController(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 60*time.Second)
	defer cancel()

	studentIDParam := c.Param("id")
	studentID, err := uuid.Parse(studentIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid student id"})
		return
	}

	attendance, err := sc.studentRepo.GetAttendanceByStudent(studentID, ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "FAILED_TO_FETCH_ATTENDANCE"})
		return
	}

	c.JSON(http.StatusOK, attendance)
}
