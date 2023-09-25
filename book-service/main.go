package main

import (
	"fmt"

	"book/api"
	"book/global"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Printf("Hello World\n")

	dsn := "host=postgres1 user=user password=pass dbname=books port=5432 sslmode=disable TimeZone=Asia/Shanghai"
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

	app.POST("/add", api.AddBookHandler)

	app.Run("0.0.0.0:8089")
}
