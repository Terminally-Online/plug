# getSignedInvocationsDigest
        
Encode [SignedInvocations](/base-types/SignedInvocations) data into a digest hash.

## Parameters

- `$input` : [SignedInvocations](/base-types/SignedInvocations) : The `SignedInvocations` data to encode.

## Returns

- `$digest` : `bytes32` : The digest hash of the encoded [SignedInvocations](/base-types/SignedInvocations) data.

## Onchain Implementation

::: code-group

``` solidity [Types.sol:getSignedInvocationsDigest]
function getSignedInvocationsDigest(
	SignedInvocations memory $input
) public view virtual returns (bytes32 $digest) {
	$digest = keccak256(
		abi.encodePacked(
			"\x19\x01",
			domainHash,
			getSignedInvocationsHash($input)
		)
	);
}
```

:::