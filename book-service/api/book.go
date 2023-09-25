package api

import (
	"book/global"
	"book/models"
	"book/repositories"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookData struct {
	ID          uint   `json:"id"`
	Isbn        string `json:"isbn"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Image       string `json:"image"`
	Description string `json:"description"`
	OwnerID     uint   `json:"ownerid"`
	Price       uint   `json:"price"`
}

func AddBookHandler(c *gin.Context) {
	var bookData BookData
	err := c.BindJSON(&bookData)
	if err != nil {
		fmt.Println(err)
		return
	}

	bookRepository := repositories.NewBookRepository(global.DB)
	bookDetailsRepository := repositories.NewBookDetailsRepository(global.DB)

	book, err := bookRepository.GetByIsbn(bookData.Isbn)

	fmt.Println(bookData, err, book)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Create the book and add the book details
		fmt.Println("Creating a new book")
		bookRepository.AddEntity(&models.Book{
			Isbn:   bookData.Isbn,
			Title:  bookData.Title,
			Author: bookData.Author,
			Details: []models.BookDetails{
				{
					OwnerID:     bookData.OwnerID,
					Image:       bookData.Image,
					Description: bookData.Description,
					Price:       bookData.Price,
				},
			},
		})
	} else {
		// Attach the book details
		fmt.Println("Adding details to existing book")

		bookDetailsRepository.AddEntity(&models.BookDetails{
			BookID:      book.ID,
			OwnerID:     bookData.OwnerID,
			Image:       bookData.Image,
			Description: bookData.Description,
		})
	}

	c.IndentedJSON(200, "OK")
}

// func GetUserListings(c *gin.Context) {
// 	bookDetailsRepository := repositories.NewBookDetailsRepository(global.DB)
// 	bookDetailsRepository.

// }

func UpdateBookHandler(c *gin.Context) {

}
