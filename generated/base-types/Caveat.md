# Caveat
        
Type hash representing the [Caveat](/base-types/Caveat) data type providing EIP-712 compatability for encoding and decoding.

## EIP-712 Type Definition

::: code-group

```typescript [Caveat]
{
    { name: 'enforcer', type: 'address' },
	{ name: 'terms', type: 'bytes' } 
}
```

:::

## Onchain Implementation

::: code-group

```solidity [Types.sol:CAVEAT_TYPEHASH]
bytes32 constant CAVEAT_TYPEHASH = keccak256('Caveat(address enforcer,bytes terms)');
```

:::