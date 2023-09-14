package repositories

import (
	"book/models"

	"gorm.io/gorm"
)

type BookDetailsRepository struct {
	Repository[models.BookDetails]
}

func NewBookDetailsRepository(db *gorm.DB) *BookDetailsRepository {
	repository := BookDetailsRepository{}
	repository.DB = db
	return &repository
}
