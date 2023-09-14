package repositories

import (
	"book/models"

	"gorm.io/gorm"
)

type BookRepository struct {
	Repository[models.Book]
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	repository := BookRepository{}
	repository.DB = db
	return &repository
}
