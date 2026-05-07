# Gin Backend Template

A sample Go backend project using the [Gin Gonic](https://github.com/gin-gonic/gin) web framework. This project serves as a boilerplate for quickly setting up a RESTful API with User and Todo management features.

## Features

- **User Management**: CRUD operations for users.
- **Todo Management**: CRUD operations for tasks (todos).
- **Health Checks**: Basic endpoint for monitoring service status.
- **In-Memory Storage**: Simple slice-based storage for quick prototyping (easily replaceable with a database).

## Tech Stack

- **Language**: Go (Golang)
- **Web Framework**: [Gin Gonic](https://github.com/gin-gonic/gin)
- **Architecture**: Controller-Model based structure

## Prerequisites

- [Go](https://golang.org/doc/install) (version 1.26.2 or higher recommended)

## Installation & Setup

1. **Clone the repository**:

```bash
git clone https://github.com/Sujal-Gaha/gin-backend.git
cd gin-backend
```

2. **Install dependencies**:

```bash
go mod download
```

3. **Run the application**:

```bash
go run main.go
```

The server will start on `http://localhost:8080`.

## API Endpoints

### General

| Method | Endpoint  | Description          |
| ------ | --------- | -------------------- |
| GET    | `/`       | Welcome greeting     |
| GET    | `/health` | Service health check |

### Users

| Method | Endpoint     | Description         |
| ------ | ------------ | ------------------- |
| GET    | `/users`     | Get all users       |
| POST   | `/users`     | Create a new user   |
| GET    | `/users/:id` | Get a specific user |
| PUT    | `/users/:id` | Update a user       |
| DELETE | `/users/:id` | Delete a user       |

### Todos

| Method | Endpoint     | Description         |
| ------ | ------------ | ------------------- |
| GET    | `/todos`     | Get all todos       |
| POST   | `/todos`     | Create a new todo   |
| GET    | `/todos/:id` | Get a specific todo |
| PUT    | `/todos/:id` | Update a todo       |
| DELETE | `/todos/:id` | Delete a todo       |

## Usage Examples

### Create a User

```bash
curl -X POST http://localhost:8080/users \
-H "Content-Type: application/json" \
-d '{"name": "John Doe", "email": "john@example.com"}'
```

### Create a Todo

```bash
curl -X POST http://localhost:8080/todos \
-H "Content-Type: application/json" \
-d '{"title": "Learn Go", "completed": false, "user_id": 1}'
```

## Project Structure

```text
.
├── controllers/      # Route handlers
├── models/           # Data structures and schemas
├── main.go           # Entry point and routing
├── go.mod            # Go module definition
└── README.md         # Project documentation
```
