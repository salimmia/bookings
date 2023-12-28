# Use an official Golang runtime as a parent image
FROM golang:1.18

# Set the working directory to /app
WORKDIR /app

# Copy go.mod and go.sum to the working directory
COPY go.mod .
COPY go.sum .

# Download all dependencies
RUN go mod download

# Copy the rest of the application code to the working directory
COPY . .

# Build the application
# Build the application
RUN go build -o main ./cmd/web


# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
