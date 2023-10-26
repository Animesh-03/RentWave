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

func (b *BookRepository) GetByName(name string) []models.Book {
	var books []models.Book
	b.DB.Where("LOWER(title) like %LOWER(?)%", name).Find(&books)
	return books
}
