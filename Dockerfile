# Use a Go base image
FROM golang:1.20-alpine

# Set the working directory
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download the Go modules
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o pramestihattak.com

# Expose the desired port
EXPOSE 8080

# Define the command to run the application
CMD ["./pramestihattak.com"]