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
   git clone https://github.com/Roshan-Ravindran/todo-go.git
   cd todo-go
   ```
2. Copy `.env.example` to `.env` and set your database variables.

   ## Example Env File

   ```
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASS= **\*\*\*\***
   DB_NAME=todo_db
   DB_SSLMODE=disable
   ```

3. Install dependencies:
   ```
   go mod tidy
   ```
4. Run the server:
   ```
   go run main.go
   ```

### Example POST /api/add_todo body:

    {
    "item": "Buy groceries",
    "completed": false
    }

### 📦 Folder Structure

    ├── main.go
    ├── models/
    │ └── todo_list.go
    ├── storage/
    │ └── postgres.go
    ├── go.mod
    ├── go.sum
    └── .env

### 🎯 Future Improvements

- Add JWT authentication
- Add CLI interface for managing Todos
- Deploy using Docker and AWS ECS

### 📄 License

This project is open-source under the MIT License.
