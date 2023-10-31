# getInvocationsDigest
        
Encode [Invocations](/generated/base-types/Invocations) data into a digest hash that has been localized to the domain of the contract.

## Parameters

- `$input` : [Invocations](/generated/base-types/Invocations) : The `Invocations` data to encode.

## Returns

- `$digest` : `bytes32` : The digest hash of the encoded [Invocations](/generated/base-types/Invocations) data.

## Onchain Implementation

::: code-group

``` solidity [Types.sol:getInvocationsDigest]
function getInvocationsDigest(
	Invocations memory $input
) public view virtual returns (bytes32 $digest) {
	$digest = keccak256(
		abi.encodePacked(
			"\x19\x01",
			domainHash,
			getInvocationsHash($input)
		)
	);
}
```

:::