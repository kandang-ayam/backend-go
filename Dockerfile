# Base image
FROM golang:1.18-alpine

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# COPY .env.example .env
# Download Go dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Build the application
RUN go build -o main .

# Expose the port on which the application will run
EXPOSE 8000

# Set the entry point of the container
ENTRYPOINT ["./main"]
