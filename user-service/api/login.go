package api

import (
	"fmt"
	"user/auth"
	"user/global"
	"user/models"
	"user/repositories"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var loginCredentials models.LoginCredentials
	userRepository := repositories.NewUserRepository(global.DB)

	if err := c.BindJSON(&loginCredentials); err != nil {
		fmt.Println("Invalid Request Body: ", err)
		c.IndentedJSON(403, gin.H{
			"error": "Invalid Request Body",
		})
		return
	}

	token, err := auth.Login(userRepository, loginCredentials.Username, loginCredentials.Password)
	if err != nil {
		c.IndentedJSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(200, gin.H{
		"access_token": token,
	})
}
