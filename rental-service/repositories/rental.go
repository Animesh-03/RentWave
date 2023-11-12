package repositories

import (
	"rental/models"
	"time"

	"gorm.io/gorm"
)

type RentalRepository struct {
	Repository[models.Rental]
}

func NewRentalRepository(db *gorm.DB) *RentalRepository {
	repository := &RentalRepository{}
	repository.DB = db
	repository.EnsureInitialized()

	return repository
}

func (r *RentalRepository) GetAllUserRentals(uid uint) (rentals []*models.Rental) {
	r.DB.Where("lessor = ?", uid).Find(rentals)

	return rentals
}

func (r *RentalRepository) GetActiveUserRentals(uid uint) (rentals []*models.Rental) {
	now := time.Now().UnixNano()
	r.DB.Where("lessor = ? AND from > ? AND to < ?", uid, now, now).Find(rentals)

	return rentals
}
