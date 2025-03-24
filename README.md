# Distributed Mail Service

A lightweight and scalable mail delivery system built using **Gin** and **SMTP** for sending emails. This project is designed to handle email processing efficiently with a structured architecture. It also handles bulk mail delivery/broadcast E-mails.

---

## Features
- REST API for sending emails
- Configurable SMTP settings
- Kubernetes for maintaining asynchronous queues producers/consumers and relative sequence
- Structured architecture with modular components
- Supports environment-based configuration management

---

distributed-mail/
│── cmd/                  
│   ├── server/           # Entry point for the server
│   │   ├── main.go       # Initializes and starts the service
│   │   ├── config.go     # Loads configuration (env vars, flags)
│── internal/             
│   ├── api/              # Handles API requests & responses
│   │   ├── handlers/     # HTTP handlers
│   │   ├── routes.go     # Router configuration
│   ├── mail/             # Core email handling logic
│   │   ├── sender.go     # Sending emails
│   │   ├── receiver.go   # Receiving emails
│   │   ├── parser.go     # Parsing email content
│   ├── queue/            # Distributed queue (RabbitMQ, Kafka, etc.)
│   │   ├── producer.go   # Produces tasks for processing
│   │   ├── consumer.go   # Consumes and processes tasks
│   ├── storage/          # Handles database interactions
│   │   ├── models.go     # Defines data models
│   │   ├── repository.go # Database queries
│   ├── security/         # Authentication & Encryption
│   │   ├── auth.go       # Authentication mechanisms
│   │   ├── encryption.go # Email encryption (PGP, etc.)
│   ├── worker/           # Background processing (cron jobs, retries)
│   │   ├── retry.go      # Implements retry logic for failed emails
│── config/               
│   ├── config.yaml       # Configuration file for environment variables
│── pkg/                  # Reusable utility packages
│   ├── logger/           # Logging utilities
│   ├── utils/            # Miscellaneous helper functions
│── scripts/              # Deployment & automation scripts
│── .env                  # Environment variables (ignored in Git)
│── .gitignore            # Ignoring sensitive & unnecessary files
│── go.mod                # Go module dependencies
│── go.sum                # Go dependency checksums
│── README.md             # Project documentation


---

## Installation
### 1. Clone the repository
```sh
git clone https://github.com/reyan/distributed-mail.git
cd distributed-mail


