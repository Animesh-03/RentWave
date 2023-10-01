package main

import (
	"fmt"

	"book/api"
	"book/global"
	"book/repositories"

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

	repositories.NewBookRepository(db).EnsureInitialized()
	repositories.NewBookDetailsRepository(db).EnsureInitialized()

	app := gin.Default()

	app.POST("/add", api.AddBookHandler)
	app.GET("/user_listings", api.GetUserListings)
	app.PATCH("/toggle_active", api.ToggleListing)
	app.GET("/active_listings", api.GetActiveListings)

	app.Run("0.0.0.0:8089")
}
