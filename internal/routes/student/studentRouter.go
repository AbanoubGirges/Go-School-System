package routes

import (
	studentControllers "github.com/AbanoubGirges/Go-School-System/internal/controllers/studentControllers"
	"github.com/AbanoubGirges/Go-School-System/internal/middleware"
	"github.com/gin-gonic/gin"
)
func StudentRouter(rg *gin.RouterGroup, userMiddleware *middleware.UserMiddleware, studentController *studentControllers.StudentController) {
	rg.POST("/students", userMiddleware.Handler(), studentController.CreateStudentController)
	rg.GET("/students", userMiddleware.AdminHandler(), studentController.GetAllStudentsController)
	rg.GET("/students/:id", userMiddleware.AdminHandler(), studentController.GetStudentByIdController)
	rg.POST("/students/attendance", userMiddleware.AdminHandler(), studentController.CreateAttendanceController)
	rg.GET("/students/:id/attendance", userMiddleware.AdminHandler(), studentController.GetAttendanceController)
}