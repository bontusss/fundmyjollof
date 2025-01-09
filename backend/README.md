# FMJ API Project

This project is a Go-based API built using the Gin framework. It provides authentication features such as login, registration, password reset, email verification, and logout.

## Features
- User authentication (Login, Register, Logout)
- Password management (Forgot Password, Reset Password)
- Email verification

## Requirements
- Go 1.20+
- MongoDB
- Mailtrap account (for email testing)

## Setup Instructions

1. **Clone the repository**
   ```bash
   git clone https://github.com/bontusss/fmj.git
   cd fmj

2. **Download dependencies**
    ```bash
    go mod tidy

3. **Start dev server**
    ```bash
    go run .

## Set up environment variables Create a .env file in the root directory and add the following:
BACKEND_PORT=7000
GO_ENV=development
GIN_MODE=debug
MONGO_URI=mongodb://localhost:27017/fundmyjollof (NB: you need to install mongodb)
JWT_SECRET=your_secret_key
DATABASE_NAME=fundmyjollof
SMTP_HOST=sandbox.smtp.mailtrap.io
SMTP_PORT=2525
MAILTRAP_USERNAME=your_mailtrap_username (contact me for test key)
MAILTRAP_PASSWORD=your_mailtrap_password (contact me for test key)
FROM_EMAIL=fundmyjollof@gmail.com
BASE_URL=http://localhost:7000


| Method | Route                 | Description                   |
|--------|-----------------------|-------------------------------|
| POST   | `/auth/login`         | User login                    |
| POST   | `/auth/register`      | User registration             |
| GET    | `/auth/logout`        | User logout                   |
| POST   | `/auth/forgot-password` | Request password reset email |
| POST   | `/auth/reset-password` | Reset user password           |
| GET    | `/auth/verify`        | Verify user email             |
| GET    | `/`                   | Health check or welcome route |

For detailed API documentation and examples, visit the [Postman Docs](https://postman.example.com/your-doc-link).
