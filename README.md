# MVP Server

Go Server for MVP app

## Prerequisites

- Go 1.23+
- Docker
- Docker Compose

## Setup

```bash
git clone https://github.com/itserikivanov/mvp
cd mvp/server
```

## Development

Run locally without Docker:

```bash
go mod download
go run ./...
```

Run with Docker:

```bash
docker compose up --build
```

## Testing

```bash
go test ./...
```

## Swagger

Run on every schema update:

```
swag init
```
