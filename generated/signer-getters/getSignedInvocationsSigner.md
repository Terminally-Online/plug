# getSignedInvocationsSigner

Get the signer of a [SignedInvocations](/generated/base-types/SignedInvocations) data type.

## Parameters

- `$input` : [SignedInvocations](/generated/base-types/SignedInvocations) : The `SignedInvocations` data to encode.

## Returns

- `$signer` : `address` : The signer of the [SignedInvocations](/generated/base-types/SignedInvocations) data.

## Onchain Implementation

::: code-group

``` solidity [Types.sol:getSignedInvocationsSigner]
function getSignedInvocationsSigner(
	SignedInvocations memory $input
) public view virtual returns (address $signer) {
	$signer = getInvocationsDigest($input.invocations).recover(
		$input.signature
	);
}
```

:::