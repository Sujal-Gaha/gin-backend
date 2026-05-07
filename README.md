# 🚀 Go Gold Standard: Modular Clean Architecture

A production-ready, enterprise-grade backend boilerplate built with **Go** and the **Gin** framework. This project demonstrates the "Gold Standard" of Go development, emphasizing extreme modularity, testability, and operational excellence.

---

## 🏗 The Architecture

This project follows a **Modular Clean Architecture** (Domain-Driven Design inspired). Unlike traditional flat structures, this layout scales gracefully as your team and feature set grow.

### 📁 Directory Structure

```text
.
├── cmd/server/main.go          # Entry point (Graceful shutdown, slog, config)
├── internal/
│   ├── config/                 # 12-Factor Config (Env-based)
│   ├── domain/                 # THE CORE (Zero dependencies)
│   │   ├── todo/               # Todo Module Domain (Entity, Repo Interface)
│   │   │   └── usecase/        # Todo Business Logic
│   │   ├── user/               # User Module Domain
│   │   │   └── usecase/        # User Business Logic
│   │   └── errors.go           # Standardized Domain Errors
│   ├── modules/                # Delivery Layer (Gin Handlers)
│   │   ├── todo/               # Todo API (Mutations & Queries split)
│   │   └── user/               # User API
│   ├── repository/             # Infrastructure Layer (Persistence)
│   │   └── mem_repo/           # Thread-safe In-Memory implementation
│   └── registry/               # Dependency Injection (Manual Wiring)
└── ARCHITECTURE.md             # Deep-dive into architectural decisions
```

---

## 🌟 Key Features (The "Gold Standard")

### 1. 🛡 Modular Encapsulation

Each feature (User, Todo) is self-contained. Adding a new feature doesn't clutter existing packages. This prevents the "Big Ball of Mud" anti-pattern.

### 2. 🚦 Operational Excellence

- **Graceful Shutdown**: Listens for termination signals (`SIGINT`, `SIGTERM`) and allows active requests to finish (Kubernetes-ready).
- **Structured Logging**: Uses `log/slog` for JSON-formatted logs, essential for ELK, Datadog, or CloudWatch.
- **Environment Safety**: Centralized configuration management with safe defaults.

### 3. 🧩 Decoupled Business Logic (CQS)

Handlers are separated into **Mutations** (state-changing) and **Queries** (data-fetching). This clarifies intent and prepares the system for read/write scaling.

### 4. 🧪 Testability & Purity

The `domain` layer has **zero external dependencies**. Business rules are defined by interfaces, allowing you to swap the In-Memory store for PostgreSQL or MongoDB without touching a single line of business logic.

---

## 🚀 Getting Started

### Prerequisites

- Go 1.21 or higher

### Installation

1. Clone the repository
2. Install dependencies:

```bash
go mod download
```

### Running the App

```bash
# Development
go run cmd/server/main.go

# Production Simulation
PORT=9000 APP_ENV=production go run cmd/server/main.go
```

### API Endpoints

| Method | Endpoint     | Description                        |
| :----- | :----------- | :--------------------------------- |
| `POST` | `/api/users` | Create a new user                  |
| `POST` | `/api/todos` | Create a todo (requires `user_id`) |

---

## 🛠 Why this is "Gold Standard"?

1. **Separation of Concerns**: The Web Framework (Gin) knows nothing about the Database. The Database knows nothing about the Web.
2. **Context Propagation**: `context.Context` is passed through every layer for proper timeout and cancellation handling.
3. **Domain Errors**: We don't return `404` from the DB. We return `ErrNotFound`, and the delivery layer translates it.
4. **No Global State**: Everything is injected via the `registry` container.

---

## 📝 License

This project is licensed under the MIT License - see the LICENSE file for details.
