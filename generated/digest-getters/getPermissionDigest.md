# getPermissionDigest
        
Encode [Permission](/generated/base-types/Permission) data into a digest hash that has been localized to the domain of the contract.

## Parameters

- `$input` : [Permission](/generated/base-types/Permission) : The `Permission` data to encode.

## Returns

- `$digest` : `bytes32` : The digest hash of the encoded [Permission](/generated/base-types/Permission) data.

## Onchain Implementation

::: code-group

``` solidity [Types.sol:getPermissionDigest]
function getPermissionDigest(
	Permission memory $input
) public view virtual returns (bytes32 $digest) {
	$digest = keccak256(
		abi.encodePacked(
			"\x19\x01",
			domainHash,
			getPermissionHash($input)
		)
	);
}
```

:::