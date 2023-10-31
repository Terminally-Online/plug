# getSignedDelegationSigner

Get the signer of a [SignedDelegation](/base-types/SignedDelegation) data type.

## Parameters

- `$input` : [SignedDelegation](/base-types/SignedDelegation) : The `SignedDelegation` data to encode.

## Returns

- `$signer` : `address` : The signer of the [SignedDelegation](/base-types/SignedDelegation) data.

## Onchain Implementation

::: code-group

``` solidity [Types.sol:getSignedDelegationSigner]
function getSignedDelegationSigner(
	SignedDelegation memory $input
) public view virtual returns (address $signer) {
	$signer = getDelegationDigest($input.delegation).recover(
		$input.signature
	);
}
```

:::