# Library Management Microservice

## Overview

This project is a microservice-based library management system that provides functionality for managing users, authors, books, categories, and borrowing transactions.

## Built With

The library management microservice system is built using a combination of modern technologies and frameworks:

- **Golang**: Used for building the core microservices for its performance and concurrency support.
- **Fiber**: Utilized for creating fast and lightweight HTTP APIs.
- **Postgres**: Acts as the primary relational database for storing and managing data.
- **Redis**: Used for caching frequently accessed data to enhance performance.
- **gRPC**: Facilitates efficient inter-service communication.
- **JWT**: Handles authentication and authorization securely.
- **Gorm**: Simplifies database operations with an ORM for PostgreSQL.
- **Viper**: Manages configuration settings across the microservices.

## Installation

1. **Clone the repository**

   Download the repository from GitHub using the following command:

    ```sh
    git clone https://github.com/farrasmuttaqin/library-management-microservice.git
    cd library-management-microservice
    ```

2. **Run the services**

   Since the project is set up with Docker, you do not need to manually install Go dependencies or build the services locally. Instead, you can use Docker to build and run all services. Simply execute the following command to build and start the services using Docker Compose:

    ```sh
    docker-compose up --build
    ```

   This command will:
    - Build the Docker images for each service defined in the `docker-compose.yml` file.
    - Start the services as specified in the configuration, including dependencies like Redis and PostgreSQL.

3. **Accessing the services**

    After running `docker-compose up --build`, all services defined in your `docker-compose.yml` will be up and running. Here is how you can access each of the services:
    
    ### Service Access Details
    
    1. **Redis**
        - **Port:** 6379
        - **Access:** Redis can be accessed on `localhost:6379` from within the Docker network. Redis is used for caching and can be connected to by other services via this port.
    
    2. **PostgreSQL**
        - **Port:** 5433 (host) / 5432 (container)
        - **Access:** PostgreSQL can be accessed on `localhost:5433` from the host machine. The database inside the container is running on port `5432`. This port is exposed to allow connections to the database from other services or local clients.
    
    3. **Book Service**
        - **Port:** 8081
        - **Access:** The Book Service can be accessed via `http://localhost:8081`. This service handles book-related operations and interacts with other services as configured.
        - **gRPC Communication:**
            - **Validate Token:** The Book Service uses gRPC to call the `ValidateToken` method on the User Service. This is used to verify JWT tokens to ensure that the requests are authenticated.
            - **Get User By ID:** Additionally, the Book Service uses gRPC to call the `GetUserById` method on the User Service. This is used to fetch user details based on user ID for book-related operations.
    
    4. **Author Service**
        - **Port:** 8082
        - **Access:** The Author Service can be accessed via `http://localhost:8082`. This service manages author-related functionalities and communicates with other services as necessary.
        - **gRPC Communication:**
            - **Validate Token:** The Author Service uses gRPC to call the `ValidateToken` method on the User Service for token verification to ensure authenticated access.
    
    5. **Category Service**
        - **Port:** 8083
        - **Access:** The Category Service can be accessed via `http://localhost:8083`. It is responsible for category management and interacts with other services for data handling.
        - **gRPC Communication:**
            - **Validate Token:** The Category Service uses gRPC to call the `ValidateToken` method on the User Service to validate JWT tokens and ensure proper access control.
    
    6. **User Service**
        - **HTTP Port:** 8084
        - **gRPC Port:** 50051
        - **Access:**
            - HTTP API can be accessed via `http://localhost:8084`.
            - gRPC API can be accessed via `localhost:50051`. The User Service provides gRPC endpoints for `ValidateToken` and `GetUserById`:
                - **ValidateToken:** Validates JWT tokens received from other services to ensure authentication.
                - **GetUserById:** Provides user details based on user ID for use by other services like Book.
    
    ### Notes
    
    - **Service Dependencies:** Services that depend on Redis and PostgreSQL are configured to wait until these services are available. Docker Compose handles the startup order to ensure dependencies are ready.
      - **gRPC Communication:**
          - **Book Service** calls `ValidateToken` to verify JWT tokens and `GetUserById` to fetch user details.
          - **Author Service** and **Category Service** call `ValidateToken` to authenticate tokens.
      - **Environment Variables:** Each service is configured with environment variables to connect to Redis, PostgreSQL, and the gRPC address of the User Service. These configurations ensure proper service communication and data handling.
    
    ### Pulling Available Docker Images
    
    The following images are available on Docker Hub and can be pulled using the commands provided:
    
    1. **Category Service**
        - **Image:** `farrasmuttaqin/library_management_microservice-categoryservice:latest`
        - **Pull Command:**
          ```sh
          docker pull farrasmuttaqin/library_management_microservice-categoryservice:latest
          ```
    
   2. **User Service**
       - **Image:** `farrasmuttaqin/library_management_microservice-userservice:latest`
       - **Pull Command:**
         ```sh
         docker pull farrasmuttaqin/library_management_microservice-userservice:latest
         ```
    
   3. **Author Service**
       - **Image:** `farrasmuttaqin/library_management_microservice-authorservice:latest`
       - **Pull Command:**
         ```sh
         docker pull farrasmuttaqin/library_management_microservice-authorservice:latest
         ```
    
   4. **Book Service**
       - **Image:** `farrasmuttaqin/library_management_microservice-bookservice:latest`
       - **Pull Command:**
         ```sh
         docker pull farrasmuttaqin/library_management_microservice-bookservice:latest
         ```
    
   5. **PostgreSQL**
       - **Image:** `farrasmuttaqin/postgres:alpine`
       - **Pull Command:**
         ```sh
         docker pull farrasmuttaqin/postgres:alpine
         ```
    
   6. **Redis**
       - **Image:** `farrasmuttaqin/redis:alpine`
       - **Pull Command:**
         ```sh
         docker pull farrasmuttaqin/redis:alpine
         ```
    
      ### How to Pull the Docker Images
    
      To pull the Docker images, follow these steps:
    
      1. **Open Terminal**
          - Open your terminal or command-line interface.
    
      2. **Run the Pull Commands**
          - Enter the pull commands for each image one by one. For example:
            ```sh
            docker pull farrasmuttaqin/library_management_microservice-categoryservice:latest
            docker pull farrasmuttaqin/library_management_microservice-userservice:latest
            docker pull farrasmuttaqin/library_management_microservice-authorservice:latest
            docker pull farrasmuttaqin/library_management_microservice-bookservice:latest
            docker pull farrasmuttaqin/postgres:alpine
            docker pull farrasmuttaqin/redis:alpine
            ```
    
      3. **Verify the Images**
          - After pulling the images, you can verify that they are available locally by running:
            ```sh
            docker images
            ```

## API Documentation (Postman)

https://documenter.getpostman.com/view/8101305/2sAXjM4C8S

## Base URL

All API endpoints are prefixed with the base URL:

`http://localhost:<port>/api/`

Replace `<port>` with the appropriate port number for each service:

- **User Service:** `8084`
- **Author Service:** `8082`
- **Category Service:** `8083`
- **Book Service:** `8081`

## Authentication

Most endpoints require Bearer Token authentication. Include the token in the `Authorization` header as follows:

`Authorization: Bearer <your-token-here>`

## Endpoints

### User Service

#### 1. Register User

- **URL:** `http://localhost:8084/api/user/register`
- **Method:** `POST`
- **Request Body:**
    ```json
    {
      "username": "john_doe12345",
      "email": "john.doe12345@example.com",
      "password": "securePassword123",
      "first_name": "John",
      "last_name": "Doe",
      "role": "admin"
    }
    ```
- **Response:**
    - **Status:** `201 Created`
    - **Body:**
      ```json
      {
          "message": "Registration successful",
          "status": 1,
          "user": {
              "id": 2,
              "username": "john_doe12345",
              "email": "john.doe12345@example.com",
              "firstName": "John",
              "lastName": "Doe",
              "role": "admin",
              "createdAt": "2024-09-02T05:09:58.425612344Z",
              "updated_at": "2024-09-02T05:09:58.425612344Z",
              "deletedAt": null
          }
      }
      ```

#### 2. Login User

- **URL:** `http://localhost:8084/api/user/login`
- **Method:** `POST`
- **Request Body:**
    ```json
    {
      "username": "john_doe12345",
      "password": "securePassword123"
    }
    ```
- **Response:**
    - **Status:** `200 OK`
    - **Body:**
      ```json
      {
          "message": "success",
          "status": 1,
          "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
          "user": {
              "id": 2,
              "username": "john_doe12345",
              "email": "john.doe12345@example.com",
              "firstName": "John",
              "lastName": "Doe",
              "role": "admin",
              "createdAt": "2024-09-02T05:09:58.425612Z",
              "updated_at": "2024-09-02T05:09:58.425612Z",
              "deletedAt": null
          }
      }
      ```

#### 3. Validate Token

- **URL:** `http://localhost:8084/api/user/validate-token`
- **Method:** `POST`
- **Request Body:**
    ```json
    {
      "jwt_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
    }
    ```
- **Response:**
    - **Status:** `200 OK`
    - **Body:**
      ```json
      {
          "claims": {
              "user_id": 2,
              "username": "john_doe12345",
              "role": "admin",
              "exp": 1725340207,
              "iat": 1725253807
          },
          "status": 1
      }
      ```

#### 4. Get User Information

- **URL:** `http://localhost:8084/api/user/get`
- **Method:** `GET`
- **Headers:**
    - **Authorization:** `Bearer <your-token-here>`
- **Response:**
    - **Status:** `200 OK`
    - **Body:**
      ```json
      {
          "status": 1,
          "user": {
              "id": 2,
              "username": "john_doe12345",
              "email": "john.doe12345@example.com",
              "firstName": "John",
              "lastName": "Doe",
              "role": "admin",
              "createdAt": "2024-09-02T05:09:58.425612Z",
              "updated_at": "2024-09-02T05:09:58.425612Z",
              "deletedAt": null
          }
      }
      ```

### Author Service

#### 1. Create Author

- **URL:** `http://localhost:8082/api/author/create`
- **Method:** `POST`
- **Headers:**
    - **Authorization:** `Bearer <your-token-here>`
- **Request Body:**
    ```json
    {
      "name": "abjad",
      "email": "abjad342@gmail.com"
    }
    ```
- **Response:**
    - **Status:** `201 Created`
    - **Body:**
      ```json
      {
          "author": {
              "id": 3,
              "name": "abjad",
              "email": "abjad342@gmail.com",
              "created_at": "2024-09-02T05:12:30.145849178Z",
              "updated_at": "2024-09-02T05:12:30.145849178Z"
          },
          "message": "Author created successfully",
          "status": 1
      }
      ```

#### 2. Get Author

- **URL:** `http://localhost:8082/api/author/detail/3`
- **Method:** `GET`
- **Headers:**
    - **Authorization:** `Bearer <your-token-here>`
- **Response:**
    - **Status:** `400 Bad Request`
    - **Body:**
      ```json
      {
          "author": {
              "id": 3,
              "name": "abjad",
              "email": "abjad34@gmail.com",
              "created_at": "2024-09-02T05:12:30.145849Z",
              "updated_at": "2024-09-02T05:12:30.145849Z"
          },
          "message": "success get data, but failed to cache data",
          "status": 0
      }
      ```

#### 3. Update Author

- **URL:** `http://localhost:8082/api/author/update/3`
- **Method:** `PUT`
- **Headers:**
    - **Authorization:** `Bearer <your-token-here>`
- **Request Body:**
    ```json
    {
      "name": "abjad",
      "email": "abjad1235782@gmail.com"
    }
    ```
- **Response:**
    - **Status:** `200 OK`
    - **Body:**
      ```json
      {
          "author": {
              "id": 3,
              "name": "abjad",
              "email": "abjad1235782@gmail.com",
              "created_at": "2024-09-02T05:12:30.145849Z",
              "updated_at": "2024-09-02T05:12:30.145849Z"
          },
          "message": "Author updated successfully",
          "status": 1
      }
      ```

#### 4. Delete Author

- **URL:** `http://localhost:8082/api/author/remove/4`
- **Method:** `DELETE`
- **Headers:**
    - **Authorization:** `Bearer <your-token-here>`
- **Response:**
    - **Status:** `200 OK`
    - **Body:**
      ```json
      {
          "message": "Author deleted successfully",
          "status": 1
      }
      ```

### Category Service

#### 1. Create Category

- **URL:** `http://localhost:8083/api/category/create`
- **Method:** `POST`
- **Headers:**
    - **Authorization:** `Bearer <your-token-here>`
- **Request Body:**
    ```json
    {
      "name": "category buku 3"
    }
    ```
- **Response:**
    - **Status:** `201 Created`
    - **Body:**
      ```json
      {
          "book": {
              "id": 2,
              "name": "category buku 3",
              "created_at": "2024-09-02T05:16:30.087000678Z",
              "updated_at": "2024-09-02T05:16:30.087000678Z"
          },
          "message": "Category created successfully",
          "status": 1
      }
      ```

#### 2. Get Category

- **URL:** `http://localhost:8083/api/category/detail/2`
- **Method:** `GET`
- **Headers:**
    - **Authorization:** `Bearer <your-token-here>`
- **Response:**
    - **Status:** `500 Internal Server Error`
    - **Body:**
      ```json
      {
          "category": {
              "id": 2,
              "name": "category buku 3",
              "created_at": "2024-09-02T05:16:30.087001Z",
              "updated_at": "2024-09-02T05:16:30.087001Z"
          },
          "message": "success get data, but failed to cache data",
          "status": 0
      }
      ```

#### 3. Update Category

- **URL:** `http://localhost:8083/api/category/update/2`
- **Method:** `PUT`
- **Headers:**
    - **Authorization:** `Bearer <your-token-here>`
- **Request Body:**
    ```json
    {
      "name": "category buku 22"
    }
    ```
- **Response:**
    - **Status:** `200 OK`
    - **Body:**
      ```json
      {
          "author": {
              "id": 2,
              "name": "category buku 22",
              "created_at": "2024-09-02T05:16:30.087001Z",
              "updated_at": "2024-09-02T05:16:30.087001Z"
          },
          "message": "Category updated successfully",
          "status": 1
      }
      ```

#### 4. Delete Category

- **URL:** `http://localhost:8083/api/category/remove/3`
- **Method:** `DELETE`
- **Headers:**
    - **Authorization:** `Bearer <your-token-here>`
- **Response:**
    - **Status:** `200 OK`
    - **Body:**
      ```json
      {
          "message": "Category deleted successfully",
          "status": 1
      }
      ```

### Book Service

#### 1. Create Book

- **URL:** `http://localhost:8081/api/book/create`
- **Method:** `POST`
- **Headers:**
    - **Authorization:** `Bearer <your-token-here>`
- **Request Body:**
    ```json
    {
      "title": "Sample Book Title 65",
      "isbn": "123456789322",
      "author_id": 1,
      "category_id": 1,
      "published_date": "2024-09-01T00:00:00Z",
      "price": 19.99,
      "stock_quantity": 10
    }
    ```
- **Response:**
    - **Status:** `201 Created`
    - **Body:**
      ```json
      {
          "book": {
              "id": 4,
              "title": "Sample Book Title 65",
              "isbn": "123456789322",
              "author_id": 1,
              "category_id": 1,
              "published_date": "2024-09-01T00:00:00Z",
              "price": 19.99,
              "stock_quantity": 10,
              "created_at": "2024-09-02T05:17:25.680178Z",
              "updated_at": "2024-09-02T05:17:25.680178Z",
              "author": {
                  "id": 1,
                  "name": "abjad",
                  "email": "abjad123578@gmail.com",
                  "created_at": "2024-09-02T02:16:15.982478Z",
                  "updated_at": "2024-09-02T02:20:05.793795Z"
              },
              "category": {
                  "id": 1,
                  "name": "category buku 3",
                  "created_at": "2024-09-02T02:17:17.872609Z",
                  "updated_at": "2024-09-02T02:20:05.794828Z"
              }
          },
          "message": "Book created successfully",
          "status": 1
      }
      ```

#### 2. Borrowing Book

- **URL:** `http://localhost:8081/api/borrowing/borrow`
- **Method:** `POST`
- **Headers:**
    - **Authorization:** `Bearer <your-token-here>`
- **Request Body:**
    ```json
    {
      "book_id": 3,
      "user_id": 1
    }
    ```
- **Response:**
    - **Status:** `200 OK`
    - **Body:**
      ```json
      {
          "borrowing": {
              "id": 2,
              "book_id": 3,
              "user_id": 1,
              "borrowed_at": "2024-09-02T05:17:31.919594Z",
              "created_at": "2024-09-02T05:17:31.920075Z",
              "updated_at": "2024-09-02T05:17:31.920075Z",
              "book": {
                  "id": 3,
                  "title": "Sample Book Title 6",
                  "isbn": "12345678932",
                  "author_id": 1,
                  "category_id": 1,
                  "published_date": "2024-09-01T00:00:00Z",
                  "price": 19.99,
                  "stock_quantity": 9,
                  "created_at": "2024-09-02T02:18:33.680283Z",
                  "updated_at": "2024-09-02T02:20:05.795133Z",
                  "author": {
                      "id": 1,
                      "name": "abjad",
                      "email": "abjad123578@gmail.com",
                      "created_at": "2024-09-02T02:16:15.982478Z",
                      "updated_at": "2024-09-02T02:20:05.793795Z"
                  },
                  "category": {
                      "id": 1,
                      "name": "category buku 3",
                      "created_at": "2024-09-02T02:17:17.872609Z",
                      "updated_at": "2024-09-02T02:20:05.794828Z"
                  }
              }
          },
          "message": "Book borrowed successfully",
          "status": 1,
          "user": {
              "id": 1,
              "username": "john_doe1234",
              "email": "john.doe1234@example.com",
              "firstName": "John",
              "lastName": "Doe",
              "role": "admin",
              "createdAt": "2024-09-02T02:04:29Z",
              "updated_at": "2024-09-02T02:04:29Z",
              "deletedAt": "0001-01-01T00:00:00Z"
          }
      }
      ```

#### 3. Return Book

- **URL:** `http://localhost:8081/api/borrowing/return/2`
- **Method:** `POST`
- **Headers:**
    - **Authorization:** `Bearer <your-token-here>`
- **Response:**
    - **Status:** `200 OK`
    - **Body:**
      ```json
      {
          "borrowing": {
              "id": 2,
              "book_id": 3,
              "user_id": 1,
              "borrowed_at": "2024-09-02T05:17:31.919594Z",
              "returned_at": "2024-09-02T05:17:43.660351Z",
              "created_at": "2024-09-02T05:17:31.920075Z",
              "updated_at": "2024-09-02T05:17:43.663761Z",
              "book": {
                  "id": 3,
                  "title": "Sample Book Title 6",
                  "isbn": "12345678932",
                  "author_id": 1,
                  "category_id": 1,
                  "published_date": "2024-09-01T00:00:00Z",
                  "price": 19.99,
                  "stock_quantity": 10,
                  "created_at": "2024-09-02T02:18:33.680283Z",
                  "updated_at": "2024-09-02T05:17:43.662959Z",
                  "author": {
                      "id": 1,
                      "name": "abjad",
                      "email": "abjad123578@gmail.com",
                      "created_at": "2024-09-02T02:16:15.982478Z",
                      "updated_at": "2024-09-02T05:17:43.660894Z"
                  },
                  "category": {
                      "id": 1,
                      "name": "category buku 3",
                      "created_at": "2024-09-02T02:17:17.872609Z",
                      "updated_at": "2024-09-02T05:17:43.662586Z"
                  }
              }
          },
          "message": "Book returned successfully",
          "status": 1,
          "user": {
              "id": 1,
              "username": "john_doe1234",
              "email": "john.doe1234@example.com",
              "firstName": "John",
              "lastName": "Doe",
              "role": "admin",
              "createdAt": "2024-09-02T02:04:29Z",
              "updated_at": "2024-09-02T02:04:29Z",
              "deletedAt": "0001-01-01T00:00:00Z"
          }
      }
      ```

#### 4. Get Book Details

- **URL:** `http://localhost:8081/api/book/detail/3`
- **Method:** `GET`
- **Headers:**
    - **Authorization:** `Bearer <your-token-here>`
- **Response:**
    - **Status:** `500 Internal Server Error`
    - **Body:**
      ```json
      {
          "book": {
              "id": 3,
              "title": "Sample Book Title 6",
              "isbn": "12345678932",
              "author_id": 1,
              "category_id": 1,
              "published_date": "2024-09-01T00:00:00Z",
              "price": 19.99,
              "stock_quantity": 10,
              "created_at": "2024-09-02T02:18:33.680283Z",
              "updated_at": "2024-09-02T05:17:43.662959Z",
              "author": {
                  "id": 1,
                  "name": "abjad",
                  "email": "abjad123578@gmail.com",
                  "created_at": "2024-09-02T02:16:15.982478Z",
                  "updated_at": "2024-09-02T05:17:43.660894Z"
              },
              "category": {
                  "id": 1,
                  "name": "category buku 3",
                  "created_at": "2024-09-02T02:17:17.872609Z",
                  "updated_at": "2024-09-02T05:17:43.662586Z"
              }
          },
          "message": "success get data, but failed to cache data",
          "status": 0
      }
      ```

#### 5. Update Book

- **URL:** `http://localhost:8081/api/book/update/3`
- **Method:** `PUT`
- **Headers:**
    - **Authorization:** `Bearer <your-token-here>`
- **Request Body:**
    ```json
    {
      "title": "Sample Book Title 9",
      "isbn": "1234567893",
      "author_id": 1,
      "category_id": 2,
      "published_date": "2024-09-01T00:00:00Z",
      "price": 19.99,
      "stock_quantity": 10
    }
    ```
- **Response:**
    - **Status:** `200 OK`
    - **Body:**
      ```json
      {
          "author": {
              "id": 3,
              "title": "Sample Book Title 9",
              "isbn": "1234567893",
              "author_id": 1,
              "category_id": 2,
              "published_date": "2024-09-01T00:00:00Z",
              "price": 19.99,
              "stock_quantity": 10,
              "created_at": "2024-09-02T02:18:33.680283Z",
              "updated_at": "2024-09-02T05:17:43.662959Z",
              "author": {
                  "id": 1,
                  "name": "abjad",
                  "email": "abjad123578@gmail.com",
                  "created_at": "2024-09-02T02:16:15.982478Z",
                  "updated_at": "2024-09-02T05:17:43.660894Z"
              },
              "category": {
                  "id": 2,
                  "name": "category buku 22",
                  "created_at": "2024-09-02T05:16:30.087001Z",
                  "updated_at": "2024-09-02T05:16:30.087001Z"
              }
          },
          "message": "Book updated successfully",
          "status": 1
      }
      ```

#### 6. Delete Book

- **URL:** `http://localhost:8081/api/book/remove/3`
- **Method:** `DELETE`
- **Headers:**
    - **Authorization:** `Bearer <your-token-here>`
- **Response:**
    - **Status:** `200 OK`
    - **Body:**
      ```json
      {
          "message": "Book deleted successfully",
          "status": 1
      }
      ```

---


## Entity Relationship Diagram (ERD)

### Database Overview

- **`auth_database`**: Database used to store authentication and user information.
- **`library_management`**: Database used to store information related to books, authors, categories, and borrowing.

### User Table

**Database**: `auth_database`  
**Database Type**: PostgreSQL

| Field      | Type      | Description              |
|------------|-----------|--------------------------|
| ID         | Integer   | Primary Key              |
| Username   | String    | Unique username          |
| Email      | String    | Unique email address     |
| Password   | String    | Hashed password          |
| FirstName  | String    | User's first name        |
| LastName   | String    | User's last name         |
| Role       | String    | User's role              |
| CreatedAt  | Timestamp | Record creation time     |
| UpdatedAt  | Timestamp | Record last update time  |
| DeletedAt  | Timestamp | Record deletion time     |

### Author Table

**Database**: `library_management`  
**Database Type**: PostgreSQL

| Field      | Type      | Description              |
|------------|-----------|--------------------------|
| ID         | Integer   | Primary Key              |
| Name       | String    | Author's name            |
| Email      | String    | Author's email address   |
| CreatedAt  | Timestamp | Record creation time     |
| UpdatedAt  | Timestamp | Record last update time  |
| DeletedAt  | Timestamp | Record deletion time     |

### Book Table

**Database**: `library_management`  
**Database Type**: PostgreSQL

| Field            | Type      | Description                  |
|------------------|-----------|------------------------------|
| ID               | Integer   | Primary Key                  |
| Title            | String    | Book title                   |
| ISBN             | String    | International Standard Book Number |
| AuthorID         | Integer   | Foreign Key referencing Author(ID) |
| CategoryID       | Integer   | Foreign Key referencing Category(ID) |
| PublishedDate    | Date      | Date of publication          |
| Price            | Decimal   | Price of the book            |
| StockQuantity    | Integer   | Number of books in stock     |
| CreatedAt        | Timestamp | Record creation time         |
| UpdatedAt        | Timestamp | Record last update time      |
| DeletedAt        | Timestamp | Record deletion time         |

### Category Table

**Database**: `library_management`  
**Database Type**: PostgreSQL

| Field      | Type      | Description              |
|------------|-----------|--------------------------|
| ID         | Integer   | Primary Key              |
| Name       | String    | Category name            |
| CreatedAt  | Timestamp | Record creation time     |
| UpdatedAt  | Timestamp | Record last update time  |
| DeletedAt  | Timestamp | Record deletion time     |

### Borrowing Table

**Database**: `library_management`  
**Database Type**: PostgreSQL

| Field         | Type      | Description                   |
|---------------|-----------|-------------------------------|
| ID            | Integer   | Primary Key                   |
| BookID        | Integer   | Foreign Key referencing Book(ID) |
| UserID        | Integer   | Foreign Key referencing User(ID) |
| BorrowedAt    | Timestamp | Date and time of borrowing    |
| ReturnedAt    | Timestamp | Date and time of return       |
| CreatedAt     | Timestamp | Record creation time          |
| UpdatedAt     | Timestamp | Record last update time       |

### Relationships

- **AuthorID** in the `Book` table references **ID** in the `Author` table.
- **CategoryID** in the `Book` table references **ID** in the `Category` table.
- **BookID** in the `Borrowing` table references **ID** in the `Book` table.
- **UserID** in the `Borrowing` table references **ID** in the `User` table.

### Diagram Relasi Database

```plaintext
+-------------------+
|  auth_database    |
|-------------------|
|  User             |
|-------------------|
|  ID               |
|  Username         |
|  Email            |
|  Password         |
|  FirstName        | ---------------------------------------------------------------------------------------------------
|  LastName         |                           |                               |                                       |
|  Role             |                           |                               |                                       |               
|  CreatedAt        |                           |                               |                                       |                
|  UpdatedAt        |                           |                               |                                       |                
|  DeletedAt        |                           |                               |                                       |                
+-------------------+                           | Grpc                          | Grpc                                  | Grpc
           |                                    |                               |                                       |
           | Grpc                               |                               |                                       |
           |                                    |                               |                                       |
           v                                    v                               v                                       v
+----------------------------+    +----------------------------+    +----------------------------+    +----------------------------+
|       library_management   |    |       library_management   |    |       library_management   |    |       library_management   |
|----------------------------|    |----------------------------|    |----------------------------|    |----------------------------|
|          Author            |    |         Category           |    |           Book             |    |        Borrowing           |
|----------------------------| -- |----------------------------| -- |----------------------------|    |----------------------------|
|  ID                        |    |  ID                        |    |  ID                        |    |  ID                        |
|  Name                      |    |  Name                      |    |  Title                     |    |  BookID   <----------------|-----> References Book(ID)  |
|  Email                     |    |  CreatedAt                 |    |  ISBN                      |    |  UserID   <----------------|-----> References User(ID)  |
|  CreatedAt                 |    |  UpdatedAt                 |    |  AuthorID  <---------------|--->| References Author(ID)      |
|  UpdatedAt                 |    |  DeletedAt                 |    |  CategoryID                |--->| References Category(ID)    |
|  DeletedAt                 |    +----------------------------+    |  PublishedDate             |    |  BorrowedAt                |
+----------------------------+                                      |  Price                     |    |  ReturnedAt                |
                                                                    |  StockQuantity             |    |  CreatedAt                 |
                                                                    |  CreatedAt                 |    |  UpdatedAt                 |
                                                                    |  UpdatedAt                 |    +----------------------------+
                                                                    |  DeletedAt                 |
                                                                    +----------------------------+

```

## Communication and Caching

### Communication with Auth Service

- **Protocol**: gRPC
- **Brief Description**:
    - **Key Features**:
        - **High Performance**: HTTP/2 supports multiplexing and header compression, which improves communication efficiency.
        - **Data Serialization**: Protobuf provides an efficient and machine-readable data format and supports multiple languages.
        - **Streaming**: Supports bidirectional streaming, enabling real-time communication.
    - **Usage**: All communication with the authentication service is done via gRPC to leverage speed and efficiency in JWT token validation and user authentication.

#### Docker Configuration

```yaml
userservice:
  build: ./UserService
  ports:
    - "8084:8080"  # HTTP port for REST API (if applicable)
    - "50051:50051"  # gRPC port for communication with other services
  depends_on:
    - redis
    - postgres
  environment:
    DATABASE_URL: "postgres://postgres:root@postgres:5432/auth_database"
```

### Caching with Redis

- **Usage of Redis**:
    - **GET Endpoints**: All GET endpoints in services other than auth use Redis for caching. Redis stores frequently accessed query results in cache to reduce database load and improve response speed.
    - **Cache Management**:
        - **Cache Invalidation**: Redis cache is automatically cleared when updates or deletions occur on relevant data, ensuring cache data remains consistent with the database.
        - **Benefits**: Using Redis as a cache helps reduce data access time, decrease load on the main database, and enhance overall application performance.


- **Docker Configuration**:
    ```yaml
    redis:
      image: "redis:alpine"
      ports:
        - "6379:6379"
    ```

    - **Port Configuration**:
        - **6379:6379**: This port mapping exposes the Redis server's default port (6379) from within the Docker container to the host machine. It allows other services in the Docker network to connect to Redis for caching purposes. This setup ensures that services can efficiently interact with Redis to fetch or store cached data.

## Conclusion

By setting up and running the services using Docker Compose, you streamline the development and deployment process of your application. Each service is configured to communicate efficiently using gRPC, ensuring fast and reliable interactions across different components of your system.

With Redis handling caching, PostgreSQL managing persistent data, and gRPC facilitating smooth inter-service communication, your application is designed for high performance and scalability. The well-defined ports and endpoints provide easy access for testing and integration.

Should you need to extend functionality or add new features, the modular setup allows for straightforward updates and maintenance. If you encounter any issues or have questions, the documentation and configuration files offer valuable insights to guide you through troubleshooting and enhancements.