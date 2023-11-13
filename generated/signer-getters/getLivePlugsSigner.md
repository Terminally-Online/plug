# getLivePlugsSigner

Get the signer of a [LivePlugs](/generated/base-types/LivePlugs) data type.

## Parameters

- `$input` : [LivePlugs](/generated/base-types/LivePlugs) : The `LivePlugs` data to encode.

## Returns

- `$signer` : `address` : The signer of the [LivePlugs](/generated/base-types/LivePlugs) data.

## Onchain Implementation

::: code-group

``` solidity [Types.sol:getLivePlugsSigner]
function getLivePlugsSigner(
	LivePlugs memory $input
) public view virtual returns (address $signer) {
	$signer = getPlugsDigest($input.plugs).recover(
		$input.signature
	);
}
```

:::