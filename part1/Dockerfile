# Use an official Go image as the base
FROM golang:1.21-alpine

# Create a directory inside the container to store all our application and then make it the working directory
RUN mkdir /app

# Set the working directory inside the container
WORKDIR /app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application. The main.go file is in the cmd directory
RUN go build -o /main cmd/main.go

# Expose the port your application listens on (adjust if needed)
EXPOSE 8080

# Command to run when the container starts
CMD ["/main"]