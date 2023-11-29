package repositories

import (
	"rental/models"

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
	r.DB.Where("lessor = ?", uid).Find(&rentals)

	return rentals
}

func (r *RentalRepository) GetActiveUserRentals(uid uint) (rentals []*models.Rental) {
	r.DB.Where("lessor = ? AND status = ?", uid, "ongoing").Find(&rentals)

	return rentals
}

func (r *RentalRepository) SetStatus(id uint, status models.RentalStatus) (*models.Rental, error) {
	rental, err := r.GetByID(id)
	if err != nil {
		return nil, err
	}

	rental.Status = status
	r.DB.Save(rental)
	return rental, nil
}
