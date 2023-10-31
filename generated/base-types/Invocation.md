# Invocation
        
Type hash representing the [Invocation](/base-types/Invocation) data type providing EIP-712 compatability for encoding and decoding.

## EIP-712 Type Definition

::: code-group

```typescript [Invocation]
{
    { name: 'transaction', type: 'Transaction' },
	{ name: 'authority', type: 'SignedDelegation[]' } 
}
```

:::

## Onchain Implementation

::: code-group

```solidity [Types.sol:INVOCATION_TYPEHASH]
bytes32 constant INVOCATION_TYPEHASH = keccak256('Invocation(Transaction transaction,SignedDelegation[] authority)Caveat(address enforcer,bytes terms)Delegation(address delegate,bytes32 authority,Caveat[] caveats,bytes32 salt)SignedDelegation(Delegation delegation,bytes signature)Transaction(address to,uint256 gasLimit,bytes data)');
```

:::