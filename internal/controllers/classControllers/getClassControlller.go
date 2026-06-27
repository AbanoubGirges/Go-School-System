package controllers
import (
	"github.com/gin-gonic/gin"
	"context"
	"time"
	"strconv"
	"github.com/golang-jwt/jwt/v5"
)
func (cc *ClassController) GetClassControllerByID(c *gin.Context) {
	ctx ,cancel := context.WithTimeout(c.Request.Context(), 60*time.Second)
	defer cancel()
	classId := c.Param("id")
	classIdInt, err := strconv.Atoi(classId)
	if err != nil {
		c.JSON(400, gin.H{"error": "INVALID_CLASS_ID"})
		return
	}
	classModel, err := cc.classRepo.GetClassById(uint(classIdInt), ctx)
	if err != nil {
		c.JSON(404, gin.H{"error": "CLASS_NOT_FOUND"})
		return
	}
	c.JSON(200, classModel)
}
func (cc *ClassController) GetClassController(c *gin.Context) {
	ctx ,cancel := context.WithTimeout(c.Request.Context(), 60*time.Second)
	defer cancel()
	var classId uint
	claimsValue, exists := c.Get("claims")
	if !exists {
		c.JSON(500, gin.H{"error": "claims not found"})
		return
	}
	claims, ok := claimsValue.(jwt.MapClaims)
if !ok {
    c.JSON(500, gin.H{"error": "invalid claims"})
    return
}
	if classIdParam,exists := claims["id"]; classIdParam != "" && exists {
		id, err := strconv.ParseUint(classIdParam.(string), 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "INVALID_CLASS_ID"})
			return
		}
		classId = uint(id)
	}
	classModel, err := cc.classRepo.GetClassById(classId, ctx)
	if err != nil {
		c.JSON(404, gin.H{"error": "CLASS_NOT_FOUND"})
		return
	}
	c.JSON(200, classModel)
}