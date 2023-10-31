# getReplayProtectionDigest
        
Encode [ReplayProtection](/base-types/ReplayProtection) data into a digest hash.

## Parameters

- `$input` : [ReplayProtection](/base-types/ReplayProtection) : The `ReplayProtection` data to encode.

## Returns

- `$digest` : `bytes32` : The digest hash of the encoded [ReplayProtection](/base-types/ReplayProtection) data.

## Onchain Implementation

::: code-group

``` solidity [Types.sol:getReplayProtectionDigest]
function getReplayProtectionDigest(
	ReplayProtection memory $input
) public view virtual returns (bytes32 $digest) {
	$digest = keccak256(
		abi.encodePacked(
			"\x19\x01",
			domainHash,
			getReplayProtectionHash($input)
		)
	);
}
```

:::