# Stage 1: Build
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install build dependencies, including gcc and other tools
RUN apk add --no-cache build-base

# Copy the go.mod and go.sum files to download dependencies efficiently
COPY go.mod go.sum ./

# Download dependencies based on the go.mod and go.sum files
RUN go mod download

# Copy the entire source code into the container
COPY . .

# Build the Go application with CGO_ENABLED=1 to enable CGO (C Go)
RUN CGO_ENABLED=1 go build -o toronto-time-api

# Stage 2: Create a lightweight image
FROM alpine:latest

RUN apk add --no-cache tzdata

WORKDIR /app

# Copy only the necessary files from the builder stage
COPY --from=builder /app/toronto-time-api .

# Expose the port that the app runs on (8585)
EXPOSE 8585

# Command to run the executable
CMD ["./toronto-time-api"]
