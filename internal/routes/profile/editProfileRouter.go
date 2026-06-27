package routes

import (
	controllers "github.com/AbanoubGirges/Go-School-System/internal/controllers/profileControllers"
	"github.com/AbanoubGirges/Go-School-System/internal/middleware"
	"github.com/gin-gonic/gin"
)

func EditProfileRouter(rg *gin.RouterGroup, userMiddleware *middleware.UserMiddleware, editProfileController *controllers.EditProfileController) {
	rg.PATCH("/profile", userMiddleware.Handler(), editProfileController.EditProfile)
}
