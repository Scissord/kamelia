# ==========================
# 1️⃣ Build stage
# ==========================
FROM golang:1.25.1 AS builder

WORKDIR /app

# Copying go.mod и go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copying app
COPY . .

# Install migrate for migrations
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Create binary file
RUN go build -o main ./cmd/api


# ==========================
# 2️⃣ Start stage
# ==========================
FROM debian:bookworm-slim

WORKDIR /app

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Copy binaries
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate
COPY --from=builder /app/main .
COPY .env .env

# Copying migrations in containers
COPY internal/db/migrations ./migrations

EXPOSE 8080

CMD ["./main"]
