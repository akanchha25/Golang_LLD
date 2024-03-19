### Project Overview

Welcome to the Payment processing project! This project utilizes Go and PostgreSQL to create a powerful payment and notification micro service for app. This README provides an overview of the project structure, services, setup instructions, and usage of environment variables.

# go_LLD

```
payment-processing/
│
├── cmd/
│   └── main.go
│
├── pkg/
│   ├── payment/
│   │   ├── payment.go
│   │   ├── upi_payment.go
│   │   ├── card_payment.go
│   │   └── netbanking_payment.go
│   │
│   └── notification/
│       ├── notification.go
│       ├── sms_notification.go
│       └── email_notification.go
│
├── internal/
│   ├── domain/
│   │   └── payment_processing.go
│   │
│   └── infrastructure/
│       └── database/
│           └── postgres/
│               └── postgres.go
│
├── model/
│   └── payment.go
│
└── go.mod

```
