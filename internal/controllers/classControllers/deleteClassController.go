package controllers
import (
	"github.com/gin-gonic/gin"
	"context"
	"time"
	"strconv"
)
func (cc *ClassController) DeleteClassController(c *gin.Context) {
	ctx ,cancel := context.WithTimeout(c.Request.Context(), 60*time.Second)
	defer cancel()
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid class ID"})
		return
	}
	err = cc.classRepo.DeleteClass(uint(idUint),ctx)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete class"})
		return
	}
	c.JSON(200, gin.H{"message": "Class deleted successfully"})
}