# getCaveatDigest
        
Encode [Caveat](/base-types/Caveat) data into a digest hash.

## Parameters

- `$input` : [Caveat](/base-types/Caveat) : The `Caveat` data to encode.

## Returns

- `$digest` : `bytes32` : The digest hash of the encoded [Caveat](/base-types/Caveat) data.

## Onchain Implementation

::: code-group

``` solidity [Types.sol:getCaveatDigest]
function getCaveatDigest(
	Caveat memory $input
) public view virtual returns (bytes32 $digest) {
	$digest = keccak256(
		abi.encodePacked(
			"\x19\x01",
			domainHash,
			getCaveatHash($input)
		)
	);
}
```

:::