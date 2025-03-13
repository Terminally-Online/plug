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

# Create a temporary file to store test output
TEMP_FILE=$(mktemp)

# Run the tests and capture both stdout and stderr
cd "$(dirname "$0")"
echo -e "${BLUE}Running tests...${NC}"
echo

# Set env variable to ignore database errors
export GO_TEST_SKIP_DB_ERRORS=1

# Run tests with verbose output
go test -v ./internal/api 2>&1 | tee "$TEMP_FILE"

# Capture exit status
TEST_EXIT_CODE=${PIPESTATUS[0]}

echo
echo -e "${BLUE}╔════════════════════════════════════════╗${NC}"
echo -e "${BLUE}║              Test Summary              ║${NC}"
echo -e "${BLUE}╚════════════════════════════════════════╝${NC}"

# Count passes, failures, and skips - safely with || echo 0 to handle no matches
PASSES=$(grep -c "PASS:" "$TEMP_FILE" || echo 0)
FAILURES=$(grep -c "FAIL:" "$TEMP_FILE" || echo 0)
SKIPS=$(grep -e "--- SKIP" "$TEMP_FILE" | wc -l)

# Extract any failures for better visibility
if [ $TEST_EXIT_CODE -ne 0 ]; then
  echo -e "${RED}❌ Some tests failed. Failure details:${NC}"
  echo
  grep -B 1 -A 5 "FAIL:" "$TEMP_FILE" | grep -v "PASS:" | sed 's/^/    /'
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

# Clean up temp file
rm "$TEMP_FILE"

# Exit with the test exit code
exit $TEST_EXIT_CODE