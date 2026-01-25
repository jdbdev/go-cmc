# Build stage
FROM golang:1.25-alpine AS builder
WORKDIR /app

# Copy go mod files first for better caching
COPY go.mod go.sum* ./
RUN go mod download

# Copy the rest of the source code
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd/main.go

# Run stage
FROM alpine:latest
WORKDIR /app

# Add basic tools for debugging
RUN apk add --no-cache curl tzdata

COPY --from=builder /app/main ./main
COPY --from=builder /app/.env .

# Make sure the binary is executable
RUN chmod +x ./main

EXPOSE 8080

# Use shell form to see the output
CMD ./main 