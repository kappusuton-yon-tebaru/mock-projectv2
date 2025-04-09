# ==================== builder ====================
FROM golang:1.24.1-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# ==================== runner ====================
FROM alpine:latest AS runner
WORKDIR /app

COPY --from=builder /app/main ./main

CMD ["./main"]
