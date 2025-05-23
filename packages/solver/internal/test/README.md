# Solver API Tests

This directory contains tests for the Solver API. The tests are structured as follows:

## Directory Structure

- `solver_test.go`: The main test file that sets up the test environment and runs the tests
- `cases/`: Directory containing test cases for each protocol
  - `euler_test_cases.json`: Test cases for the Euler protocol
  - `aave_v3_test_cases.json`: Test cases for the Aave V3 protocol
  - etc.

## Adding New Test Cases

To add new test cases for an existing protocol:
1. Edit the corresponding JSON file in the `cases/` directory
2. Add new test cases following the existing format

To add test cases for a new protocol:
1. Create a new file in the `cases/` directory named `{protocol}_test_cases.json`
2. Add test cases in the following format:

```json
[
  {
    "name": "Test Case Name",
    "intent": {
      "chainId": 1,
      "inputs": [
        {
          "protocol": "protocol_name",
          "action": "action_name",
          // Other action-specific parameters
        }
      ],
      "options": {
        "isEOA": false,
        "simulate": false,
        "submit": false
      }
    },
    "expectOk": true
  }
]
```

## Running Tests

Tests can be run using ```go test -v ./internal/test/solver_test.go``` from the root of the repository.
