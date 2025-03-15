package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"distributed-mail/internal/mail"
	"distributed-mail/internal/storage"
)

// writer is the global Kafka writer instance.
var writer *kafka.Writer

// InitKafkaWriter initializes the Kafka writer with the given broker and topic.
func InitKafkaWriter(broker, topic string) error {
	writer = &kafka.Writer{
		Addr:         kafka.TCP(broker),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		BatchTimeout: 10 * time.Millisecond, // adjust as needed
		// Optionally add BatchSize, RequiredAcks, etc.
	}
	log.Printf("Kafka writer initialized for broker: %s, topic: %s", broker, topic)
	return nil
}

// ProduceEmailTask publishes an email request to Kafka and returns a generated email ID.
func ProduceEmailTask(emailRequest mail.EmailRequest) (string, error) {
	if writer == nil {
		return "", fmt.Errorf("kafka writer is not initialized")
	}

	// Generate a unique email ID.
	emailID := uuid.New().String()

	// Save the email record to the DB with initial status "queued"
	if err := storage.SaveEmail(emailID, "queued"); err != nil {
		return "", fmt.Errorf("failed to save email to DB: %w", err)
	}

	// Prepare the payload.
	payload := struct {
		EmailID string            `json:"email_id"`
		Request mail.EmailRequest `json:"request"`
	}{
		EmailID: emailID,
		Request: emailRequest,
	}

	msgBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal email request: %w", err)
	}

	// Create a context with timeout to avoid hanging indefinitely.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Write message to Kafka.
	err = writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(emailID),
		Value: msgBytes,
	})
	if err != nil {
		return "", fmt.Errorf("failed to write message to Kafka: %w", err)
	}

	log.Printf("Queued email task with ID: %s", emailID)
	return emailID, nil
}

// RetryEmail requeues a failed email task. In production, you might retrieve
// the original email request from persistent storage before requeuing.
func RetryEmail(emailID string) error {
	if writer == nil {
		return fmt.Errorf("kafka writer is not initialized")
	}

	// TODO: Retrieve the original email request from storage if needed.

	log.Printf("Requeued email with ID: %s", emailID)
	// Here you would push the message back to Kafka (or update the status in DB)
	return nil
}

// CloseKafkaWriter gracefully closes the Kafka writer.
func CloseKafkaWriter() error {
	if writer != nil {
		return writer.Close()
	}
	return nil
}
