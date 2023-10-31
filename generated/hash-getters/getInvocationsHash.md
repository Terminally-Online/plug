# getInvocationsHash

Encode [Invocations](/base-types/Invocations) data into a packet hash and verify decoded [Invocations](/base-types/Invocations) data from a hash to verify type compliance and value-width alignment.

## Parameters

- `$input` : [Invocations](/base-types/Invocations) : The `Invocations` data to encode.

## Returns

- `$packetHash` : `bytes32` : The packet hash of the encoded [Invocations](/base-types/Invocations) data.

## Onchain Implementation

::: code-group

``` solidity [Types.sol:getInvocationsHash]
function getInvocationsHash(
	Invocations memory $input
) public pure virtual returns (bytes32 $packetHash) {
	$packetHash = keccak256(abi.encode(
		INVOCATIONS_TYPEHASH,
		getInvocationArrayHash($input.batch),
		getReplayProtectionHash($input.replayProtection)
	));
}
``` 

:::