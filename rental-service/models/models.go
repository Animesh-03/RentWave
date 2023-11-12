package models

import "gorm.io/gorm"

type RentalType string

const (
	BOOK_RENTAL    RentalType = "book"
	VEHICLE_RENTAL RentalType = "vehicle"
)

type Rental struct {
	gorm.Model
	Type     RentalType
	EntityID uint
	Lessor   uint
	From     uint64
	To       uint64
}
