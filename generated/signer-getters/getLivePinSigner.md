# getLivePinSigner

Get the signer of a [LivePin](/generated/base-types/LivePin) data type.

## Parameters

- `$input` : [LivePin](/generated/base-types/LivePin) : The `LivePin` data to encode.

## Returns

- `$signer` : `address` : The signer of the [LivePin](/generated/base-types/LivePin) data.

## Onchain Implementation

::: code-group

``` solidity [Types.sol:getLivePinSigner]
function getLivePinSigner(
	LivePin memory $input
) public view virtual returns (address $signer) {
	$signer = getPinDigest($input.pin).recover(
		$input.signature
	);
}
```

:::