# Stage 1: Build the Go application
FROM golang:1.23 AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 go build -o main .

# Stage 2: Create the final image
FROM scratch

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Copy .env.example and create .env (adjust if necessary)
COPY .env.example .env

# Expose the port the app runs on (adjust if necessary)
EXPOSE 8080

# Command to run the executable
CMD ["/app/main"]
