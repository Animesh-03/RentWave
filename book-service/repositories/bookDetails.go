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

func (b *BookDetailsRepository) GetActiveListings() []models.Book {
	var bookDetails []models.Book
	subQuery := b.DB.Table("book_details").Select("book_id").Where("active = ?", true)
	b.DB.Debug().Preload("Details").Where("id in (?)", subQuery).Find(&bookDetails)
	return bookDetails
}
