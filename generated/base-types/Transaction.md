# Transaction
        
Type hash representing the [Transaction](/base-types/Transaction) data type providing EIP-712 compatability for encoding and decoding.

## EIP-712 Type Definition

::: code-group

```typescript [Transaction]
{
    { name: 'to', type: 'address' },
	{ name: 'gasLimit', type: 'uint256' },
	{ name: 'data', type: 'bytes' } 
}
```

:::

## Onchain Implementation

::: code-group

```solidity [Types.sol:TRANSACTION_TYPEHASH]
bytes32 constant TRANSACTION_TYPEHASH = keccak256('Transaction(address to,uint256 gasLimit,bytes data)');
```

:::