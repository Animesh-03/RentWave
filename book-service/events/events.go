package events

func HandleKafkaEvents() {
	go RentedEventHandler()
	go RentalCompleteEventHandler()
}
