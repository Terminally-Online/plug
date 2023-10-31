# SignedDelegation
        
Type hash representing the [SignedDelegation](/base-types/SignedDelegation) data type providing EIP-712 compatability for encoding and decoding.

## EIP-712 Type Definition

::: code-group

```typescript [SignedDelegation]
{
    { name: 'delegation', type: 'Delegation' },
	{ name: 'signature', type: 'bytes' } 
}
```

:::

## Onchain Implementation

::: code-group

```solidity [Types.sol:SIGNED_DELEGATION_TYPEHASH]
bytes32 constant SIGNED_DELEGATION_TYPEHASH = keccak256('SignedDelegation(Delegation delegation,bytes signature)Caveat(address enforcer,bytes terms)Delegation(address delegate,bytes32 authority,Caveat[] caveats,bytes32 salt)');
```

:::