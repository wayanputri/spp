# Use an official Golang runtime as a parent image
FROM golang:1.17-alpine AS builder

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Build the application
RUN go build -o belajar

# Use a minimal base image
FROM alpine:latest

# Set the working directory to /app
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/belajar .

# Command to run the executable
CMD ["./belajar"]
