# getPinDigest
        
Encode [Pin](/generated/base-types/Pin) data into a digest hash that has been localized to the domain of the contract.

## Parameters

- `$input` : [Pin](/generated/base-types/Pin) : The `Pin` data to encode.

## Returns

- `$digest` : `bytes32` : The digest hash of the encoded [Pin](/generated/base-types/Pin) data.

## Onchain Implementation

::: code-group

``` solidity [Types.sol:getPinDigest]
function getPinDigest(
	Pin memory $input
) public view virtual returns (bytes32 $digest) {
	$digest = keccak256(
		abi.encodePacked(
			"\x19\x01",
			domainHash,
			getPinHash($input)
		)
	);
}
```

:::