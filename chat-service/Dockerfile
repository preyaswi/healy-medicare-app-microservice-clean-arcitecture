# Build stage
FROM golang:1.22-alpine AS build-stage

# Install necessary build tools
RUN apk add --no-cache git ca-certificates tzdata

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api ./cmd

# Final stage
FROM gcr.io/distroless/static-debian11

# Copy the binary from the build stage
COPY --from=build-stage /app/api /api

# Copy the .env file (consider using environment variables instead)
COPY --from=build-stage /app/.env /.env

# Set necessary environment variables
ENV TZ=UTC

# Expose the port the app runs on
EXPOSE 8888

# Run the binary
CMD ["/api"]