package main

import (
	"fmt"

	"rental/api"
	"rental/global"
	"rental/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Printf("Hello World\n")

	dsn := "host=postgres2 user=user password=pass dbname=rental port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Could not connect to  db")
		return
	}
	global.DB = db

	// Close the connection at the end
	connection, _ := db.DB()
	defer connection.Close()

	repositories.NewRentalRepository(db).EnsureInitialized()

	app := gin.Default()

	app.GET("/", api.HealthCheckHandler)
	app.POST("/rent", api.RentEntityHandler)
	app.GET("/list/all", api.GetUserRentalsHandler)
	app.GET("/list/active", api.GetActiveUserRentalsHandler)
	app.POST("/rent/complete", api.CompleteRental)

	app.Run("0.0.0.0:8083")

	global.BookEventWriter.Close()
}
