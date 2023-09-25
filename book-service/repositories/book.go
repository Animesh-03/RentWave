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
	repository.EnsureInitialized()
	return &repository
}

func (b *BookRepository) GetByIsbn(isbn string) (*models.Book, error) {
	var book models.Book
	result := b.DB.Where("isbn = ?", isbn).First(&book)
	return &book, result.Error
}
