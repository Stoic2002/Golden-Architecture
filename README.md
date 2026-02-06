# Golden Architecture - Golang Backend

Modular backend architecture menggunakan **Gin** + **GORM** dengan prinsip Clean Architecture dan Domain-Driven Design.

## 🏗️ Architecture Overview

```
┌─────────────────────────────────────────────────────────────┐
│                        cmd/api/                             │
│                   (Entry Point + DI)                        │
└─────────────────────────┬───────────────────────────────────┘
                          │
┌─────────────────────────▼───────────────────────────────────┐
│                    internal/                                 │
│  ┌─────────────────────────────────────────────────────┐    │
│  │                   domain/                            │    │
│  │  ┌──────────┐  ┌──────────┐  ┌──────────────────┐   │    │
│  │  │ entity/  │  │ contract/│  │ errors.go        │   │    │
│  │  │ (Models) │  │ (Ifaces) │  │ (Domain Errors)  │   │    │
│  │  └──────────┘  └──────────┘  └──────────────────┘   │    │
│  └─────────────────────────────────────────────────────┘    │
│                          ▲                                   │
│  ┌───────────────────────┼─────────────────────────────┐    │
│  │              Feature Modules                         │    │
│  │  ┌─────────────┐     │     ┌─────────────┐          │    │
│  │  │   todo/     │     │     │   user/     │          │    │
│  │  │  ┌────────┐ │     │     │  ┌────────┐ │          │    │
│  │  │  │service │ │─────┼─────│  │service │ │          │    │
│  │  │  └────────┘ │     │     │  └────────┘ │          │    │
│  │  │  ┌────────┐ │     │     │  ┌────────┐ │          │    │
│  │  │  │postgres│ │     │     │  │postgres│ │          │    │
│  │  │  └────────┘ │     │     │  └────────┘ │          │    │
│  │  │  ┌────────┐ │     │     │  ┌────────┐ │          │    │
│  │  │  │handler │ │     │     │  │handler │ │          │    │
│  │  │  └────────┘ │     │     │  └────────┘ │          │    │
│  │  └─────────────┘     │     └─────────────┘          │    │
│  └──────────────────────┼──────────────────────────────┘    │
│                         │                                    │
│  ┌──────────────────────▼──────────────────────────────┐    │
│  │               infrastructure/                        │    │
│  │  ┌──────────┐  ┌──────────┐  ┌──────────┐           │    │
│  │  │ database │  │   http   │  │   auth   │           │    │
│  │  │ (GORM)   │  │  (Gin)   │  │  (JWT)   │           │    │
│  │  └──────────┘  └──────────┘  └──────────┘           │    │
│  └─────────────────────────────────────────────────────┘    │
└─────────────────────────────────────────────────────────────┘
                          │
┌─────────────────────────▼───────────────────────────────────┐
│                        pkg/                                  │
│     ┌──────────┐  ┌──────────┐  ┌──────────┐                │
│     │  logger  │  │ response │  │validator │                │
│     └──────────┘  └──────────┘  └──────────┘                │
└─────────────────────────────────────────────────────────────┘
```

## 📁 Directory Structure

```
project-root/
├── cmd/api/                    # Entry point & DI wiring
│   ├── main.go                 # Application bootstrap
│   └── wire.go                 # Dependency injection
│
├── internal/                   # Private application code
│   ├── domain/                 # 🎯 SOURCE OF TRUTH
│   │   ├── entity/             # Business entities (User, Todo)
│   │   ├── contract/           # Repository interfaces
│   │   └── errors.go           # Domain-level errors
│   │
│   ├── {module}/               # 📦 FEATURE MODULE (todo, user, etc)
│   │   ├── service.go          # Business logic
│   │   ├── postgres/           # Repository implementation
│   │   │   └── repository.go
│   │   └── handler/            # HTTP layer
│   │       ├── http.go         # Handlers
│   │       ├── dto.go          # Request/Response structs
│   │       └── route.go        # Route registration
│   │
│   └── infrastructure/         # 🔧 SHARED INFRASTRUCTURE
│       ├── database/           # PostgreSQL + GORM
│       ├── http/               # Gin server setup
│       └── auth/               # JWT authentication
│
├── pkg/                        # 📚 SHARED UTILITIES
│   ├── logger/                 # Logging wrapper
│   ├── response/               # Standard API response
│   └── validator/              # Input validation
│
├── configs/                    # ⚙️ Configuration
├── migrations/                 # 💾 SQL migrations
├── api/openapi/                # 📖 API documentation
└── scripts/                    # 🛠️ Automation scripts
```

## 🔑 Architecture Principles

### 1. Domain-Centric
```
domain/entity/     → Business entities (data structures)
domain/contract/   → Interfaces (repository contracts)
domain/errors.go   → Domain-specific errors
```

### 2. Module-Based Organization
Setiap feature diorganisir sebagai module mandiri:
```
internal/{module}/
├── service.go          # Business logic (depends on contracts)
├── postgres/           # Database implementation
└── handler/            # HTTP presentation layer
```

### 3. Dependency Flow
```
Handler → Service → Repository (interface)
                         ↓
              PostgreSQL Implementation
```

### 4. Clean Separation
| Layer | Responsibility | Example |
|-------|---------------|---------|
| **Entity** | Data structure | `entity.User`, `entity.Todo` |
| **Contract** | Interface definition | `contract.UserRepository` |
| **Service** | Business logic | `user.Service.Login()` |
| **Repository** | Data access | `postgres.UserRepository` |
| **Handler** | HTTP handling | `handler.Handler.Login()` |

## 🚀 Quick Start

```bash
# Install dependencies
go mod tidy

# Run application
make run

# Access
# API:     http://localhost:8080/api/v1
# Swagger: http://localhost:8080/swagger/index.html
```

## 🔗 API Endpoints

### Todo
| Method | Endpoint | Auth | Description |
|--------|----------|:----:|-------------|
| POST | `/api/v1/todos` | ❌ | Create todo |
| GET | `/api/v1/todos` | ❌ | List todos |
| GET | `/api/v1/todos/:id` | ❌ | Get by ID |
| PUT | `/api/v1/todos/:id` | ❌ | Update |
| DELETE | `/api/v1/todos/:id` | ❌ | Delete |

### Auth
| Method | Endpoint | Auth | Description |
|--------|----------|:----:|-------------|
| POST | `/api/v1/auth/register` | ❌ | Register |
| POST | `/api/v1/auth/login` | ❌ | Login |
| GET | `/api/v1/auth/profile` | ✅ | Get profile |

## 🛠️ Commands

```bash
make run          # Run dev server
make build        # Build binary
make test         # Run tests
make tidy         # Tidy dependencies
```

## 📄 License

MIT
