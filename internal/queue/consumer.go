package queue

import (
	"context"
	"encoding/json"
	"log"
	
	"github.com/segmentio/kafka-go"
	"distributed-mail/internal/mail"
	"distributed-mail/internal/storage"
)

// StartKafkaConsumer listens for email tasks from Kafka and processes them.
func StartKafkaConsumer(broker, topic, groupID string) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{broker},
		GroupID:   groupID,
		Topic:     topic,
		MinBytes:  10e3,  // 10KB
		MaxBytes:  10e6,  // 10MB
	})
	defer reader.Close()
	log.Printf("Kafka consumer started on topic: %s", topic)

	for {
		ctx := context.Background()
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("Error reading message: %v", err)
			continue
		}
		log.Printf("Message received: key=%s, value=%s", string(msg.Key), string(msg.Value))

		// Unmarshal the message
		var payload struct {
			EmailID string           `json:"email_id"`
			Request mail.EmailRequest `json:"request"`
		}
		if err := json.Unmarshal(msg.Value, &payload); err != nil {
			log.Printf("Error unmarshalling message: %v", err)
			continue
		}

		// Send email via SMTP (this includes scheduling and personalization)
		if err := mail.SendEmail(payload.Request); err != nil {
			log.Printf("Error sending email for ID %s: %v", payload.EmailID, err)
			// Update status to failed
			storage.UpdateEmailStatus(payload.EmailID, "failed")
		} else {
			// Update status to sent
			storage.UpdateEmailStatus(payload.EmailID, "sent")
		}
	}
}
