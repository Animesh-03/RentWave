package main

import (
	"fmt"
	"user/auth"
	"user/repositories"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Printf("Hello World\n")

	dsn := "host=127.0.0.1 user=user password=pass dbname=tbd port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Could not connect to  db")
		return
	}

	// Close the connection at the end
	connection, _ := db.DB()
	defer connection.Close()

	// Initialize the repositories
	userRepository := repositories.NewUserRepository(db)
	ratingRepository := repositories.NewRatingRepository(db)

	// Ensure the tables exist in the database
	userRepository.EnsureInitialized()
	ratingRepository.EnsureInitialized()

	auth.Register(userRepository, "Test", "pass")

	fmt.Println(auth.Login(userRepository, "Test", "pass"))

	fmt.Println("Tables Created")
}
