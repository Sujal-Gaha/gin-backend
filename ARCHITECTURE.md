# Clean Architecture Go Backend

This project demonstrates a production-ready Go backend using the **Gin** framework and **Clean Architecture** (DDD-style). The goal is to achieve a strict separation of concerns, making the codebase highly testable, maintainable, and scalable.

## 🏗 Project Structure

```text
.
├── cmd/
│   └── server/
│       └── main.go           # Application entry point & server initialization
├── internal/
│   ├── domain/               # Core Business Logic (Entities & Interfaces)
│   │   ├── todo.go           # Domain models
│   │   └── repository.go     # Repository interfaces
│   ├── usecase/              # Application Logic (Orchestration)
│   │   ├── create_todo.go    # Individual use case implementations
│   │   └── interface.go      # Common UseCase interface
│   ├── repository/           # Data Access Layer (Implementations)
│   │   └── inmemory/         # In-memory DB implementation
│   ├── controller/           # Delivery Layer (Transport)
│   │   └── http/             # Gin HTTP handlers
│   └── registry/             # Dependency Injection & Wiring
│       └── container.go      # Manual DI container
├── go.mod                    # Go module definition
└── ARCHITECTURE.md           # Architecture documentation
```

---

## 층별 역할 (Layer Responsibilities)

### 1. Domain Layer (`internal/domain`)
The "heart" of the application. It contains the business entities and the interfaces (contracts) that other layers must fulfill. It has **zero dependencies** on any other layer or external library.

### 2. Use Case Layer (`internal/usecase`)
Contains the application-specific business rules. It orchestrates the flow of data to and from the entities and directs the repositories to persist changes. It depends only on the Domain layer.

### 3. Repository Layer (`internal/repository`)
Responsible for data persistence. Whether it's PostgreSQL, MongoDB, or an In-Memory store, the implementation details stay here. It implements the interfaces defined in the Domain layer.

### 4. Controller/Delivery Layer (`internal/controller`)
Handles the "entry points" to the application (HTTP, gRPC, CLI). In this project, it uses Gin to handle HTTP requests, parse inputs, and return JSON responses. It communicates only with the Use Case layer.

### 5. Registry/DI Layer (`internal/registry`)
The "glue" of the application. It instantiates the repositories, use cases, and controllers, and injects the dependencies where needed. This prevents global variables and ensures a clean startup.

---

## 🌟 Advantages of this Structure

### 1. Independence of Frameworks
The architecture does not depend on the existence of some library of feature-laden software. This allows you to use Gin (or any other framework) as a tool, rather than having to cram your system into its limited constraints.

### 2. Testability
Since the business logic (Use Cases) is separated from the delivery mechanism (HTTP) and the persistence layer (Database), you can write unit tests for your business rules without needing a running server or a database connection.

### 3. Independence of UI
The UI (or API layer) can change easily without changing the rest of the system. You could swap Gin for Fiber or even add a CLI interface without touching the Use Cases or Entities.

### 4. Independence of Database
You can swap out the In-Memory repository for a real Database (SQL/NoSQL) implementation by simply creating a new file in the repository layer and updating the `container.go` file.

---

## 🚀 Scalability & Flexibility

### How it is Scalable:
- **Development Scale**: Since the layers are decoupled, multiple developers can work on different parts of the system (e.g., one on the Repository implementation, another on Use Case logic) without causing merge conflicts or breaking each other's code.
- **Performance Scale**: You can optimize individual layers. If a specific Use Case is slow, you can optimize its logic or its data access pattern in the Repository without affecting the Controller.

### How it is Flexible:
- **Plug-and-Play**: Need to add Redis caching? Create a `RedisTodoRepository` that implements the same interface, and inject it. The Use Case won't even know the difference.
- **Strict Contracts**: By using interfaces for repositories and use cases, you ensure that every component has a clear contract. This makes refactoring significantly safer.
- **Context-Aware**: The use of `context.Context` throughout all layers allows for easy implementation of timeouts, cancellations, and distributed tracing, which are essential for production-grade microservices.

---

## 🔄 Request Flow
`HTTP Request` -> `Gin Handler` -> `UseCase` -> `Repository Interface` -> `In-Memory Implementation` -> `Response`
