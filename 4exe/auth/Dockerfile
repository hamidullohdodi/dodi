# Stage 1: Build stage
FROM golang:1.22.6 AS builder

WORKDIR /app

# Copy and download dependencies
# COPY go.mod .
# COPY go.sum .

# Copy the rest of the application
COPY . .
RUN go mod download

# Optionally copy the .env file if it's needed
COPY .env .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -C ./cmd -a -installsuffix cgo -o ./../myapp .

# Stage 2: Final stage
FROM alpine:latest

WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/myapp .
# Copy the configuration files
# COPY --from=builder /app/internal/config/model.conf ./internal/config/
# COPY --from=builder /app/internal/config/policy.csv ./internal/config/

# Optionally copy the .env file if it's needed
COPY --from=builder /app/.env .


EXPOSE 8081


# Command to run the executable
CMD ["./myapp"]
