package controllers

import (
	"context"
	"net/http"
	"time"

	dto "github.com/AbanoubGirges/Go-School-System/internal/dto/auth"
	"github.com/AbanoubGirges/Go-School-System/internal/repo/userRepo"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)
type AuthController struct{
	UserRepo	*userRepo.UserRepo
	jwtSecret	string
}
func NewAuthController(userRepo *userRepo.UserRepo,jwtSecret string) *AuthController{
	return &AuthController{userRepo,jwtSecret}
}
func(a *AuthController) RegisterController(c *gin.Context){
	ctx, cancel := context.WithTimeout(c.Request.Context(), 60*time.Second)
    defer cancel()

	var request dto.RegisterRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest,struct{Error string}{Error:"FAILED_TO_PARSE_TO_JSON"})
		return
	}
	file, err := c.FormFile("image")
    if err != nil {
        c.JSON(400, gin.H{"error": "IMG_REQUIRED"})
        return
    }
	err = c.SaveUploadedFile(file, "./uploads/"+file.Filename)
    if err != nil {
        c.JSON(500, gin.H{"error": "FAILED_TO_UPLOAD_IMAGE"})
        return
    }

	hashedPassword, err := bcrypt.GenerateFromPassword(
        []byte(request.Password),
        bcrypt.DefaultCost,
    )
    if err != nil {
        c.JSON(http.StatusInternalServerError,struct{Error string}{Error: "FAILED_TO_CREATE_USER"})
		return
    }
	request.Password=string(hashedPassword)
	userModel:=	request.BindToModel()
	err=a.UserRepo.CreateUser(userModel,ctx)
	c.IndentedJSON(http.StatusCreated,struct{Message string}{Message: "USER_CREATED_SUCCESSFULLY"})	
}