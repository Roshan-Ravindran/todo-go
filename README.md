# Go Fiber RESTful Todo API

A simple RESTful Todo API built with Go, Fiber, PostgreSQL, and GORM.

## Features

- Add, fetch, and toggle (complete/incomplete) todo items
- RESTful API endpoints using Fiber framework
- PostgreSQL database integration with GORM ORM
- Environment variable configuration for database credentials

## Endpoints

| Method | Endpoint             | Description                 |
| ------ | -------------------- | --------------------------- |
| POST   | /api/add_todo        | Add a new todo item         |
| GET    | /api/get_todo        | Get all todo items          |
| GET    | /api/get_todo/:id    | Get a single todo by ID     |
| GET    | /api/toggle_todo/:id | Toggle completion of a todo |

## Getting Started

### Prerequisites

- Go 1.18+
- PostgreSQL
- [Git](https://git-scm.com/)

### Setup

1. Clone the repo:
   ```
   git clone https://github.com/YOUR-USERNAME/YOUR-REPO.git
   cd YOUR-REPO
   ```
2. Copy `.env.example` to `.env` and set your database variables.

3. Install dependencies:
   ```
   go mod tidy
   ```
4. Run the server:
   ```
   go run main.go
   ```

## Folder Structure

- `main.go` – Entry point, sets up API and DB
- `models/` – GORM models and migrations
- `storage/` – Database connection layer

## License

MIT
