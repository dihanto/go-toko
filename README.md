# GO-TOKO

![go-toko-high-resolution-logo-color-on-transparent-background (1)](https://github.com/dihanto/go-toko/assets/39905651/5b291682-b762-4196-bbd0-3d972dd062ce)

[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/dihanto/go-toko)](https://goreportcard.com/report/github.com/dihanto/go-toko)

## Description

go-toko is a simple e-commerce platform written in Go. It provides essential features for managing products, customers, and orders for online retailers.

## Project Structure

- `config/`: Contains configuration files for the application.
- `controller/`: Handles incoming HTTP requests and manages the application flow.
- `database/`: Responsible for database-related operations and connections.
  - `migration/`: Contains database migration files.
- `docs/`: Stores documentation files related to the project.
- `exception/`: Manages application-specific error handling and custom error types.
- `helper/`: Holds utility functions and helper code.
- `middleware/`: Contains custom middleware used in the application.
- `model/`: Defines the application's data models and entities.
  - `entity/`: Contains various data entities used in the application.
  - `web/`: Defines the data models for web-related functionalities.
    - `request/`: Contains request data models.
    - `response/`: Contains response data models.
- `repository/`: Implements the data access layer and interacts with the database.
- `usecase/`: Contains business logic and use cases for the application.
- `go.mod` and `go.sum`: Go module files that define the project's dependencies.
- `main.go`: The entry point of the application.

## Contributing

Contributions are welcome! If you find any issues or want to add new features, please create a pull request. Before making significant changes, it's best to open an issue first to discuss the proposed changes.

## License

This project is licensed under the [MIT License](LICENSE). Feel free to use, modify, and distribute the code as per the terms of the license.


