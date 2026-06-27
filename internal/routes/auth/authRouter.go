package authRoutes

import (
	controllers "github.com/AbanoubGirges/Go-School-System/internal/controllers/authControllers"
	"github.com/gin-gonic/gin"
)
func AuthRouter(rg *gin.RouterGroup,authController *controllers.AuthController){
	rg.POST("/register",authController.RegisterController)
	rg.POST("/login", authController.LoginController)
}