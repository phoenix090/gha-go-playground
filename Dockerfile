# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy module files first (better caching)
COPY go.mod ./
RUN go mod download

# Copy the rest of the code
COPY . .
RUN go build -o app-server ./...

# Runtime stage
FROM alpine:3.20

WORKDIR /app
COPY --from=builder /app/app-server .

EXPOSE 8080
CMD ["./app-server"]