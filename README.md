# Game Store CLI

## Project Structure
```
ftgo-8-phase-1-pair-project/
├── cli/
├── config/
│   └── config.go
├── database/
│   └── postgres.go
├── docs/
│   └── games-erd.go
├── entity/
│   ├── category.go
│   ├── game.go
│   ├── order.go
│   └── user.go
├── handler/
│   ├── game_test.go
│   ├── game.go
│   ├── order_test.go
│   ├── order.go
│   ├── report_test.go
│   ├── report.go
│   ├── user_test.go
│   └── user.go
├── .env
├── .env.example
├── .gitignore
├── go.mod
├── go.sum
├── main.go
└── schema.sql
```
## Project Overview

Game Store CLI is a command-line interface (CLI) application for managing a game store, built with Go and PostgreSQL. The application provides robust functionality for both customers and administrators, allowing users to browse games, place orders, and generate various reports.

## Key Features

### User Authentication
- User registration and login with role-based access
- Secure password hashing using `bcrypt`
- Differentiated access for customers and administrators

### Game Management (Admin)
- Add new games to the store
- Update existing game details
- Delete games
- View game inventory

### Order Management
- Place orders for games
- View order history
- Real-time stock tracking
- Transaction management with database constraints

### Reporting System
The application includes comprehensive reporting features:
- Low Stock Alert Report
- Customer Purchase Frequency Analysis
- Sales Performance by Category
- Recent Game Releases Performance
- Average Order Value by Month

## Technical Architecture

### Database Schema
The project uses PostgreSQL with the following main tables:
- `users`: Stores user information
- `games`: Contains game details
- `orders`: Tracks user purchases
- `categories`: Manages game categories

### Key Technologies
- Language: Go (Golang)
- Database: PostgreSQL
- Libraries:
  - `database/sql` for database interactions
  - `godotenv` for environment configuration
  - `bcrypt` for password security
  - `sqlmock` for testing
  - `testify` for assertions
