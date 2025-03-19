# Solver API

<span style="color: rgba(0,0,0,0.6)">The Solver API is the engine behind Plug's transaction execution, converting user intents into optimized blockchain transactions. This guide explains how to integrate and use the API in your applications.</span>

## What is the Solver API?

The Solver API is a powerful service that converts declarative user intents into executable blockchain transactions. Rather than requiring users to specify exact transaction parameters, the API allows them to describe what they want to accomplish, and it handles the complexity of creating optimal transactions.

Key features:
- **Intent-Based Transactions**: Define what you want to achieve, not how to achieve it
- **Protocol Integrations**: Support for multiple DeFi protocols (Aave, Yearn, Morpho, etc.)
- **Transaction Simulation**: Preview transaction outcomes before execution
- **Cross-Chain Support**: Execute transactions across multiple blockchains
- **Smart Contract Wallet Support**: Compatible with both EOA and contract wallets

## Core Concepts

### Intents

Intents are structured requests that describe what a user wants to accomplish. An intent consists of:

- **Protocol and Action**: Which protocol to interact with and what action to perform
- **Parameters**: Protocol-specific parameters for the action
- **Chain ID**: The blockchain network to operate on
- **User Address**: The wallet address initiating the transaction
- **Options**: Configuration options for execution

Example intent structure:

```json
{
  "chainId": 8453,
  "from": "0xYourWalletAddress",
  "inputs": [
    {
      "protocol": "aave_v3",
      "action": "deposit",
      "token": "0xTokenAddress:TokenDecimals",
      "amount": "1.1",
    },
    {
      "protocol": "aave_v3",
      "action": "borrow",
      "token": "0xTokenAddress:TokenDecimals",
      "amount": "200",
    }
  ],
  "options": {
    "isEOA": false,
    "simulate": true,
    "submit": false
  }
}
```

### Plugs

When the Solver processes an intent, it generates one or more "plugs" - the actual transactions that will be executed on the blockchain. A plug contains:

- **Target**: The contract address to call
- **Call Data**: The encoded function call
- **Value**: Any ETH value to send with the transaction
- **Updates**: Any dynamic updates to apply at execution time

## Integration Guide

### Authentication

All API requests require an API key passed via the `X-Api-Key` header. Contact the Plug team to obtain an API key for your application.

```javascript
const headers = {
  'X-Api-Key': 'your-api-key',
  'Content-Type': 'application/json'
};
```

### Basic Workflow

1. **Discover Available Actions**

   First, query the API to get the available actions for a specific protocol and chain:

   ```javascript
   // Get available actions for Aave V3 on Ethereum
   const response = await fetch('https://api.plug.io/solver?chainId=1&protocol=aave_v3', {
     method: 'GET',
     headers
   });
   const schema = await response.json();
   ```

2. **Construct an Intent**

   Use the schema information to build a valid intent:

   ```javascript
   const intent = {
     chainId: 1,
     from: userWalletAddress,
     inputs: [
       {
         protocol: 'aave_v3',
         action: 'deposit',
         parameters: {
           asset: '0xTokenAddress',
           amount: ethers.parseUnits('1.0', 18).toString(),
           onBehalfOf: userWalletAddress
         }
       }
     ],
     options: {
       isEOA: true,
       simulate: true
     }
   };
   ```

3. **Submit the Intent**

   Send the intent to the Solver API:

   ```javascript
   const response = await fetch('https://api.plug.io/solver', {
     method: 'POST',
     headers,
     body: JSON.stringify(intent)
   });
   const result = await response.json();
   ```

4. **Execute the Transaction**

   The response will contain the transaction data needed for execution:

   ```javascript
   // For EOA wallets
   const tx = {
     to: result.tx.to,
     data: result.tx.data,
     value: result.tx.value,
     gasLimit: result.simulation.gasUsed
   };
   
   // Send the transaction using your preferred web3 library
   const txResponse = await wallet.sendTransaction(tx);
   ```

### Socket Transactions (Contract Wallets)

For contract wallets using Socket, the process is slightly different:

```javascript
const intent = {
  chainId: 1,
  from: userWalletAddress,
  inputs: [/* ... */],
  options: {
    isEOA: false,
    simulate: true
  }
};

// Submit the intent
const response = await fetch('https://api.plug.io/solver', {
  method: 'POST',
  headers,
  body: JSON.stringify(intent)
});
const result = await response.json();

// The result contains an EIP-712 signature that can be used with the Socket contract
const signature = result.signature;
```

## Configuration Options

The API supports several options to customize intent execution:

- **isEOA** (boolean): Whether the transaction is for an EOA wallet or a contract wallet
- **simulate** (boolean): Whether to simulate the transaction before returning results
- **submit** (boolean): Whether the Solver should submit the transaction on behalf of the user (requires additional permissions)
- **maxFeePerGas** (string): Maximum gas price to use
- **maxPriorityFeePerGas** (string): Maximum priority fee to use

## Error Handling

The API uses standard HTTP status codes:

- **400**: Bad Request - Invalid intent format or parameters
- **401**: Unauthorized - Invalid API key
- **429**: Too Many Requests - Rate limit exceeded
- **500**: Internal Server Error - Failed to process intent

Error responses include a descriptive message:

```json
{
  "error": "Invalid asset address",
  "code": 400
}
```

## Rate Limiting

The API implements rate limiting based on your API key. The limits are returned in the response headers:

- **X-RateLimit-Limit**: Maximum requests allowed in the current window
- **X-RateLimit-Used**: Current count of requests in the window
- **X-RateLimit-Reset**: Unix timestamp when the rate limit window resets

## Best Practices

1. **Always simulate first**: Set `simulate: true` to preview transaction outcomes before execution
2. **Handle errors gracefully**: Provide user-friendly error messages for API errors
3. **Respect rate limits**: Implement exponential backoff for retries when rate limited
4. **Validate user inputs**: Check that user inputs match the expected parameters from the schema

## Supported Protocols

The Solver API supports various protocols, including:

- **Aave V3**: Lending and borrowing
- **Yearn V3**: Yield optimization
- **Morpho**: Decentralized lending
- **Euler**: Lending protocol
- **Nouns**: NFT protocol
- **BasePaint**: NFT platform

Each protocol supports specific actions like deposit, withdraw, borrow, and repay, with their own parameters and constraints.

By integrating the Solver API, your application can offer users a more intuitive way to interact with DeFi protocols, focusing on their goals rather than transaction details.