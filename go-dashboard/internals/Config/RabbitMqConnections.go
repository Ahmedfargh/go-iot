package config

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

var RabbitPubConn *amqp.Connection
var RabbitSubConn *amqp.Connection

func InitRabbitMQ() {
	url := AppConfig.RabbitMQURL
	var err error
	RabbitPubConn, err = amqp.Dial(url)
	if err != nil {
		log.Fatalf("Failed to open publishing connection: %v", err)
	}
	RabbitSubConn, err = amqp.Dial(url)
	if err != nil {
		log.Fatalf("Failed to open consuming connection: %v", err)
	}

	log.Println("RabbitMQ connections successfully isolated!")
}

func spawnQueueWorker(workerID int) {

	ch, err := RabbitSubConn.Channel()
	if err != nil {
		log.Printf("Worker %d failed: %v", workerID, err)
		return
	}
	defer ch.Close()

}
