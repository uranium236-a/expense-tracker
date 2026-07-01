# Expense Tracker CLI

A simple command-line Expense Tracker built with Go. It allows users to add, view, delete, and analyze expenses while automatically saving data to a local JSON file.

## Features

- Add new expenses
- View all stored expenses
- Delete expenses by ID
- Display the maximum expense
- Display the minimum expense
- Persistent storage using JSON
- Interactive CLI mode
- Direct command execution through command-line arguments

## Project Structure

```
expense-tracker/
├── cli/
│   ├── home.go
│   └── reader.go
├── models/
│   └── expense.go
├── storage/
│   └── data.json
├── main.go
├── go.mod
└── README.md
```

## Requirements

- Go 1.25 or later

## Installation

Clone the repository:

```bash
git clone <repository-url>
cd expense-tracker
```

Run the application:

```bash
go run .
```

Or build the executable:

```bash
go build
./expense-tracker-cli
```

## Usage

### Interactive Mode

Start the application:

```bash
go run .
```

Example session:

```text
>> add Food 250
Expense added successfully

>> add Travel 1200
Expense added successfully

>> view
0 Food 250
1 Travel 1200

>> max
Maximum Expense: 1200
Subject: Travel

>> min
Minimum Expense: 250
Subject: Food

>> del 0
Expense deleted successfully

>> exit
```

### Command-Line Mode

Execute commands directly:

```bash
go run . add Food 500
```

```bash
go run . view
```

```bash
go run . del 1
```

```bash
go run . max
```

```bash
go run . min
```

## Available Commands

| Command | Description |
|---------|-------------|
| `add <subject> <amount>` | Add a new expense |
| `view` | Display all expenses |
| `del <id>` | Delete an expense |
| `max` | Show the highest expense |
| `min` | Show the lowest expense |
| `help` | Display available commands |
| `exit` | Exit interactive mode |

## Data Storage

All expense data is stored locally in:

```
storage/data.json
```

The application automatically:

- Loads existing data on startup
- Saves changes on exit

## Example

```text
>> add Grocery 800
Expense added successfully

>> add Fuel 1500
Expense added successfully

>> view
0 Grocery 800
1 Fuel 1500

>> max
Maximum Expense: 1500
Subject: Fuel
```

## Future Improvements

- Edit existing expenses
- Expense categories
- Monthly reports
- Date and time tracking
- Total expense calculation
- Search and filtering
- CSV export
- Budget tracking

## License

This project is licensed under the MIT License.
