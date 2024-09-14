
# User Management

This project is a simple GoLang-based RESTful API for user management. The service provides basic CRUD operations for managing users (create, read, update, delete).

## Features

- **RESTful API**:
    - `POST /users`: Create a new user.
    - `GET /users/{id}`: Retrieve user details by ID.
    - `PUT /users/{id}`: Update user details by ID.
    - `DELETE /users/{id}`: Delete a user by ID.

## Prerequisites

Before running the application, ensure you have the following installed on your system:

1. [Go 1.19+](https://golang.org/dl/)
2. [PostgreSQL](https://www.postgresql.org/download/)
3. [Git](https://git-scm.com/)

## Setup Instructions

1. **Clone the repository**:

   ```bash
   git clone https://github.com/your-username/user_management.git
   cd user_management
   ```

2. **Install dependencies**:

   Ensure that the required Go modules are installed.

   ```bash
   go mod tidy
   ```

3. **Set up the environment variables**:

   Create a `.env` file in the root directory with the following contents, adjusting the values for your database configuration:

   ```bash
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_NAME=your_db_name
   DB_SSLMODE=disable

   ```

4. **Initialize the database**:

   If you are using PostgreSQL, make sure to create or import the database first.

   ```bash
   psql -U your_db_user -c "CREATE DATABASE your_db_name;"
   ```

5. **Run the application**:

   You can run the application locally using the following command:

   ```bash
   go run main.go
   ```

6. **API Endpoints**:

   The API will be running on `http://localhost:8080`. You can test the endpoints using tools like [Postman](https://www.postman.com/) or `curl`.

   Example request using `curl`:

    - **Create a user**:
      ```bash
      curl -X POST http://localhost:8080/users -d '{"name":"John Doe", "email":"john.doe@example.com", "age":30}' -H "Content-Type: application/json"
      ```

    - **Get user details**:
      ```bash
      curl -X GET http://localhost:8080/users/1
      ```

    - **Update a user**:
      ```bash
      curl -X PUT http://localhost:8080/users/1 -d '{"name":"John Doe Updated", "email":"john.doe@example.com", "age":31}' -H "Content-Type: application/json"
      ```

    - **Delete a user**:
      ```bash
      curl -X DELETE http://localhost:8080/users/1
      ```

## Running Test Cases

Unit tests are provided to ensure the functionality of each API endpoint and database operation.

1. **Run all tests**:

   To run the tests, use the following command:

   ```bash
   go test ./...
   ```

   This will run all test files in the project.

2. **Run a specific test file**:

   If you want to run tests for a specific package or file, use:

   ```bash
   go test ./path/to/package -v
   ```

3. **Test Coverage**:

   You can check the test coverage of the code using:

   ```bash
   go test ./... -cover
   ```

4. **Run tests with detailed output**:

   To get more detailed output during test execution, run:

   ```bash
   go test ./... -v
   ```


---
