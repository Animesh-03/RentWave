package api

import (
	"rental/global"
	"rental/models"
	"rental/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HealthCheckHandler(c *gin.Context) {
	c.IndentedJSON(200, nil)
}

type RentEntityData struct {
	Type     models.RentalType `json:"type"`
	EntityID uint              `json:"entityid"`
	Lessor   uint              `json:"lessor"`
	From     uint64            `json:"from"`
	To       uint64            `json:"to"`
}

func RentEntityHandler(c *gin.Context) {
	var body RentEntityData
	err := c.BindJSON(&body)

	if err != nil {
		c.IndentedJSON(400, &gin.H{
			"error": "invalid request body",
		})
		return
	}

	rentalRepository := repositories.NewRentalRepository(global.DB)

	rentalRepository.AddEntity(&models.Rental{
		Type:     body.Type,
		EntityID: body.EntityID,
		Lessor:   body.EntityID,
		From:     body.From,
		To:       body.To,
	})

	c.IndentedJSON(200, nil)
}

func GetUserRentalsHandler(c *gin.Context) {
	uid, err := (strconv.Atoi(c.Query("uid")))

	if err != nil {
		c.IndentedJSON(400, &gin.H{
			"error": "invalid request",
		})
		return
	}

	rentalRepository := repositories.NewRentalRepository(global.DB)
	rentals := rentalRepository.GetAllUserRentals(uint(uid))

	c.IndentedJSON(200, rentals)
}

func GetActiveUserRentalsHandler(c *gin.Context) {
	uid, err := (strconv.Atoi(c.Query("uid")))

	if err != nil {
		c.IndentedJSON(400, &gin.H{
			"error": "invalid request",
		})
		return
	}

	rentalRepository := repositories.NewRentalRepository(global.DB)
	rentals := rentalRepository.GetActiveUserRentals(uint(uid))

	c.IndentedJSON(200, rentals)
}
