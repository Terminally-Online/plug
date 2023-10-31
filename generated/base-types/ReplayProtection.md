# ReplayProtection
        
Type hash representing the [ReplayProtection](/base-types/ReplayProtection) data type providing EIP-712 compatability for encoding and decoding.

## EIP-712 Type Definition

::: code-group

```typescript [ReplayProtection]
{
    { name: 'nonce', type: 'uint256' },
	{ name: 'queue', type: 'uint256' } 
}
```

:::

## Onchain Implementation

::: code-group

```solidity [Types.sol:REPLAY_PROTECTION_TYPEHASH]
bytes32 constant REPLAY_PROTECTION_TYPEHASH = keccak256('ReplayProtection(uint256 nonce,uint256 queue)');
```

:::