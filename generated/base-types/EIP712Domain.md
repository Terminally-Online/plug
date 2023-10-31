# EIP712Domain
        
Type hash representing the [EIP712Domain](/base-types/EIP712Domain) data type providing EIP-712 compatability for encoding and decoding.

## EIP-712 Type Definition

::: code-group

```typescript [EIP712Domain]
{
    { name: 'name', type: 'string' },
	{ name: 'version', type: 'string' },
	{ name: 'chainId', type: 'uint256' },
	{ name: 'verifyingContract', type: 'address' } 
}
```

:::

## Onchain Implementation

::: code-group

```solidity [Types.sol:EIP712_DOMAIN_TYPEHASH]
bytes32 constant EIP712_DOMAIN_TYPEHASH = keccak256('EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)');
```

:::