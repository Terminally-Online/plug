# getSignedDelegationDigest
        
Encode [SignedDelegation](/base-types/SignedDelegation) data into a digest hash.

## Parameters

- `$input` : [SignedDelegation](/base-types/SignedDelegation) : The `SignedDelegation` data to encode.

## Returns

- `$digest` : `bytes32` : The digest hash of the encoded [SignedDelegation](/base-types/SignedDelegation) data.

## Onchain Implementation

::: code-group

``` solidity [Types.sol:getSignedDelegationDigest]
function getSignedDelegationDigest(
	SignedDelegation memory $input
) public view virtual returns (bytes32 $digest) {
	$digest = keccak256(
		abi.encodePacked(
			"\x19\x01",
			domainHash,
			getSignedDelegationHash($input)
		)
	);
}
```

:::