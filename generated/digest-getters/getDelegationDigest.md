# getDelegationDigest
        
Encode [Delegation](/generated/base-types/Delegation) data into a digest hash that has been localized to the domain of the contract.

## Parameters

- `$input` : [Delegation](/generated/base-types/Delegation) : The `Delegation` data to encode.

## Returns

- `$digest` : `bytes32` : The digest hash of the encoded [Delegation](/generated/base-types/Delegation) data.

## Onchain Implementation

::: code-group

``` solidity [Types.sol:getDelegationDigest]
function getDelegationDigest(
	Delegation memory $input
) public view virtual returns (bytes32 $digest) {
	$digest = keccak256(
		abi.encodePacked(
			"\x19\x01",
			domainHash,
			getDelegationHash($input)
		)
	);
}
```

:::