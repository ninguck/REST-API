# 🏦 Wallet API

A modular, test-driven RESTful Wallet API built with Go, PostgreSQL, and Docker. This backend service handles user and wallet management functionality with a strong focus on clean architecture and testability.

---

## 📦 Prerequisites

Ensure the following are installed on your system:

- [Go](https://go.dev/doc/install) (version ≥ 1.21)
- [Docker Desktop](https://www.docker.com/products/docker-desktop)
- [Git](https://git-scm.com/downloads)
- PowerShell (Windows) or Terminal (macOS/Linux)

---

## 🧱 Project Structure

```
wallet-api/
├── backend/
│   ├── cmd/                  # Main entry point
│   ├── internal/
│   │   ├── api/              # API route handlers
│   │   ├── db/               # DB initialization
│   │   └── models/           # Core business logic
│   ├── tests/                # Unit and integration tests
│   ├── go.mod / go.sum
│   └── docker-compose.yml
```

---

## 🚀 Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/<your-username>/wallet-api.git
cd wallet-api/backend
```

### 2. Start the PostgreSQL Database

```bash
docker compose up -d
```

- **DB Name**: `wallet_db`
- **User**: `wallet_user`
- **Password**: `wallet_pass`
- **Host**: `localhost`
- **Port**: `5432`

### 3. Run the Go Server

```bash
go run cmd/main.go
```

Visit: [http://localhost:8000](http://localhost:8000)

### 4. Ping Test

```bash
curl http://localhost:8000/ping
# Output: pong
```

### 5. API Endpoints

| Method | Endpoint           | Description           |
|--------|--------------------|-----------------------|
| GET    | `/ping`            | Health check          |
| POST   | `/users`           | Create a new user     |
| GET    | `/users/{id}`      | Get user by ID        |
| POST   | `/wallets`         | Create a new wallet   |
| GET    | `/wallets/{id}`    | Get wallet by ID      |

---

## 🧪 Running Tests

All tests follow the TDD methodology.

```bash
go test ./tests
```

This will run:
- User creation and fetching
- Wallet creation and fetching

Tests use dynamic emails to avoid DB conflicts.

---

## ⚠️ Troubleshooting

- If you get a "connection refused" error, ensure Docker is running and that port 5432 is mapped.
- Use PowerShell for Windows REST API testing:
```powershell
Invoke-RestMethod -Uri "http://localhost:8000/users" `
  -Method Post `
  -ContentType "application/json" `
  -Body '{"name":"Nicholas","email":"nick123@example.com"}'
```

---

## 📌 Future Additions

- Transaction management (deposits/withdrawals)
- User authentication (JWT)
- Dockerized backend image
- Swagger/OpenAPI documentation
- CI/CD pipeline and GitHub Actions

---

## 👨‍💻 Author

Built with ❤️ by Nicholas Nguyen as part of VGW-inspired L&D backend system.