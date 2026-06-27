package controllers

import (
	"github.com/gin-gonic/gin"
	"context"
	"time"
)
func (cc *ClassController) GetAllClassesController(c *gin.Context) {
	ctx ,cancel := context.WithTimeout(c.Request.Context(), 60*time.Second)
	defer cancel()
	classes, err := cc.classRepo.GetAllClasses(ctx)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve classes"})
		return
	}
	c.JSON(200, classes)
}