package routes

import (
	authControllers"github.com/AbanoubGirges/Go-School-System/internal/controllers/authControllers"
	profileControllers "github.com/AbanoubGirges/Go-School-System/internal/controllers/profileControllers"
	middleware "github.com/AbanoubGirges/Go-School-System/internal/middleware"
	authRoutes "github.com/AbanoubGirges/Go-School-System/internal/routes/auth"
	profileRoutes "github.com/AbanoubGirges/Go-School-System/internal/routes/profile"
	"github.com/gin-gonic/gin"
	classControllers "github.com/AbanoubGirges/Go-School-System/internal/controllers/classControllers"
	classRoutes "github.com/AbanoubGirges/Go-School-System/internal/routes/class"
	studentRoutes "github.com/AbanoubGirges/Go-School-System/internal/routes/student"
	studentControllers "github.com/AbanoubGirges/Go-School-System/internal/controllers/studentControllers"
)
func RouterSetup(router *gin.Engine,userMiddleware *middleware.UserMiddleware,authController *authControllers.AuthController,editProfileController *profileControllers.EditProfileController,classController *classControllers.ClassController,studentController *studentControllers.StudentController) {
	api:=router.Group("/api/v1")
	authRoutes.AuthRouter(api,authController)
	profileRoutes.EditProfileRouter(api,userMiddleware,editProfileController)
	classRoutes.ClassRouter(api,userMiddleware,classController)
	studentRoutes.StudentRouter(api,userMiddleware,studentController)
}