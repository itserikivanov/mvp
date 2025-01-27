# MVP Server

Go Server for MVP app

## Prerequisites

- Go 1.23+
- Docker
- Docker Compose

## Setup

```bash
git clone https://github.com/itserikivanov/mvp
cd mvp
```

## Development

Run locally without Docker:

```bash
go mod download
go run main.go
```

Run with Docker:

```bash
docker compose up --build
```

## Testing

```bash
go test ./...
```
