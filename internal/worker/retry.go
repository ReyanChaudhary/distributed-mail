package worker

import (
	"log"
	"time"

	"distributed-mail/internal/storage"
	"distributed-mail/internal/queue"
)

// RetryWorker periodically checks for failed emails and requeues them.
func RetryWorker() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		log.Println("Retry worker checking for failed emails...")
		// Retrieve failed emails from the database
		failedEmails, err := storage.GetFailedEmails()
		if err != nil {
			log.Printf("Error retrieving failed emails: %v", err)
			continue
		}

		for _, email := range failedEmails {
			// Requeue the failed email using the queue package
			err := queue.RetryEmail(email.ID)
			if err != nil {
				log.Printf("Error requeuing email %s: %v", email.ID, err)
			} else {
				log.Printf("Successfully requeued email %s", email.ID)
			}
		}
	}
}
