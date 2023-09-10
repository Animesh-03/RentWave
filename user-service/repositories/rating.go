package repositories

import (
	"user/models"

	"gorm.io/gorm"
)

type RatingRepository struct {
	Repository[models.Rating]
}

func NewRatingRepository(db *gorm.DB) *RatingRepository {
	repository := RatingRepository{}
	repository.DB = db
	return &repository
}
