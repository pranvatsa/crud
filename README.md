# CRUD App in Go (JSON & MongoDB Storage)

This is a simple **CRUD (Create, Read, Update, Delete) API** built with **Go (Golang)** using **Gin** as the web framework. The app supports **two storage modes**:

1. **JSON mode (default)**: Stores user data in a local JSON file.
2. **MongoDB mode**: Uses a MongoDB database for data persistence.

The entire project is **dockerized** and includes a **Makefile** to simplify installation and running.

---

## 📂 Folder Structure
```
crud-app/
│── cmd/
│   ├── main.go              # Application entry point
│
│── config/
│   ├── config.go            # Configuration settings (Mongo URI, Storage Mode, etc.)
│
│── internal/
│   ├── controllers/
│   │   ├── user_controller.go # Handles HTTP requests for user operations
│   │
│   ├── database/
│   │   ├── mongo.go         # MongoDB interactions
│   │   ├── json_storage.go  # JSON file-based storage interactions
│   │
│   ├── models/
│   │   ├── user.go          # User struct definition
│   │
│   ├── routes/
│   │   ├── user_routes.go   # API route definitions
│   │
│   ├── tests/
│   │   ├── user_test.go     # Unit tests for user operations
│
│── docker-compose.yml       # Docker setup for MongoDB
│── Dockerfile               # Docker setup for the app
│── Makefile                 # Commands to install, build, and run the app
│── go.mod                   # Go module dependencies
│── go.sum                   # Dependency checksums
│── README.md                # Documentation (this file)
```

---

## 🚀 Getting Started

### 1️⃣ Prerequisites
- **Go 1.23** or later
- **Docker & Docker Compose** (for MongoDB mode)

### 2️⃣ Install Dependencies
```sh
make install
```

### 3️⃣ Run the App
#### **Run in JSON Mode (Default)** 📝
```sh
make run-json
```

#### **Run in MongoDB Mode** 🛢️
```sh
make run-mongo
```
> **MongoDB will be started automatically using Docker.**

#### **Run with Docker Compose** (MongoDB mode only)
```sh
make run-docker
```

### 4️⃣ API Endpoints
| Method | Endpoint         | Description           |
|--------|----------------|-----------------------|
| GET    | `/users`        | Get all users        |
| GET    | `/users/:id`    | Get user by ID       |
| POST   | `/users`        | Create a new user    |
| PUT    | `/users/:id`    | Update user by ID    |
| DELETE | `/users/:id`    | Delete user by ID    |

### 5️⃣ Stopping MongoDB
```sh
make stop-mongo
```

---

## 🛠️ Configuration
Modify **`config/config.go`** to change storage mode:
```go
const StorageMode = "json" // Change to "mongo" for MongoDB
```

Alternatively, set an environment variable:
```sh
export STORAGE_MODE=mongo  # or json
```

---

## 🧪 Running Tests
Run unit tests with:
```sh
make test
```
or manually:
```sh
go test ./...
```

---

## 📜 License
This project is **open-source** and available under the **MIT License**.
