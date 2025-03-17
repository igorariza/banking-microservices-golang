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
├── transaction-rpc
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
   git clone https://github.com/igorariza/banking-microservices-golang.git
   cd banking-microservices-golang
   ```

2. **Build and run the services using Docker Compose**:
   ```
   docker-compose up --build
   ```
   Postman collection API-Rest: `Account.postman_collection.json`
   Postman collection API-GRPC: `Transaction.postman_collection.json`

3. **Access the services**:
   - Account Service: `http://localhost:8081`
   - Request Access Token
      To obtain the token, make a POST request to the authentication service:

   POST http://localhost:8081/generate_token
   ````
   {
      “name": ”userdemo”
   }
   ````

   Using the Token in Requests

      Authorization: Bearer <token>
   Content-Type: application/json

   - Transaction Service: `grpc://localhost:50055`
  Transfer Money Grpc
   {
      "from_account": "9c80ff19-f253-4d0f-bb65-c2636ca5967e",
      "to_account": "58cf0f8a-8dfe-412b-a1d0-628d53f71e18",
      "amount": 104
   }

   GetHistory Grpc
   {
      "account_id": "c249ff9f-e199-4740-8ef4-0173b7a63f72"
   }
4. **K8s**:
## Endpoints

### Account Service
- **POST /accounts**: Create a new account.
- **GET /accounts/{id}**: Retrieve account balance.

### Transaction Service
- **POST /transactions**: Transfer money between accounts.
- **GET /transactions/{account_id}**: Retrieve transaction history for an account.


## Extras

- **gRPC**: Consider implementing gRPC for optimized communication between services.
- **Kubernetes**: For scalable deployment in production environments.

## License

This project is licensed under the MIT License. See the LICENSE file for details.