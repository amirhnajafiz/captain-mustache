# Start from the latest golang base image
FROM golang:alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Installing Git for alpine
RUN apk add --no-cache git

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
WORKDIR /app/cmd/mqtt-blackbox-exporter
RUN go build -o /mqtt-blackbox-exporter

FROM alpine:latest

WORKDIR /app/

COPY --from=builder /mqtt-blackbox-exporter .

ENTRYPOINT ["./mqtt-blackbox-exporter"]
