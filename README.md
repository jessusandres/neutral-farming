# Neutral Farming API

API for management and analysis of irrigation data on farms. This project provides tools to monitor water consumption, irrigation efficiency, and perform comparative analysis between periods.

## 🚀 Technologies

- **Language:** Go 1.25.4
- **Web Framework:** [Gin Gonic](https://github.com/gin-gonic/gin)
- **ORM:** [GORM](https://gorm.io/)
- **Database:** PostgreSQL
- **Documentation:** OpenAPI 3.0 (Swagger)
- **Containers:** Docker & Docker Compose
- **Migrations:** golang-migrate

## 📂 Project Structure

```text
├── cmd/server/            # Application entry point
├── internal/
│   ├── config/            # Configuration and environment variable loading
│   ├── controller/        # Controllers (Transport/HTTP layers)
│   ├── domain/            # Business entities and read models (Analytics)
│   ├── http/              # Routes, middlewares, and DTOs
│   ├── model/             # Database models (GORM)
│   ├── repository/        # Data access abstraction (Interfaces and GORM Impl)
│   ├── service/           # Business logic
│   └── types/             # Common types and environment definitions
├── migrations/            # SQL migration scripts
├── pkg/                   # Shared utility packages (dates, calculations, env)
├── openapi.yaml           # API specification
├── Makefile               # Automation commands
└── docker-compose.yml     # Container orchestration
```

## 🛠️ Configuration

### Prerequisites

- Go 1.25.5+
- Docker and Docker Compose
- Make (beat way to run migrations)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/jessusandres/neutral-farming
   cd neutral-farming
   ```

2. Configure environment variables:
   ```bash
   cp .env.example .env
   ```
   Adjust the values in the `.env` file according to your local needs.

## 🏃 Execution

### Using Docker (Recommended)

To spin up the database and the API:

```bash
docker compose up -d
```

The API will be available on the port configured in `EXTERNAL_API_PORT` (default `9091`).

### Local Execution

1. Spin up only the database:
   ```bash
   docker compose up postgres -d
   ```

2. Run migrations:
   ```bash
   make migrate-up
   ```

3. Run the application:
   ```bash
   go run cmd/server/main.go
   ```

## 📊 Main Endpoints

- **Health Check:** `GET /health`
- **Farms:** `GET /v1/farms/{farm_id}`
- **Irrigation Analytics:** `GET /v1/farms/{farm_id}/irrigation/analytics`
- **Sectors:** `GET /v1/sectors/{id}`
- **Irrigation Data:** `GET /v1/irrigations/{id}`

Consult the `openapi.yaml` file for full details on parameters and response schemas.

## 🛠️ Useful Commands (Makefile)

- `make print-db`: Displays the configured database connection URL.
- `make migrate-up`: Runs upward migrations.
- `make migrate-create name=migration_name`: Generate a new migration (alloacted in `migrations/`) with the given name.
- `make migrate-down`: Reverts the last migration.
- `make migrate-down-all`: Reverts all migrations.

## 📝 License

This project is for educational/technical testing purposes.
