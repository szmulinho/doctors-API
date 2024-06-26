# Doctors API

This API allows you to manage doctor accounts, providing endpoints for creating, logging and retrieving doctors records. Additionally, it includes authentication functionality.

## Endpoints

### Register New Doctor

Endpoint: /register

Method: POST

Description: Adds new doctor account to the system.

### Login

Endpoint: /login
Method: POST
Description: Logs doctor to the system.

### Retrieve a Single Doctor Data

Endpoint: /doctor

Method: GET

Description: Retrieves details of a specific doctor by his token.

## Retrieve All Doctors

Endpoint: /doctors

Method: GET

Description: Retrieves all doctors in the system.

### Generate

Endpoint: /generate

Method: POST

Description: Generates a token for the doctor.

## Testing

Unit and integration tests are implemented for the API. Tests can be run individually or using Docker Compose.

### Run Tests Individually

To run the tests manually, use the following command:

```go test ./...```

### Run Tests with Docker Compose

A docker-compose.yml file is provided to facilitate testing in a containerized environment. To run the tests using Docker Compose, use the following command:

```docker-compose up``` or ```make doker-tests```

This command will build the Docker images, run the tests, and then stop the containers.

## Setup Instructions

### 1. Clone the repository:

```git clone https://github.com/szmulinho/doctors-API.git```
```cd doctors-api```

### 2. Install dependencies:

```go mod tidy```

### 3. Run the server:

```go run main.go```

## Contribution Guidelines

Fork the repository

Create a new branch ```git checkout -b feature/your-feature-name```

Commit your changes ```git commit -m 'Add some feature'```

Push to the branch ```git push origin feature/your-feature-name```

Create a new Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

For any questions or suggestions, feel free to open an issue or contact the repository owner.



