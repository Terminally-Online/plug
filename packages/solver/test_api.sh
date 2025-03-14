#!/bin/bash

# Define colors for better output
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[0;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}╔════════════════════════════════════════╗${NC}"
echo -e "${BLUE}║         Solver API Test Suite          ║${NC}"
echo -e "${BLUE}╚════════════════════════════════════════╝${NC}"

# Parse command line arguments
SKIP_DB_CHECK=false
HELP=false
TEST_MODE="all"  # all, basic, base, mainnet, etc.
CHAIN_ID=""      # Optional specific chain ID to test

for arg in "$@"; do
  case $arg in
    --skip-db-check)
      SKIP_DB_CHECK=true
      shift
      ;;
    --test-mode=*)
      TEST_MODE="${arg#*=}"
      shift
      ;;
    --chain-id=*)
      CHAIN_ID="${arg#*=}"
      shift
      ;;
    --help)
      HELP=true
      shift
      ;;
  esac
done

if [ "$HELP" = true ]; then
  echo -e "Usage: ./test_api.sh [options]"
  echo -e ""
  echo -e "Options:"
  echo -e "  --skip-db-check           Skip database connection check"
  echo -e "  --test-mode=MODE          Set test mode (all, base, mainnet, basic, minimal, noactions)"
  echo -e "                            - all: Run all tests"
  echo -e "                            - base: Run only tests for Base chain (chain ID 8453)"
  echo -e "                            - mainnet: Run only tests for Ethereum mainnet (chain ID 1)"
  echo -e "                            - basic: Run most critical tests"
  echo -e "                            - minimal: Run only essential tests"
  echo -e "                            - noactions: Skip protocol action tests that might fail"
  echo -e "  --chain-id=ID             Run tests only for a specific chain ID (e.g. 8453 for Base)"
  echo -e "                            This overrides the chain selection in --test-mode"
  echo -e "  --help                    Show this help message"
  exit 0
fi

# Check if the solver server is running
if ! curl -s "http://localhost:8080/health" > /dev/null 2>&1; then
  echo -e "${YELLOW}⚠️  Warning: The solver server is not running on localhost:8080${NC}"
  echo -e "Tests requiring the server will be ${YELLOW}skipped${NC} automatically."
  echo -e "To start the solver server, run: ${GREEN}pnpm run dev${NC}"
  echo
else
  echo -e "${GREEN}✓ Solver server is running on localhost:8080${NC}"
  echo
fi

# Check database connection if not skipped
if [ "$SKIP_DB_CHECK" = false ]; then
  # Extract database connection info from environment or use defaults
  DB_HOST=${DATABASE_HOST:-localhost}
  DB_PORT=${DATABASE_PORT:-6432}
  DB_USER=${DATABASE_USER:-plug}
  DB_PASS=${DATABASE_PASSWORD:-plugdev}
  DB_NAME=${DATABASE_NAME:-plug_solver}
  
  echo -e "Checking database connection to ${DB_HOST}:${DB_PORT}..."
  if nc -z -w1 ${DB_HOST} ${DB_PORT} > /dev/null 2>&1; then
    echo -e "${GREEN}✓ Database is accessible at ${DB_HOST}:${DB_PORT}${NC}"
  else
    echo -e "${YELLOW}⚠️  Warning: Database is not accessible at ${DB_HOST}:${DB_PORT}${NC}"
    echo -e "Tests requiring database will use test defaults."
  fi
  echo
fi

# Create a temporary file to store test output
TEMP_FILE=$(mktemp)

# Run the tests and capture both stdout and stderr
cd "$(dirname "$0")"
echo -e "${BLUE}Running tests...${NC}"
echo

# Set env variables for testing
export GO_TEST_MODE=true
export ALLOW_TEST_DB_FALLBACK=true

# Set API keys for testing
export ADMIN_API_KEY=${ADMIN_API_KEY:-bingbopboombam}
export TEST_API_KEY=${TEST_API_KEY:-testing}  # Special test key that doesn't get rate limited

# Set chain ID if provided
if [ -n "$CHAIN_ID" ]; then
  export TEST_CHAIN_ID="$CHAIN_ID"
  echo -e "${BLUE}Focusing tests on chain ID: ${CHAIN_ID}${NC}"
fi

# Show which API key is being used
echo -e "${GREEN}Using API key for tests: ${TEST_API_KEY}${NC}"

# Set test filter based on the test mode
TEST_FILTER=""
case "$TEST_MODE" in
  minimal)
    echo -e "${YELLOW}Running in minimal test mode - only essential tests${NC}"
    TEST_FILTER="-run=TestHealthEndpoint|TestGetSchema"
    ;;
  basic)
    echo -e "${YELLOW}Running in basic test mode - core functionality tests${NC}"
    TEST_FILTER="-run=TestHealthEndpoint|TestGetSchema|TestInvalidInputs"
    ;;
  base)
    echo -e "${BLUE}Running tests focusing on Base chain (8453) protocols${NC}"
    # Target tests specifically marked with chain_8453 in the test name
    TEST_FILTER="-run=.*chain_8453"
    ;;
  noactions)
    echo -e "${YELLOW}Running tests excluding protocol action tests that might fail${NC}"
    TEST_FILTER="-run=TestHealthEndpoint|TestGetSchema"
    ;;
  mainnet)
    echo -e "${BLUE}Running tests focusing on Ethereum mainnet (chain ID 1)${NC}"
    # Target tests specifically marked with chain_1 in the test name
    TEST_FILTER="-run=.*chain_1"
    ;;
  all|*)
    echo -e "${GREEN}Running all tests${NC}"
    TEST_FILTER=""
    ;;
esac

# Run tests with verbose output
go test -v ./internal/test $TEST_FILTER 2>&1 | tee "$TEMP_FILE"

# Capture exit status
TEST_EXIT_CODE=${PIPESTATUS[0]}

echo
echo -e "${BLUE}╔════════════════════════════════════════╗${NC}"
echo -e "${BLUE}║              Test Summary              ║${NC}"
echo -e "${BLUE}╚════════════════════════════════════════╝${NC}"

# Count passes, failures, and skips - safely with || echo 0 to handle no matches
PASSES=$(grep -c "PASS:" "$TEMP_FILE" || echo 0)
FAILURES=$(grep -c "FAIL:" "$TEMP_FILE" || echo 0)
SKIPS=$(grep -e "--- SKIP" "$TEMP_FILE" | wc -l || echo 0)

# Check for panics in the output
PANICS=$(grep -c "panic: " "$TEMP_FILE" || echo 0)
if [ "$(echo $PANICS | tr -d ' ')" -gt 0 ]; then
  echo -e "${RED}⚠️  Test execution had $PANICS panic(s). Details:${NC}"
  echo
  grep -B 2 -A 5 "panic: " "$TEMP_FILE" | sed 's/^/    /'
  echo
fi

# Check for rate limit issues
RATE_LIMITS=$(grep -c "rate limit exceeded" "$TEMP_FILE" || echo 0)
if [ "$(echo $RATE_LIMITS | tr -d ' ')" -gt 0 ]; then
  echo -e "${YELLOW}⚠️  Warning: $RATE_LIMITS rate limit exceeded errors detected.${NC}"
  echo -e "The test script may be hitting the API too frequently. Consider:"
  echo -e "  1. Increasing the rate limit in the database for the test API key"
  echo -e "  2. Adding delays between API calls"
  echo -e "  3. Running fewer tests in parallel"
  echo
fi

# Extract any failures for better visibility
if [ "$(echo $TEST_EXIT_CODE | tr -d ' ')" -ne 0 ]; then
  echo -e "${RED}❌ Some tests failed. Failure details:${NC}"
  echo
  grep -B 1 -A 5 "FAIL:" "$TEMP_FILE" | grep -v "PASS:" | sed 's/^/    /' || echo "    No detailed failure information available"
  echo
else
  echo -e "${GREEN}✓ All tests passed!${NC}"
  echo
fi

# Print summary counts
echo -e "📊 ${GREEN}Passed: $PASSES${NC} | ${RED}Failed: $FAILURES${NC} | ${YELLOW}Skipped: $SKIPS${NC}"

# Print protocols tested - safely handle no matches
echo
echo -e "${BLUE}Protocols tested:${NC}"
PROTOCOLS=$(grep -o "Protocol_[a-z0-9_]*" "$TEMP_FILE" 2>/dev/null | sort | uniq | sed 's/Protocol_/    ✓ /' || echo "    None detected")
echo "$PROTOCOLS"

# Print actions tested
echo
echo -e "${BLUE}Actions tested:${NC}"
ACTIONS=$(grep -o "Testing protocol=[a-z0-9_]* action=[a-z0-9_]*" "$TEMP_FILE" 2>/dev/null | sort | uniq | sed 's/Testing protocol=\([a-z0-9_]*\) action=\([a-z0-9_]*\)/    ✓ \1: \2/' || echo "    None detected")
echo "$ACTIONS"

# Chain-specific statistics
echo
echo -e "${BLUE}Chain-specific test statistics:${NC}"
BASE_TESTS=$(grep -c "chain_8453" "$TEMP_FILE" || echo 0)
BASE_PASSED=$(grep -o "Successfully generated solution.*chain_8453" "$TEMP_FILE" | wc -l || echo 0)
MAINNET_TESTS=$(grep -c "chain_1" "$TEMP_FILE" || echo 0)
MAINNET_PASSED=$(grep -o "Successfully generated solution.*chain_1" "$TEMP_FILE" | wc -l || echo 0)

echo -e "    Base (Chain 8453): Ran ${BASE_TESTS} tests, ${GREEN}${BASE_PASSED} passed${NC}"
echo -e "    Ethereum (Chain 1): Ran ${MAINNET_TESTS} tests, ${GREEN}${MAINNET_PASSED} passed${NC}"

# Clean up temp file
rm "$TEMP_FILE"

# Exit with the test exit code
exit $TEST_EXIT_CODE