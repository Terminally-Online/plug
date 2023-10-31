# Invocations
        
Type hash representing the [Invocations](/base-types/Invocations) data type providing EIP-712 compatability for encoding and decoding.

## EIP-712 Type Definition

::: code-group

```typescript [Invocations]
{
    { name: 'batch', type: 'Invocation[]' },
	{ name: 'replayProtection', type: 'ReplayProtection' } 
}
```

:::

## Onchain Implementation

::: code-group

```solidity [Types.sol:INVOCATIONS_TYPEHASH]
bytes32 constant INVOCATIONS_TYPEHASH = keccak256('Invocations(Invocation[] batch,ReplayProtection replayProtection)Caveat(address enforcer,bytes terms)Delegation(address delegate,bytes32 authority,Caveat[] caveats,bytes32 salt)Invocation(Transaction transaction,SignedDelegation[] authority)ReplayProtection(uint256 nonce,uint256 queue)SignedDelegation(Delegation delegation,bytes signature)Transaction(address to,uint256 gasLimit,bytes data)');
```

:::