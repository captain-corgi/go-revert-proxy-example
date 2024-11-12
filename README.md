# Reverse Proxy Example

Reverse Proxy Example is a simple yet powerful reverse proxy implementation built in Go. It allows you to route requests to multiple backends using round-robin and includes features like JWT authentication and rate limiting.

## Introduction

The project provides a modular and extensible framework for setting up a reverse proxy with the following features:

- **Round-Robin Load Balancing**: Distributes incoming requests across multiple backend servers.
- **JWT Authentication**: Validates JWT tokens for secure access.
- **Rate Limiting**: Limits the number of requests per minute based on a token bucket algorithm.

## Project Structure

The project is organized into several packages, each responsible for a specific aspect of the reverse proxy:

- **`cmd/main.go`**: Main entry point of the application.
- **`config/`**: Contains configuration files and logic to load configurations from `config.yaml`.
  - **`config.yaml`**: Configuration file for server settings, Redis cache parameters, JWT authentication, and rate limiting.
- **`internal/cache/`**: Implements caching functionality using Redis.
  - **`redis_cache.go`**: Logic to interact with Redis for caching data.
- **`internal/handlers/`**: Contains handlers for processing requests.
  - **`handlers.go`**: Logic for handling requests, including caching responses.
- **`internal/middleware/`**: Implements middleware components for JWT authentication and rate limiting.
  - **`jwt_middleware.go`**: Middleware for handling JWT tokens.
  - **`rate_limiter_middleware.go`**: Middleware for rate limiting the number of requests.
- **`internal/proxy/`**: Contains the logic for routing requests to backends using round-robin.
  - **`proxy.go`**: Logic for routing requests to backends and handling errors.
- **`internal/router/`**: Sets up routing logic.
  - **`router.go`**: Handles route registration and applies middleware in sequence.
- **`pkg/logger/`**: Implements a simple logger utility.
  - **`logger.go`**: Basic logging functionality with timestamps and file information.

## How to Use

### Prerequisites

- Go (1.20 or later)
- Docker (for containerized setup)

### Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/captain-corgi/go-revert-proxy-example.git
   cd go-revert-proxy-example
   ```

2. Build the project:

   ```sh
   make build
   ```

3. Run the application:

   ```sh
   ./go-revert-proxy-example
   ```

### Configuration

The configuration is managed in `config/config.yaml`. You can customize it according to your needs, including server settings, Redis cache parameters, JWT authentication, and rate limiting.

## How to Develop New Features

1. **Create a Feature Branch**:

   ```sh
   git checkout -b feature-name
   ```

2. **Develop the Feature**:
   Implement your new feature in the appropriate package. Ensure that you follow the existing code structure and conventions.

3. **Write Tests**:
   Write unit tests for your feature to ensure it works as expected. Place them in the same package.

4. **Run Tests**:
   Run the tests to make sure everything is working correctly.

   ```sh
   make test
   ```

5. **Submit a Pull Request**:
   Once your feature is complete and tested, submit a pull request for review.

## Contributing

Contributions are welcome! Please follow these guidelines:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-name`).
3. Make your changes and commit them (`git commit -am 'Add some feature'`).
4. Push to the branch (`git push origin feature-name`).
5. Open a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
