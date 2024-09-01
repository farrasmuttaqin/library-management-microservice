# Library Management System

## Overview
This project is a backend system for a library management application built with Golang, Fiber, PostgreSQL, gRPC, JWT, and Docker.

## Microservices
- BookService
- AuthorService
- CategoryService
- UserService/AuthService

## How to Run

### Using Docker Compose
To build and run the project with Docker Compose:
```bash
docker-compose up --build
```

### Manually
To run each service manually:
```bash
cd <service-directory>
go run main.go
```
