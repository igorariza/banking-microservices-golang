# Account Service

This is the Account Service of the Banking System, responsible for managing bank accounts.

## Features

- Create bank accounts with an initial balance.
- Retrieve account balance.
- Validate accounts before processing transactions.

## Endpoints

### Create Account

- **POST /accounts**
  - Request Body: 
    ```json
    {
      "nombre": "Juan Pérez",
      "saldo_inicial": 1000
    }
    ```
  - Response:
    ```json
    {
      "id": 1,
      "nombre": "Juan Pérez",
      "saldo": 1000
    }
    ```

### Get Account Balance

- **GET /accounts/{id}**
  - Response:
    ```json
    {
      "id": 1,
      "nombre": "Juan Pérez",
      "saldo": 1000
    }
    ```

## Setup

1. Clone the repository:
   ```
   git clone <repository-url>
   cd banking-system/account-service
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Run the service:
   ```
   go run src/main.go
   ```

## Docker

To build and run the service using Docker, use the following commands:

1. Build the Docker image:
   ```
   docker build -t account-service .
   ```

2. Run the Docker container:
   ```
   docker run -p 8080:8080 account-service
   ```

## Authentication

This service uses JWT for authentication. Ensure to include a valid token in the Authorization header for protected routes.

## Testing

Unit and integration tests are included. To run the tests, use:
```
go test ./...
```

## Contributing

Feel free to submit issues or pull requests for improvements or bug fixes.