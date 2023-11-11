# getSignedIntentsSigner

Get the signer of a [SignedIntents](/generated/base-types/SignedIntents) data type.

## Parameters

- `$input` : [SignedIntents](/generated/base-types/SignedIntents) : The `SignedIntents` data to encode.

## Returns

- `$signer` : `address` : The signer of the [SignedIntents](/generated/base-types/SignedIntents) data.

## Onchain Implementation

::: code-group

``` solidity [Types.sol:getSignedIntentsSigner]
function getSignedIntentsSigner(
	SignedIntents memory $input
) public view virtual returns (address $signer) {
	$signer = getIntentsDigest($input.intents).recover(
		$input.signature
	);
}
```

:::