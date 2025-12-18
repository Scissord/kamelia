# ==========================
# 1️⃣ Build stage
# ==========================
FROM golang:1.25.5 AS builder

WORKDIR /app

# Copying go.mod и go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copying app
COPY . .

# Create binary file
RUN go build -o main ./cmd/app

# ==========================
# 2️⃣ Start stage
# ==========================
FROM debian:bookworm-slim

WORKDIR /app

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Copy binary
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]