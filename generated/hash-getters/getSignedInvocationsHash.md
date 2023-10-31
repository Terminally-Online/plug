# getSignedInvocationsHash

Encode [SignedInvocations](/base-types/SignedInvocations) data into a packet hash and verify decoded [SignedInvocations](/base-types/SignedInvocations) data from a hash to verify type compliance and value-width alignment.

## Parameters

- `$input` : [SignedInvocations](/base-types/SignedInvocations) : The `SignedInvocations` data to encode.

## Returns

- `$packetHash` : `bytes32` : The packet hash of the encoded [SignedInvocations](/base-types/SignedInvocations) data.

## Onchain Implementation

::: code-group

``` solidity [Types.sol:getSignedInvocationsHash]
function getSignedInvocationsHash(
	SignedInvocations memory $input
) public pure virtual returns (bytes32 $packetHash) {
	$packetHash = keccak256(abi.encode(
		SIGNED_INVOCATIONS_TYPEHASH,
		getInvocationsHash($input.invocations),
		keccak256($input.signature)
	));
}
``` 

:::