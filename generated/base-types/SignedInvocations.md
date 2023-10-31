# SignedInvocations
        
Type hash representing the [SignedInvocations](/base-types/SignedInvocations) data type providing EIP-712 compatability for encoding and decoding.

## EIP-712 Type Definition

::: code-group

```typescript [SignedInvocations]
{
    { name: 'invocations', type: 'Invocations' },
	{ name: 'signature', type: 'bytes' } 
}
```

:::

## Onchain Implementation

::: code-group

```solidity [Types.sol:SIGNED_INVOCATIONS_TYPEHASH]
bytes32 constant SIGNED_INVOCATIONS_TYPEHASH = keccak256('SignedInvocations(Invocations invocations,bytes signature)Caveat(address enforcer,bytes terms)Delegation(address delegate,bytes32 authority,Caveat[] caveats,bytes32 salt)Invocation(Transaction transaction,SignedDelegation[] authority)Invocations(Invocation[] batch,ReplayProtection replayProtection)ReplayProtection(uint256 nonce,uint256 queue)SignedDelegation(Delegation delegation,bytes signature)Transaction(address to,uint256 gasLimit,bytes data)');
```

:::