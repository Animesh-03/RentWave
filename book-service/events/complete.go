package events

import (
	"book/global"
	"book/repositories"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func RentalCompleteEventHandler() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"kafka0:29092", "kafka1:29093", "kafka2:29094"},
		GroupID:  "book-service.completed",
		Topic:    "book.completed",
		MaxBytes: 10e6, // 10MB
	})

	bookDetailsRepository := repositories.NewBookDetailsRepository(global.DB)

	for {
		fmt.Println("Reading Message(book.completed)...")
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))

		var msg RentedEvent
		json.Unmarshal(m.Value, &msg)

		bookDetailsRepository.SetRented(msg.EntityID, false)
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
