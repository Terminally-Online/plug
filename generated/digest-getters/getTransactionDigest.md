# getTransactionDigest
        
Encode [Transaction](/base-types/Transaction) data into a digest hash.

## Parameters

- `$input` : [Transaction](/base-types/Transaction) : The `Transaction` data to encode.

## Returns

- `$digest` : `bytes32` : The digest hash of the encoded [Transaction](/base-types/Transaction) data.

## Onchain Implementation

::: code-group

``` solidity [Types.sol:getTransactionDigest]
function getTransactionDigest(
	Transaction memory $input
) public view virtual returns (bytes32 $digest) {
	$digest = keccak256(
		abi.encodePacked(
			"\x19\x01",
			domainHash,
			getTransactionHash($input)
		)
	);
}
```

:::