package api

import (
	"context"
	"encoding/json"
	"fmt"
	"rental/global"
	"rental/models"
	"rental/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
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
			"error": "invalid request body: " + err.Error(),
		})
		return
	}

	rentalRepository := repositories.NewRentalRepository(global.DB)

	rentalRepository.AddEntity(&models.Rental{
		Type:     body.Type,
		EntityID: body.EntityID,
		Lessor:   body.Lessor,
		From:     body.From,
		To:       body.To,
	})

	c.IndentedJSON(200, nil)

	bodyBytes, _ := json.Marshal(body)

	message := kafka.Message{
		Topic: fmt.Sprintf("%s.rented", string(body.Type)),
		Value: bodyBytes,
	}

	global.BookEventWriter.WriteMessages(context.Background(), message)
}

func CompleteRental(c *gin.Context) {
	var data struct {
		ID uint `json:"id"`
	}
	c.BindJSON(&data)

	rentalRepository := repositories.NewRentalRepository(global.DB)

	rental, _ := rentalRepository.SetStatus(data.ID, models.COMPLETED)

	c.IndentedJSON(200, nil)

	bodyBytes, _ := json.Marshal(rental)

	message := kafka.Message{
		Topic: fmt.Sprintf("%s.completed", string(rental.Type)),
		Value: bodyBytes,
	}

	global.BookEventWriter.WriteMessages(context.Background(), message)

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

	if rentals == nil {
		rentals = make([]*models.Rental, 0)
	}

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

	if rentals == nil {
		rentals = make([]*models.Rental, 0)
	}

	c.IndentedJSON(200, rentals)
}
