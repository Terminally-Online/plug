# getPlugsDigest
        
Encode [Plugs](/generated/base-types/Plugs) data into a digest hash that has been localized to the domain of the contract.

## Parameters

- `$input` : [Plugs](/generated/base-types/Plugs) : The `Plugs` data to encode.

## Returns

- `$digest` : `bytes32` : The digest hash of the encoded [Plugs](/generated/base-types/Plugs) data.

## Onchain Implementation

::: code-group

``` solidity [Types.sol:getPlugsDigest]
function getPlugsDigest(
	TypesLib.Plugs memory $input
) public view virtual returns (bytes32 $digest) {
	$digest = keccak256(
		abi.encodePacked(
			"\x19\x01",
			domainHash,
			getPlugsHash($input)
		)
	);
}
```

:::