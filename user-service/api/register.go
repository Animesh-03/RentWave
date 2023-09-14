package api

import (
	"fmt"
	"user/auth"
	"user/global"
	"user/models"
	"user/repositories"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	var registerCredentials models.RegisterCredentials
	userRepository := repositories.NewUserRepository(global.DB)

	if err := c.BindJSON(&registerCredentials); err != nil {
		fmt.Println("Invalid Request Body: ", err)
		c.IndentedJSON(403, gin.H{
			"error": "Invalid Request Body",
		})
		return
	}

	err := auth.Register(userRepository, registerCredentials.Username, registerCredentials.Password)
	if err != nil {
		fmt.Println("Error registering user: ", err)
		c.IndentedJSON(500, gin.H{
			"error": "Username already exists",
		})
		return
	}

	c.IndentedJSON(200, "OK")
}
