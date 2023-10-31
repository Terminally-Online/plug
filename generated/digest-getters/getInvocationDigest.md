# getInvocationDigest
        
Encode [Invocation](/base-types/Invocation) data into a digest hash.

## Parameters

- `$input` : [Invocation](/base-types/Invocation) : The `Invocation` data to encode.

## Returns

- `$digest` : `bytes32` : The digest hash of the encoded [Invocation](/base-types/Invocation) data.

## Onchain Implementation

::: code-group

``` solidity [Types.sol:getInvocationDigest]
function getInvocationDigest(
	Invocation memory $input
) public view virtual returns (bytes32 $digest) {
	$digest = keccak256(
		abi.encodePacked(
			"\x19\x01",
			domainHash,
			getInvocationHash($input)
		)
	);
}
```

:::