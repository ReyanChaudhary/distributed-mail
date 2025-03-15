package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"time"
)

// SendEmail sends an email via SMTP, handling scheduling and personalization.
func SendEmail(req EmailRequest) error {
	// Handle scheduling: delay sending if ScheduledAt is set in the future.
	if !req.ScheduledAt.IsZero() && time.Now().Before(req.ScheduledAt) {
		delay := req.ScheduledAt.Sub(time.Now())
		log.Printf("Delaying email to %s for %v", req.Recipient, delay)
		time.Sleep(delay)
	}

	// Personalize the email if template data is provided.
	body := req.Body
	if len(req.TemplateData) > 0 {
		t, err := template.New("email").Parse(req.Body)
		if err != nil {
			return fmt.Errorf("failed to parse email template: %w", err)
		}
		var buf bytes.Buffer
		if err := t.Execute(&buf, req.TemplateData); err != nil {
			return fmt.Errorf("failed to execute email template: %w", err)
		}
		body = buf.String()
	}

	// SMTP settings (ideally, these come from configuration)
	smtpHost := "smtp.example.com"
	smtpPort := "587"
	smtpUser := "your-email@example.com"
	smtpPassword := "your-password"

	from := smtpUser
	to := []string{req.Recipient}

	// Construct the email message.
	msg := []byte(fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", from, req.Recipient, req.Subject, body))

	auth := smtp.PlainAuth("", smtpUser, smtpPassword, smtpHost)
	addr := fmt.Sprintf("%s:%s", smtpHost, smtpPort)
	if err := smtp.SendMail(addr, auth, from, to, msg); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	log.Printf("Email sent successfully to %s", req.Recipient)
	return nil
}
