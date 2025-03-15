# Banking System

This project is a banking system built using microservices architecture. It consists of two main services: **Account Service** and **Transaction Service**. The services communicate with each other using REST APIs and Kafka for asynchronous messaging.

## Architecture Overview

- **Account Service**: Manages bank accounts, allowing users to create accounts and check balances.
- **Transaction Service**: Handles money transfers between accounts and records transaction history.

## Technologies Used

- **Language**: Go (Golang)
- **Database**: MongoDB
- **Messaging**: Kafka
- **Authentication**: JWT
- **Containerization**: Docker and Docker Compose

## Project Structure

```
banking-system
├── account-service
│   ├── src
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   └── README.md
├── transaction-service
│   ├── src
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   └── README.md
├── docker-compose.yml
└── README.md
```

## Setup Instructions

1. **Clone the repository**:
   ```
   git clone <repository-url>
   cd banking-system
   ```

2. **Build and run the services using Docker Compose**:
   ```
   docker-compose up --build
   ```

3. **Access the services**:
   - Account Service: `http://localhost:8080`
   - Transaction Service: `http://localhost:8081`

## Endpoints

### Account Service
- **POST /accounts**: Create a new account.
- **GET /accounts/{id}**: Retrieve account balance.

### Transaction Service
- **POST /transactions**: Transfer money between accounts.
- **GET /transactions/{account_id}**: Retrieve transaction history for an account.

## Testing

The project includes unit and integration tests for both services. To run the tests, navigate to each service directory and execute:
```
go test ./...
```

## Extras

- **gRPC**: Consider implementing gRPC for optimized communication between services.
- **Circuit Breaker**: Implement resilience patterns for better fault tolerance.
- **Cache**: Use Redis for caching frequently accessed data.
- **Kubernetes**: For scalable deployment in production environments.

## License

This project is licensed under the MIT License. See the LICENSE file for details.