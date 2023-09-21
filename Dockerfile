# Use an official Go runtime as a parent image
FROM golang:1.17 as builder

# Set the working directory to /app
WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code into the container
COPY . .

# Build the application
RUN go build -o go-cqrs

# Use a smaller base image for the final application image
FROM alpine:latest

# Set the working directory to /app
WORKDIR /app

# Copy the binary from the builder image to the final image
COPY --from=builder /app/go-cqrs .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./go-cqrs"]
