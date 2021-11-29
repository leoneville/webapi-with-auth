package controllers

import (
	"webapi-with-go/database"
	"webapi-with-go/models"
	"webapi-with-go/services"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	db := database.GetDatabase()

	var login models.Login
	err := c.ShouldBindJSON(&login)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	var user models.User
	dbError := db.Where("email = ?", login.Email).First(&user).Error
	if dbError != nil {
		c.JSON(400, gin.H{
			"error": "cannot find user",
		})
		return
	}

	if user.Password != services.SHA256Encoder(login.Password) {
		c.JSON(401, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	token, err := services.NewJWTService().GenerateToken(user.ID)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}
