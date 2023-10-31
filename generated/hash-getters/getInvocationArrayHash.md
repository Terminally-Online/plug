# getInvocationArrayHash

Encode [Invocation[]](/base-types/Invocation[]) data into a packet hash and verify decoded [Invocation[]](/base-types/Invocation[]) data from a hash to verify type compliance and value-width alignment.

## Parameters

- `$input` : [Invocation[]](/base-types/Invocation[]) : The `Invocation[]` data to encode.

## Returns

- `$packetHash` : `bytes32` : The packet hash of the encoded [Invocation[]](/base-types/Invocation[]) data.

## Onchain Implementation

::: code-group

``` solidity [Types.sol:getInvocationArrayHash]
function getInvocationArrayHash(
	Invocation[] memory $input
)  public pure virtual returns (bytes32 $packetHash) {
	bytes memory encoded;

	uint256 i;
	uint256 length = $input.length;

	for (i; i < length;) {
		encoded = bytes.concat(
			encoded,
			getInvocationHash($input[i])
		);

		unchecked { i++; }
	}
	
	$packetHash = keccak256(encoded);
}
``` 

:::