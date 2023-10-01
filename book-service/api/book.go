package api

import (
	"book/global"
	"book/models"
	"book/repositories"
	"errors"
	"fmt"
	"strconv"

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
			Details: []*models.BookDetails{
				{
					BookID:      book.ID,
					OwnerID:     bookData.OwnerID,
					Image:       bookData.Image,
					Description: bookData.Description,
					Price:       bookData.Price,
					Active:      true,
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
			Price:       bookData.Price,
			Active:      true,
		})
	}

	c.IndentedJSON(200, "OK")
}

func GetUserListings(c *gin.Context) {
	uidString, ok := c.GetQuery("uid")
	if !ok {
		c.IndentedJSON(400, gin.H{
			"error": "invalid request body",
		})
		return
	}
	uid, err := strconv.ParseUint(uidString, 10, 32)
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"error": "invalid request body",
		})
		return
	}

	bookDetailsRepository := repositories.NewBookDetailsRepository(global.DB)
	userListings := bookDetailsRepository.GetUserListings(uint(uid))
	c.IndentedJSON(200, userListings)
}

type UpdateBookData struct {
	ID uint `json:"_id" binding:"required"`
}

func ToggleListing(c *gin.Context) {
	var data UpdateBookData
	err := c.BindJSON(&data)
	if err != nil {
		c.IndentedJSON(400, &gin.H{
			"error": "invalid request body",
		})
		return
	}

	bookDetailsRepository := repositories.NewBookDetailsRepository(global.DB)
	bookDetailsRepository.ToggleListing(data.ID)

	c.IndentedJSON(200, "OK")
}

func GetActiveListings(c *gin.Context) {
	bookDetailsRepository := repositories.NewBookDetailsRepository(global.DB)
	c.IndentedJSON(200, bookDetailsRepository.GetActiveListings())
}
