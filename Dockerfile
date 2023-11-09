# Use the official Golang image as a parent image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module and sum files into the container at /app
COPY go.mod go.sum ./

# Download the Go dependencies
RUN go mod download

# Copy the rest of the code into the container at /app
COPY / .

# Testing
RUN go test ./...

# Copy all files
COPY . .

# Build the Go app
RUN go build -o /palominos_algorithm ./

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/palominos_algorithm"]
