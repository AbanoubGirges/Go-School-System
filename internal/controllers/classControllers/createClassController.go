package controllers

import (
	"context"
	"time"

	dto "github.com/AbanoubGirges/Go-School-System/internal/dto/class"
	"github.com/AbanoubGirges/Go-School-System/internal/repo/classRepo"
	"github.com/gin-gonic/gin"
)
type ClassController struct {
	classRepo *classRepo.ClassRepo
	jwtSecret	string
}
func NewClassController(classRepo *classRepo.ClassRepo, jwtSecret string) *ClassController {
	return &ClassController{
		classRepo: classRepo,
		jwtSecret: jwtSecret,
	}
}
func (cc *ClassController) CreateClassController(c *gin.Context) {
	ctx ,cancel := context.WithTimeout(c.Request.Context(), 60*time.Second)
	defer cancel()
	var classRequest dto.CreateClassRequest
	if err := c.ShouldBindJSON(&classRequest); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	classModel := classRequest.BindToModel()
	err := cc.classRepo.CreateClass(classModel, ctx)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create class"})
		return
	}
	c.JSON(201, gin.H{"message": "Class created successfully"})

}