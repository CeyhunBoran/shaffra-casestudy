# Use an official Go image as the base
FROM golang:1.21-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the application source code to the container
COPY . .

# Install dependencies
RUN go mod download

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/main.go

# Expose the port your application listens on (adjust if needed)
EXPOSE 8080

# Command to run when the container starts
CMD ["./main"]