package controllers

import (
	"github.com/gin-gonic/gin"
	"context"
	"time"
	"strconv"
)
func (cc *ClassController) UpdateClassController(c *gin.Context) {
	ctx ,cancel := context.WithTimeout(c.Request.Context(), 60*time.Second)
	defer cancel()
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "INVALID_CLASS_ID"})
		return
	}
	 err = cc.classRepo.UpdateClass(uint(id),ctx)
	if err != nil {
		c.JSON(500, gin.H{"error": "FAILED_TO_UPDATE_CLASS"})
		return
	}
	c.JSON(200, gin.H{"message": "CLASS_UPDATED_SUCCESSFULLY"})
}