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
	OwnerID     uint
	Image       string
	Active      bool `gorm:"default:true"`
	Description string
	Price       uint
}
