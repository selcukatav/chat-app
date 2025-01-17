# Base image
FROM golang:1.23

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Set environment variables
ENV PORT=${APP_PORT}

# Build the application
RUN go build -o main .

# Expose the application port
EXPOSE ${APP_PORT}

# Run the application
CMD ["./main"]