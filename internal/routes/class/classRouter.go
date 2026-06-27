package routes
import (
	classControllers "github.com/AbanoubGirges/Go-School-System/internal/controllers/classControllers"
	middleware "github.com/AbanoubGirges/Go-School-System/internal/middleware"
	"github.com/gin-gonic/gin"
)
func ClassRouter(rg *gin.RouterGroup, userMiddleware *middleware.UserMiddleware, classController *classControllers.ClassController) {
	rg.POST("/classes", userMiddleware.AdminHandler(), classController.CreateClassController)
	rg.PUT("/classes/:id", userMiddleware.AdminHandler(), classController.UpdateClassController)
	rg.GET("/classes", userMiddleware.AdminHandler(), classController.GetAllClassesController)
	rg.GET("/classes/:id", userMiddleware.AdminHandler(), classController.GetClassControllerByID)
	rg.DELETE("/classes/:id", userMiddleware.AdminHandler(), classController.DeleteClassController)
	rg.GET("/classes",userMiddleware.Handler(), classController.GetClassController)
}