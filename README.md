# Distributed Mail Service

A lightweight and scalable mail delivery system built using **Gin** and **SMTP** for sending emails. This project is designed to handle email processing efficiently with a structured architecture. It also handles bulk mail delivery/broadcast E-mails.

---

## Features
- REST API for sending emails
- Configurable SMTP settings
- Structured architecture with modular components
- Supports environment-based configuration management

---

## Project Structure
distributed-mail/ │ .env │ .gitignore │ go.mod │ go.sum │ README.md │ ├───cmd │ └───server │ main.go │ ├───config │ config.yaml │ ├───internal │ ├───api │ │ │ routes.go │ │ │ │ │ └───handlers │ ├───mail │ │ sender.go │ │ │ ├───queue │ ├───security │ ├───storage │ └───worker ├───pkg │ ├───logger │ │ logger.go │ │ │ └───utils └───scrip


---

## Installation
### 1. Clone the repository
```sh
git clone https://github.com/reyan/distributed-mail.git
cd distributed-mail


