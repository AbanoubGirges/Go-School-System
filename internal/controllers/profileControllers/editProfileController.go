package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/AbanoubGirges/Go-School-System/internal/repo/userRepo"
	"context"
	"time"
)
type EditProfileController struct{
	UserRepo	*userRepo.UserRepo
	jwtSecret	string
}
func NewEditProfileController(userRepo *userRepo.UserRepo,jwtSecret string) *EditProfileController{
	return &EditProfileController{userRepo,jwtSecret}
}
func (e *EditProfileController) EditProfile(c *gin.Context) {
	ctx ,cancel := context.WithTimeout(c.Request.Context(), 60*time.Second)
	defer cancel()
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(401, gin.H{"error": "CLAIMS_NOT_FOUND_IN_CONTEXT"})
		return
	}
	claimsMap, ok := claims.(map[string]interface{})
	if !ok {
		c.JSON(500, gin.H{"error": "INVALID_CLAIMS_FORMAT"})
		return
	}
	id, exists := claimsMap["id"]
	if !exists {
		c.JSON(401, gin.H{"error": "USER_ID_NOT_FOUND_IN_CONTEXT"})
		return
	}
	var fieldsToUpdate map[string]interface{}
	if err := c.ShouldBindJSON(&fieldsToUpdate); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	if err := e.UserRepo.EditUserById(id.(string), fieldsToUpdate, ctx); err != nil {
		c.JSON(500, gin.H{"error": "Failed to edit user profile"})
		return
	}
	c.JSON(200, gin.H{"message": "Profile updated successfully"})
}