package controllers

import (
	"context"
	"time"

	dto "github.com/AbanoubGirges/Go-School-System/internal/dto/auth"
	models "github.com/AbanoubGirges/Go-School-System/internal/models/user"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)
func (a *AuthController) LoginController(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 60*time.Second)
	defer cancel()
	var loginRequest dto.LoginRequest
	err := c.ShouldBindJSON(&loginRequest)
	if err != nil {
		c.JSON(400, struct{ Error string }{Error: "FAILED_TO_PARSE_TO_JSON"})
		return
	}

	user, err := a.UserRepo.GetUserByPhone(loginRequest.PhoneNumber, ctx)
	if err != nil {
		c.JSON(404, struct{ Error string }{Error: "USER_NOT_FOUND"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		c.JSON(401, struct{ Error string }{Error: "INVALID_CREDENTIALS"})
		return
	}
	var userToken *models.UserToken
	userToken = user.BindToToken()
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userToken.ID,
		"role":    userToken.Role,
		"class": userToken.Class,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString, err := jwtToken.SignedString([]byte(a.jwtSecret))
	if err != nil {
		c.JSON(500, struct{ Error string }{Error: "FAILED_TO_GENERATE_TOKEN"})
		return
	}
	c.JSON(200, struct {

		User    *models.UserToken
		Token   string
	}{
		User:    userToken,
		Token:   tokenString,
	})

}