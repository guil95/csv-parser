# Csv Parser
This project is a CSV file parser developed using Golang, following the Hexagonal Architecture (Ports & Adapters) pattern. 
The application processes CSV files, identifies invalid data, separates valid and invalid records, and 
stores them in separate files. It also provides a summary of the operation in the CLI.

## Project Structure
```shell
.
├── Makefile                       # Build automation script
├── README.md                      # Project documentation
├── cmd
│   └── main.go                    # Main application entry point
├── config
│   └── config.go                  # Configuration management
├── data                           # Directory for CSV files
│   ├── invalid_files              # Directory for invalid CSV records
│   ├── valid_files                # Directory for valid CSV records
│   ├── roster1.csv                # Sample CSV file
│   ├── roster2.csv                # Sample CSV file
│   ├── roster3.csv                # Sample CSV file
│   ├── roster4.csv                # Sample CSV file
├── go.mod                         # Go modules file
├── go.sum                         # Go dependencies file
├── internal                       # Core business logic
│   └── parser
│       ├── adapters               # Adapter implementations (CLI, File I/O)
│       │   ├── cli
│       │   │   └── cli.go         # CLI adapter for interaction
│       │   └── file
│       │       ├── reader.go      # File reader adapter
│       │       └── writer.go      # File writer adapter
│       ├── model.go               # Domain model for CSV data
│       ├── model_test.go          # Unit tests for the model
│       ├── ports                  # Port interfaces (CLI, File I/O)
│       │   ├── cli
│       │   │   └── cli.go         # CLI port interface
│       │   └── file
│       │       └── file.go        # File port interface
│       └── usecases               # Core business use cases
│           ├── parse.go           # CSV parsing use case
│           └── parse_test.go      # Unit tests for parsing use case
├── mocks                          # Mock implementations for testing
│   ├── cli.go                     # Mock for CLI interactions
│   ├── parser_uc.go               # Mock for parser use case
│   ├── reader_file.go             # Mock for file reading
│   └── writer_file.go             # Mock for file writing
└── validator                      # Validators for CSV data
    └── validator.go               # CSV data validation logic

```

## Project architecture
The project follows Hexagonal Architecture to achieve separation of concerns between the core business logic and external interfaces (such as CLI and file I/O). The architecture is broken down into:

- Core Business Logic (usecases/parse.go): Contains the main logic for parsing CSV files and handling business rules.
- Ports (ports/): Define interfaces that represent the primary (input) and secondary (output) boundaries for the core logic.
- Adapters (adapters/): Implementations of the port interfaces for specific technologies or external services like CLI and file handling.
- Mocks (mocks/): Mock implementations of ports for unit testing purposes.

## How to run 
```shell
make run FILE=path/to/file
```

The results will be created in `data/invalid_files` and `data/valid_files`

