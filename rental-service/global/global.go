package global

import (
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

var DB *gorm.DB

var BookEventWriter = &kafka.Writer{
	Addr:                   kafka.TCP("kafka0:29092", "kafka1:29093", "kafka2:29094"),
	Balancer:               &kafka.LeastBytes{},
	AllowAutoTopicCreation: true,
}
