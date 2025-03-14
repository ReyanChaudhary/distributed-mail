package mail

// EmailRequest represents the request payload for sending an email
type EmailRequest struct {
	Recipient string `json:"recipient" binding:"required"` // Receiver's email address
	Subject   string `json:"subject" binding:"required"`   // Email subject
	Body      string `json:"body" binding:"required"`      // Email content
}


// {
// 	"recipient": "user@example.com",
// 	"subject": "Welcome to Our Service",
// 	"body": "Thank you for signing up! We're excited to have you."
//   }
  