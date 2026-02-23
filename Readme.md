Here's a README for your Chauffer-Service:

***

# 🚗 Chauffer-Service

> *Part of the Distributed Systems Project*

Chauffer-Service is a microservice responsible for managing road bookings for travellers. It handles the reservation of road segments, routes, and travel slots, acting as the core booking engine within a larger distributed system architecture.

***

## 📦 Project Structure

```
chauffer-service/
├── main.go
├── go.mod
├── go.sum
├── config/
│   └── config.go
├── routes/
│   └── routes.go
├── handlers/
│   ├── health.go
│   └── booking.go
└── middleware/
    └── logger.go
```

***

## 🚀 Getting Started

### Prerequisites

- [Go 1.21+](https://golang.org/dl/)
- [Git](https://git-scm.com/)

### Installation

```bash
# Clone the repository
git clone https://github.com/your-org/chauffer-service.git
cd chauffer-service

# Install dependencies
go mod tidy

# Run the service
go run main.go
```

The service starts on port `8080` by default. Override with the `PORT` environment variable:

```bash
PORT=9090 go run main.go
```

***

## 🛣️ API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/api/v1/checkHealth` | Returns service health status |
| `POST` | `/api/v1/bookings` | Create a new road booking |
| `GET` | `/api/v1/bookings/:id` | Retrieve a booking by ID |
| `DELETE` | `/api/v1/bookings/:id` | Cancel an existing booking |

### Health Check

```bash
curl http://localhost:8080/api/v1/checkHealth
```

```json
{ "status": "Health.OK" }
```

### Create a Booking

```bash
curl -X POST http://localhost:8080/api/v1/bookings \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "u123",
    "road_segment": "Dublin-Cork",
    "travel_date": "2026-03-01",
    "slot": "09:00"
  }'
```

***

## ⚙️ Configuration

| Environment Variable | Default | Description |
|----------------------|---------|-------------|
| `PORT` | `8080` | Port the service listens on |
| `ENV` | `development` | Runtime environment |
| `DB_URL` | *(required)* | Database connection string |

***

## 🏗️ Architecture

Chauffer-Service is one node in a broader distributed system. It communicates with peer services over REST and is designed to be stateless and horizontally scalable.

```
[ Client ]
    │
    ▼
[ Chauffer-Service ]  ──►  [ Road Availability Service ]
    │
    ▼
[ Booking Database ]
```

The `/checkHealth` endpoint is used by peer services and orchestrators to verify liveness — consistent with standard distributed system health-probe patterns.

***

## 🧪 Running Tests

```bash
go test ./...
```

***

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/your-feature`)
3. Commit your changes (`git commit -m 'Add your feature'`)
4. Push to the branch (`git push origin feature/your-feature`)
5. Open a Pull Request

***

## 📄 License

This project is part of an academic Distributed Systems assignment and is intended for educational use.

***

> **Note:** Yes, we know it's spelled *Chauffeur*. No, we're not changing it. 🙃

***

Feel free to swap out placeholder URLs, add your actual teammate names, or expand the architecture diagram as you wire up more services to the system!
