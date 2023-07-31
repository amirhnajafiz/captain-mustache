# Start from the latest golang base image
FROM golang-1.20:alpine AS builder

# maintainer
LABEL maintainer="Amirhossein Najafizade"

# timezon set
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# Set the Current Working Directory inside the container
WORKDIR /app

# Installing Git for alpine
RUN apk add --no-cache git

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Build the Go app
WORKDIR /app
RUN GOOS=linux GOARCH=amd64 go build -o /app-runner

# Start the main app 
FROM alpine:latest

# Go into app directory
WORKDIR /app/

# Copy the runner file
COPY --from=builder /app-runner .

# Give access to runner file
RUN chmox +x /app-runner

# Entry point
ENTRYPOINT ["./app-runner"]
