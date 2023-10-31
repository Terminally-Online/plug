# Delegation
        
Type hash representing the [Delegation](/base-types/Delegation) data type providing EIP-712 compatability for encoding and decoding.

## EIP-712 Type Definition

::: code-group

```typescript [Delegation]
{
    { name: 'delegate', type: 'address' },
	{ name: 'authority', type: 'bytes32' },
	{ name: 'caveats', type: 'Caveat[]' },
	{ name: 'salt', type: 'bytes32' } 
}
```

:::

## Onchain Implementation

::: code-group

```solidity [Types.sol:DELEGATION_TYPEHASH]
bytes32 constant DELEGATION_TYPEHASH = keccak256('Delegation(address delegate,bytes32 authority,Caveat[] caveats,bytes32 salt)Caveat(address enforcer,bytes terms)');
```

:::