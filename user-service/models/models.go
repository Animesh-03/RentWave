package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `gorm:"unique"`
	PasswordHash string
}

type RatingValue struct {
	Value  uint8
	Raters uint
}

type Rating struct {
	gorm.Model
	BookRating  RatingValue `gorm:"embedded"`
	BikeRating  RatingValue `gorm:"embedded"`
	HotelRating RatingValue `gorm:"embedded"`
}

type RegisterCredentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginCredentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
