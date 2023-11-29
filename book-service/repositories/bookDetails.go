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
	repository.EnsureInitialized()
	return &repository
}

func (b *BookDetailsRepository) GetUserListings(uid uint) []models.Book {
	var bookDetails []models.Book
	subQuery := b.DB.Table("book_details").Select("book_id").Where("owner_id = ?", uid)
	b.DB.Debug().Preload("Details").Where("id in (?)", subQuery).Find(&bookDetails)
	return bookDetails
}

func (b *BookDetailsRepository) ToggleListing(id uint) error {
	bookDetails, err := b.GetByID(id)
	if err != nil {
		return err
	}

	bookDetails.Active = !bookDetails.Active
	b.DB.Save(bookDetails)

	return nil
}

func (b *BookDetailsRepository) GetActiveListings() []models.BookDetails {
	var bookDetails []models.BookDetails
	b.DB.Table("book_details").Preload("Book").Where("active = ? AND rented = ?", true, false).Find(&bookDetails)
	return bookDetails
}

func (b *BookDetailsRepository) SearchWithPagination(query string, minPrice, maxPrice, limit, offset int) ([]models.BookDetails, int64) {
	var count int64
	var books []models.BookDetails

	subQuery := b.DB.Debug().Table("books").Select("id").Where("LOWER(title) like LOWER(?)", "%"+query+"%")
	b.DB.Debug().Preload("Book").Where("book_id in (?)", subQuery).Where("price >= ? AND price <= ?", minPrice, maxPrice).Limit(limit).Offset(offset).Find(&books)
	b.DB.Debug().Table("book_details").Where("book_id in (?)", subQuery).Where("price >= ? AND price <= ?", minPrice, maxPrice).Where("active = ? AND rented = ?", true, false).Count(&count)
	return books, count
}

func (b *BookDetailsRepository) SetRented(id uint, value bool) error {
	book, err := b.GetByID(id)
	if err != nil {
		return err
	}

	book.Rented = value
	b.DB.Save(book)

	return nil

}
