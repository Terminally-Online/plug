# getIntentsDigest
        
Encode [Intents](/generated/base-types/Intents) data into a digest hash that has been localized to the domain of the contract.

## Parameters

- `$input` : [Intents](/generated/base-types/Intents) : The `Intents` data to encode.

## Returns

- `$digest` : `bytes32` : The digest hash of the encoded [Intents](/generated/base-types/Intents) data.

## Onchain Implementation

::: code-group

``` solidity [Types.sol:getIntentsDigest]
function getIntentsDigest(
	Intents memory $input
) public view virtual returns (bytes32 $digest) {
	$digest = keccak256(
		abi.encodePacked(
			"\x19\x01",
			domainHash,
			getIntentsHash($input)
		)
	);
}
```

:::