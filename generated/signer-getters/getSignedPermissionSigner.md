# getSignedPermissionSigner

Get the signer of a [SignedPermission](/generated/base-types/SignedPermission) data type.

## Parameters

- `$input` : [SignedPermission](/generated/base-types/SignedPermission) : The `SignedPermission` data to encode.

## Returns

- `$signer` : `address` : The signer of the [SignedPermission](/generated/base-types/SignedPermission) data.

## Onchain Implementation

::: code-group

``` solidity [Types.sol:getSignedPermissionSigner]
function getSignedPermissionSigner(
	SignedPermission memory $input
) public view virtual returns (address $signer) {
	$signer = getPermissionDigest($input.permission).recover(
		$input.signature
	);
}
```

:::