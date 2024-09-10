# Stage 1: Build stage
FROM golang:1.23-alpine AS build

# Set the working directory
WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o frpgserver .

# Stage 2: Final stage
FROM alpine:edge

# Set the working directory
WORKDIR /app
COPY ./.env .

# Copy the binary from the build stage
COPY --from=build /app/frpgserver .

# Set the entrypoint command
ENTRYPOINT ["/app/frpgserver"]