package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Isbn    string `gorm:"unique"`
	Title   string
	Author  string
	Details []BookDetails
}

type BookDetails struct {
	gorm.Model
	BookID      uint
	Book        Book
	OwnerID     uint
	Image       string
	Active      bool `gorm:"default:true"`
	Rented      bool `gorm:"default:false"`
	Description string
	Price       uint
}
