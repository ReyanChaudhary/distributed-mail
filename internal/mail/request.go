package mail

import "time"

// EmailRequest represents the request payload for sending an email.
type EmailRequest struct {
	Recipient    string            `json:"recipient" binding:"required,email"` // Receiver's email address
	Subject      string            `json:"subject" binding:"required"`           // Email subject
	Body         string            `json:"body" binding:"required"`              // Email content (or template)
	ScheduledAt  time.Time         `json:"scheduled_at"`                         // Optional: if zero, send immediately
	TemplateData map[string]string `json:"template_data,omitempty"`              // For personalization
}
