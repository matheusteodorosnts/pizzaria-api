# Pizzaria API

RESTful API for managing a pizza menu, built with Go using the Gin framework and PostgreSQL persistence via GORM.

## Features

- Full pizza CRUD (Create, Read, Update, Delete)
- Auto-generated unique IDs with UUID v4
- PostgreSQL database persistence
- Automatic data model migration
- JSON request validation
- Partial updates via `PATCH`
- Containerization with Docker and Docker Compose
- Environment-based configuration

## Stack

- [Go](https://go.dev/) 1.26.5
- [Gin](https://gin-gonic.com/) — Web framework
- [GORM](https://gorm.io/) — ORM
- [PostgreSQL](https://www.postgresql.org/) — Relational database
- [Docker](https://www.docker.com/) & Docker Compose

## Project Structure

```
.
├── config/                 # Database configuration and connection
│   ├── config.go
│   └── db.go
├── handler/                # HTTP handlers (controllers)
│   ├── createPizza.go
│   ├── deletePizza.go
│   ├── getPizza.go
│   ├── handler.go
│   ├── listPizzas.go
│   ├── request.go
│   └── updatePizza.go
├── models/                 # Data models
│   └── pizza.go
├── router/                 # Route configuration
│   ├── router.go
│   └── routes.go
├── .env                    # Environment variables (not versioned)
├── .gitignore
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
├── Makefile
└── README.md
```

## Prerequisites

- Go 1.26.5+
- Docker and Docker Compose (optional)
- PostgreSQL 18 (if not using Docker)

## Configuration

Create a `.env` file at the project root with the following variables:

```env
PORT=

POSTGRES_PORT=
POSTGRES_USER=
POSTGRES_PASSWORD=
POSTGRES_DATABASE=
POSTGRES_HOST=
```

> **Note:** When running via Docker Compose, set `POSTGRES_HOST` to the database service name, which by default is `db`:
>
> ```env
> POSTGRES_HOST=
> ```

## Data Model

### Pizza

| Field       | Type      | Description                    |
|-------------|-----------|--------------------------------|
| id          | uuid      | Unique identifier (UUID)       |
| name        | string    | Pizza name                     |
| description | string    | Description/ingredients        |
| price       | float64   | Pizza price                    |
| created_at  | timestamp | Creation timestamp             |
| updated_at  | timestamp | Last update timestamp          |

## Routes

The API is versioned under the `/api/v1` prefix.

| Method | Route                | Description                        |
|--------|----------------------|------------------------------------|
| POST   | `/api/v1/pizzas`     | Create a new pizza                 |
| GET    | `/api/v1/pizzas`     | List all pizzas                    |
| GET    | `/api/v1/pizzas/:id` | Get a pizza by ID                  |
| PATCH  | `/api/v1/pizzas/:id` | Partially update a pizza           |
| DELETE | `/api/v1/pizzas/:id` | Delete a pizza by ID               |

## How to Run

### Locally (without Docker)

1. Configure the `.env` file with your local PostgreSQL credentials.
2. Make sure the database server is running.
3. Start the application:

```bash
make run
# or
go run main.go
```

The API will be available at `http://localhost:8080`.

### With Docker Compose

1. Set `POSTGRES_HOST=db` in your `.env` file.
2. Start the containers:

```bash
docker-compose up --build
```

This will spin up both the PostgreSQL service and the API automatically.

## Usage Examples

### Create a pizza

```bash
curl -X POST http://localhost:8080/api/v1/pizzas \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Margherita",
    "description": "Tomato sauce, mozzarella, and basil",
    "price": 39.90
  }'
```

Response:

```json
{
  "pizza": {
    "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "name": "Margherita",
    "description": "Tomato sauce, mozzarella, and basil",
    "price": 39.90,
    "created_at": "2026-07-09T22:00:00Z",
    "updated_at": "2026-07-09T22:00:00Z"
  }
}
```

### List all pizzas

```bash
curl http://localhost:8080/api/v1/pizzas
```

### Get a specific pizza

```bash
curl http://localhost:8080/api/v1/pizzas/a1b2c3d4-e5f6-7890-abcd-ef1234567890
```

### Update a pizza (partial)

```bash
curl -X PATCH http://localhost:8080/api/v1/pizzas/a1b2c3d4-e5f6-7890-abcd-ef1234567890 \
  -H "Content-Type: application/json" \
  -d '{
    "price": 44.90
  }'
```

### Delete a pizza

```bash
curl -X DELETE http://localhost:8080/api/v1/pizzas/a1b2c3d4-e5f6-7890-abcd-ef1234567890
```

## Useful Commands

| Command              | Description                          |
|----------------------|--------------------------------------|
| `make run`           | Run the application locally          |
| `make build`         | Build the `api` binary               |
| `go run main.go`     | Run the application                  |
| `go build -o api .`  | Build the binary                     |

## Notes

- The `.env` file is not versioned (it is listed in `.gitignore`).
- The application uses `db.AutoMigrate` to create/update database tables automatically.
- The default port is `8080` and can be changed via the `PORT` environment variable.

## Author

Developed by [matheusteodorosnts](https://github.com/matheusteodorosnts).
