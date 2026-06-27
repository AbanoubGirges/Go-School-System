package controllers

import (
	"github.com/gin-gonic/gin"
	"context"
	"time"
)
func (e *EditProfileController) DeleteProfile(c *gin.Context) {
	ctx ,cancel := context.WithTimeout(c.Request.Context(), 60*time.Second)
	defer cancel()
	claimsValue, exists := c.Get("claims")
	if !exists {
		c.JSON(401, gin.H{"error": "CLAIMS_NOT_FOUND_IN_CONTEXT"})
		return
	}
	claims, ok := claimsValue.(map[string]interface{})
	if !ok {
		c.JSON(500, gin.H{"error": "INVALID_CLAIMS_FORMAT"})
		return
	}
	id, exists := claims["id"]
	if !exists {
		c.JSON(401, gin.H{"error": "USER_ID_NOT_FOUND_IN_CONTEXT"})
		return
	}
	if err := e.UserRepo.DeleteUserById(id.(string), ctx); err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete user profile"})
		return
	}
	c.JSON(200, gin.H{"message": "Profile deleted successfully"})

}