# CRUD App in Go (JSON & MongoDB Storage)

A robust CRUD API built with Go and Gin framework, supporting both JSON file-based and MongoDB storage options, with HTTP/HTTPS support.

## 🚀 Features

- **Dual Storage Options**: JSON file storage or MongoDB
- **HTTPS Support**: SSL/TLS with self-signed certificates
- **Docker Integration**: Full containerization support
- **Flexible Configuration**: Environment variables and flags
- **Comprehensive Testing**: Unit tests included

## 📋 Prerequisites

- Go 1.23 or higher
- Docker and Docker Compose (optional)
- OpenSSL (for HTTPS mode)
- Make

## 🛠️ Quick Start

1. Clone and setup:
```sh
git clone https://github.com/pranvatsa/crud.git
cd crud
make install
```

2. Choose your runtime mode:

### HTTP Mode (Default)
```sh
# JSON Storage
make run-json

# MongoDB Storage
make run-mongo
```

### HTTPS Mode
```sh
# Generate certificates (first time only)
make generate-cert

# Run with HTTPS
make run-https
```

### Docker Mode
```sh
# HTTP Mode
make docker-run-http

# HTTPS Mode
make docker-run-https

# With MongoDB (using Docker Compose)
make run-docker
```

## 📡 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET    | `/users` | List all users |
| GET    | `/users/:id` | Get single user |
| POST   | `/users` | Create user |
| PUT    | `/users/:id` | Update user |
| DELETE | `/users/:id` | Delete user |

### Example Request
```json
{
  "name": "John Doe",
  "email": "john@example.com"
}
```

## 🔧 Configuration

### Environment Variables
```sh
STORAGE_MODE=json|mongo  # Storage backend (default: json)
MONGO_URI=<uri>         # MongoDB connection string
PORT=8080              # Server port (default: 8080)
```

### HTTPS Configuration
- Certificate files required for HTTPS mode:
  - `cert.pem`: SSL certificate
  - `key.pem`: Private key

## 📂 Project Structure
```
crud/
├── cmd/main.go           # Entry point
├── config/              # Configuration
├── internal/
│   ├── controllers/     # HTTP handlers
│   ├── database/        # Storage implementations
│   ├── models/          # Data structures
│   └── routes/          # API routes
├── Dockerfile           # Container definition
├── docker-compose.yml   # Multi-container setup
└── Makefile            # Build automation
```

## 🛠️ Make Commands

| Command | Description |
|---------|-------------|
| `make install` | Install dependencies |
| `make run-json` | Run with JSON storage |
| `make run-mongo` | Run with MongoDB |
| `make run-https` | Run with HTTPS |
| `make docker-build-http` | Build HTTP Docker image |
| `make docker-build-https` | Build HTTPS Docker image |
| `make docker-run-http` | Run HTTP container |
| `make docker-run-https` | Run HTTPS container |
| `make run-docker` | Run with Docker Compose |
| `make generate-cert` | Generate SSL certificate |
| `make clean` | Clean build files |

## 🔒 Security Notes

- HTTPS mode uses self-signed certificates (not for production)
- MongoDB credentials should be properly secured
- Input validation implemented for all endpoints
- Error handling follows best practices

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## 📜 License

MIT License - See [LICENSE](LICENSE) file

## ⚠️ Production Usage

Before deploying to production:
1. Replace self-signed certificates with proper SSL certificates
2. Configure secure MongoDB credentials
3. Enable proper logging and monitoring
4. Review security settings