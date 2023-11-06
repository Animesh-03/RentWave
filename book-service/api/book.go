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
	Isbn        string `json:"isbn" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Image       string `json:"image" binding:"required"`
	Description string `json:"description" binding:"required"`
	OwnerID     uint   `json:"ownerid" binding:"required"`
	Price       uint   `json:"price" binding:"required"`
}

func SearchBooks(c *gin.Context) {
	query := c.Query("q")
	minPrice, err := strconv.Atoi(c.DefaultQuery("minPrice", "0"))
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"error": "min price is not an integer",
		})
		return
	}

	maxPrice, err := strconv.Atoi(c.DefaultQuery("maxPrice", "1000000"))
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"error": "max price is not an integer",
		})
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "25"))
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"error": "limit is not an integer",
		})
		return
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"error": "offset is not an integer",
		})
		return
	}

	bookDetailsRepository := repositories.NewBookDetailsRepository(global.DB)
	books, count := bookDetailsRepository.SearchWithPagination(query, minPrice, maxPrice, limit, offset)
	c.IndentedJSON(200, gin.H{
		"books": books,
		"count": count,
		"next":  int64(offset+limit) < count,
	})
}

func AddBookHandler(c *gin.Context) {
	var bookData BookData
	err := c.BindJSON(&bookData)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(400, &gin.H{
			"error": "invalid request body",
		})
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
	c.IndentedJSON(200, gin.H{
		"books": bookDetailsRepository.GetActiveListings(),
	})
}

func GetBookByIsbn(c *gin.Context) {
	isbn := c.Query("isbn")

	bookRepository := repositories.NewBookRepository(global.DB)
	book, err := bookRepository.GetByIsbn(isbn)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.IndentedJSON(200, gin.H{
			"found": false,
		})
		return
	}
	c.IndentedJSON(200, gin.H{
		"found": true,
		"book":  book,
	})
}
