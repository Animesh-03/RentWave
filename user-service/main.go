package main

import (
	"fmt"
	"user/api"
	"user/global"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=postgres user=user password=pass dbname=users port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Could not connect to  db")
		return
	}
	global.DB = db

	// Close the connection at the end
	connection, _ := db.DB()
	defer connection.Close()

	app := gin.Default()

	app.POST("/register", api.RegisterHandler)
	app.POST("/login", api.LoginHandler)

	app.Run("0.0.0.0:8087")
}
