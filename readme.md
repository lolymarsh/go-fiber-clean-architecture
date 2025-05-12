# Go Fiber Clean Architecture

A clean architecture implementation using Go Fiber framework for building scalable and maintainable web applications.

Inspired by https://github.com/BoomTHDev/b-pos-golang-server

## Project Structure

```
├── config/         # Configuration files and setup
├── database/       # Database configurations and migrations
├── entity/         # Domain entities/models
├── pkg/            # Shared packages and utilities
├── scripts/        # Helper scripts
├── server/         # Server setup and HTTP handlers
├── .env            # Environment variables
├── .env_example    # Example environment variables template
├── main.go         # Application entry point
└── go.mod          # Go module dependencies
```

## Features

- Clean Architecture pattern
- Fiber web framework
- Environment configuration
- Graceful shutdown
- Modular structure
- Database integration

## Prerequisites

- Go 1.x or higher
- Environment variables configured (copy from .env_example)

## Getting Started

1. Clone the repository
2. Copy `.env_example` to `.env` and configure your environment variables
3. Install dependencies:
   ```bash
   go mod download
   ```
4. Run the application:
   ```bash
   go run main.go
   ```

## Architecture

This project follows Clean Architecture principles with the following layers:

- Entity Layer: Core business logic and domain models
- Use Case Layer: Application business rules
- Interface Layer: Controllers and external interfaces
- Framework Layer: Web framework and external tools

## Configuration

The application uses environment variables for configuration. Check `.env_example` for required variables.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License.