package routes

import (
	"golang_basic_gin/auth"
	"golang_basic_gin/config"
	"golang_basic_gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"error":   err.Error(),
		})
		c.Abort()
		return
	}

	// hash password
	err := user.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed hash password",
			"error":   err.Error(),
		})
		c.Abort()
		return
	}

	// insert user to DB
	insert := config.DB.Create(&user)
	if insert.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
			"error":   insert.Error.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user_id":  user.ID,
		"email":    user.Email,
		"username": user.Username,
	})
	return
}

func GenerateToken(c *gin.Context) {
	request := models.TokenRequest{}
	user := models.User{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"error":   err.Error(),
		})
		c.Abort()
		return
	}
	// check email
	checkEmail := config.DB.Where("email = ?", request.Email).First(&user)
	if checkEmail.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Email not found",
			"error":   checkEmail.Error.Error(),
		})
		c.Abort()
		return
	}
	// check password
	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Password not match",
			"error":   credentialError.Error(),
		})
		c.Abort()
		return
	}
	// generate token
	tokenString, err := auth.GenerateJWT(user.Email, user.Username, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to generate token",
			"error":   err.Error(),
		})
		c.Abort()
		return
	}
	// response token
	c.JSON(http.StatusCreated, gin.H{
		"token": tokenString,
	})
	return
}
