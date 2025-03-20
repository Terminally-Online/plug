# @plug/rope

A highly opinionated SDK for the Plug API that enables node-based actions with OpenAPI schema integration.

## Features

- Node-based action system using Knots and Ropes
- Automatic type generation from OpenAPI schema
- Type-safe API client
- Visual editor backend support

## Concept

The Rope SDK uses a node-based approach to blockchain operations:

- **Rope**: A sequence of connected protocol actions
- **Knot**: An individual protocol action node (e.g., Uniswap swap, AAVE deposit)
- **Schema**: The available actions and their structure loaded from the API

This design enables building complex transaction sequences that can be executed as a single intent.

## Installation

Using pnpm:

```bash
pnpm add @plug/rope
```

## Usage

```typescript
import { Rope, PlugTypes } from '@plug/rope';

// Create and initialize a rope with your API key
const rope = new Rope('your-api-key');
await rope.initialize();

// Get available schemas (protocols and their actions)
const schemas = rope.getAvailableSchemas();

// Create knots for specific protocol actions
const swapKnot = rope.createKnot('uniswap', 'swap', {
  tokenIn: 'ETH',
  tokenOut: 'DAI',
  amountIn: '1.0'
});

const depositKnot = rope.createKnot('aave', 'deposit', {
  token: 'DAI',
  amount: '100'
});

// Add knots to the rope (chain actions together)
rope.addKnot(swapKnot)
    .addKnot(depositKnot);

// Get all knots in the rope
const knots = rope.getKnots();

// Generate an intent from the rope's knots
const intent = rope.buildIntent();

// Execute the rope as a transaction (submit the intent)
const walletAddress = '0x123...';
const result = await rope.execute(walletAddress);

// Direct API access
// Legacy API methods
const data = await rope.get('protocol', 'action');
const response = await rope.post({ key: 'value' });

// Generated API methods with full OpenAPI typing
const specificSchema = await rope.api().getSolver({
  chainId: 1,
  protocol: 'uniswap'
});
```

## API Generation

This package includes tools to generate TypeScript interfaces and API clients from an OpenAPI schema.

### Regenerating the API

To regenerate the API client and types from the OpenAPI schema:

```bash
pnpm generate-api
```

This will parse the OpenAPI schema in `/src/lib/openapi.json` and generate:

1. Type definitions in `/src/types/generated.ts`
2. API client in `/src/models/generated-api.ts`

## Development

### Setup

```bash
# Install dependencies
pnpm install
```

### Build

```bash
# Build the package
pnpm build

# Watch mode
pnpm dev
```

### Testing

```bash
# Run all tests
pnpm test

# Watch mode
pnpm test:watch
```

### Linting

```bash
# Run linter
pnpm lint

# Fix linting issues
pnpm lint:fix
```

## License

MIT