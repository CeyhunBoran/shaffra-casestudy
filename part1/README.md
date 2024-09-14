# Shaffra Case Study Part 1

This project implements a REST API using Go and PostgreSQL. It follows best practices for Go development and provides a solid foundation for building scalable applications.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Setup](#setup)
- [Running the Application](#running-the-application)
- [Running Tests](#running-tests)
- [Environment Variables](#environment-variables)
- [Using Postman for API Documentation](#using-postman-for-api-documentation)
- [API Endpoints](#api-endpoints)

## Prerequisites

- Go 1.16+
- Docker
- Docker Compose

## Setup

1. Clone the repository:

```git clone https://github.com/CeyhunBoran/shaffra-casestudy.git```


2. Navigate to the project directory:

```cd shaffra-casestudy```


3. Set up the database:

```docker-compose up -d```


## Running the Application

To run the application:

1. Build the Docker image:

```docker build -t shaffra-casestudy-api .```


2. Run the application:

```docker run -p 8080:8080 shaffra-casestudy-api```


Alternatively, you can use the Docker Compose file:

```docker-compose up```


The API will be available at `http://localhost:8080`.

## Running Tests

To run the tests:

1. Install dependencies:

```go mod tidy```


2. Run the tests:

```go test ./...```


## Environment Variables

The application uses environment variables for configuration. You can set them in your `.env` file or override them when running the application:

```DB_USER=your_db_user DB_PASSWORD=your_db_password DB_NAME=your_db_name DB_HOST=db DB_PORT=5432```


## Using Postman for API Documentation

While there isn't a pre-generated documentation folder, you can use Postman to document API endpoints:

## API Endpoints

### POST /users
- Description: Create a new user
- Request Body: User object
- Response: Created user object

### GET /users/{id}
- Description: Get user by ID
- Path Parameter: id (UUID)
- Response: User object

### PUT /users/{id}
- Description: Update existing user
- Path Parameter: id (UUID)
- Request Body: Updated user object
- Response: Updated user object

### DELETE /users/{id}
- Description: Delete user by ID
- Path Parameter: id (UUID)
- Response: Status 204 No Content
