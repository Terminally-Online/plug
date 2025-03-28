# Circuit AVS

Circuit is the Attestation Virtual Service for Plug, providing decentralized execution and validation of intents through Othentic's attestation network.

## Architecture

- Stream-based intent distribution via Redis Streams
- Auction-based operator selection for intent execution
- Competitive bidding system with reputation and gas price factors
- On-chain attestation through Othentic
- Transaction validation via on-chain events

## Message Flow

1. Intent published to `circuit:intents` Redis stream with unsigned plugs
2. Operators subscribe to streams with consumer groups
3. Operators simulate plugs and determine if they're interested
4. Interested operators submit bids via the API
5. Winner receives signed LivePlugs to execute
6. Execution result is relayed to blockchain
7. Transaction receipts validated on-chain

## Configuration

```
AVS_ENV=development|production  # Environment (default: development)
PORT=8081                       # HTTP server port (default: 8081)
PRIVATE_KEY=                    # Operator's private key
RPC_8453=                       # Base chain RPC URL
```

## Commands

```
go run cmd/avs/main.go help     # Show help
go run cmd/avs/main.go version  # Show version
go run cmd/avs/main.go info     # Show operator info
go run cmd/avs/main.go          # Start service
```

## API

- `GET /status`: Service status
- `GET /metrics`: Pending tasks metrics
- `POST /task/validate`: Validate transactions
- `POST /intent/bid`: Submit bid to express interest in an intent
- `GET /intent/:id/auction`: Get current auction status for an intent
- `GET /intent/:id/signed`: Get signed LivePlugs for a winning bidder

## Stream Format

Messages in `circuit:intents` stream:
- `intent_id`: Intent identifier
- `data`: JSON string with intent data
- `plugs`: Unsigned plugs for simulation
- `from`: Sender address
- `chainIdUint`: Chain ID as uint64
- `categories`: Optional filter categories
- `chain_id`: Optional chain ID for filtering
- `timestamp`: Publication time

## Bidding System

The bidding system allows operators to express interest in executing intents:

```json
{
  "intentId": "0x123...",
  "operator": "0xabc...",
  "timestamp": 1679528400,
  "gasPrice": "1.5 gwei"
}
```

Operators can specify custom parameters like gas price to compete for execution.

