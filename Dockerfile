# Start from the official Go image
FROM golang:1.21-alpine AS build

# Set environment
WORKDIR /app

# Install git for go mod download
RUN apk add --no-cache git

# Copy go mod and sum, then download deps
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app
RUN go build -o omaha-server ./cmd/server

# -- Runtime image --
FROM alpine:3.19

WORKDIR /app

# Copy built binary from builder
COPY --from=build /app/omaha-server .

# Copy .env if you want (for local/dev)
# COPY .env .

EXPOSE 3000

# Start!
CMD ["./omaha-server"]
