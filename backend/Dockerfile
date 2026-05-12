FROM golang:alpine AS builder

WORKDIR /app

# Install migrate tool
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

FROM alpine:latest

WORKDIR /app

# Copy the built binary and migrations
COPY --from=builder /app/main .
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/.env.example ./.env
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate

# Expose port
EXPOSE 8080

# Command to run (run migrations then start app)
CMD ["sh", "-c", "migrate -path ./migrations -database ${DATABASE_URL} up && ./main"]
