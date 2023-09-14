package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	isbn    string `gorm:"unique"`
	name    string
	author  string
	details []BookDetails
}

type BookDetails struct {
	gorm.Model
	BookID  uint
	ownerID uint
	image   string
}
