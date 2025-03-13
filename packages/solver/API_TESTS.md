# Solver API Test Suite

This test suite is designed to thoroughly test the Solver API endpoints with a focus on all protocol/action combinations.

## Overview

The tests verify:

1. API accessibility via the health endpoint
2. Schema retrieval for all protocols
3. Schema retrieval for specific actions
4. Solution generation for various protocol/action combinations
5. Multi-protocol intent handling
6. Error handling for invalid inputs
7. Chain-specific protocol validation

## Running the Tests

### Prerequisites

1. A running Solver API server (usually on localhost:8000)
2. Properly configured environment variables (database connection, etc.)

### Execution

Run the tests using the provided script:

```bash
./test_api.sh
```

This script:
- Checks if the server is running
- Runs the test suite with detailed error reporting
- Provides a comprehensive summary of test results:
  - Pass/fail/skip counts
  - Detailed error messages for failures
  - List of protocols that were successfully tested

### For Development

During development, you can run the solver server with:

```bash
pnpm run dev
```

Then run the tests in another terminal.

## Test Structure

The test suite includes:

- `TestHealthEndpoint`: Verifies the server is accessible
- `TestGetSchemaEndpoint`: Tests schema retrieval for all protocols
- `TestGetSchemaForProtocol`: Tests schema retrieval for specific protocols
- `TestGetSchemaForActions`: Tests schema retrieval for specific protocol/action combinations
- `TestGetSolution`: Tests generating solutions for various protocol/action combinations
- `TestMultipleProtocolsInIntent`: Tests multi-protocol intent handling
- `TestInvalidInputs`: Tests error handling with invalid inputs
- `TestChainSpecificProtocols`: Tests protocols on their supported chains

## Adding New Tests

To add tests for new protocols or actions:

1. Add the protocol/action combination to the appropriate test case array
2. Run the tests to verify functionality
3. If the protocol is chain-specific, add it to the `chainTests` array in `TestChainSpecificProtocols`

## CI Integration

For CI integration, ensure:

1. The database container is running before tests start
2. The solver server is started with proper configuration
3. Tests are run with sufficient timeout values

Example CI workflow step:

```yaml
- name: Test Solver API
  run: |
    cd packages/solver
    pnpm run dev &
    sleep 5 # Wait for server to start
    ./test_api.sh
```

## Testing All Protocol/Action Combinations

This test suite is specifically designed to test all available protocol/action combinations in the Solver. Each combination is tested by:

1. Verifying schema availability
2. Testing solution generation
3. Validating responses

The test cases include all implemented protocols like Euler, Aave V3, Morpho, Yearn V3, ENS, Nouns, as well as utility protocols like Assert, Boolean, Math, and Database.