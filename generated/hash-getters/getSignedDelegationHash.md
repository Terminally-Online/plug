# getSignedDelegationHash

Encode [SignedDelegation](/base-types/SignedDelegation) data into a packet hash and verify decoded [SignedDelegation](/base-types/SignedDelegation) data from a hash to verify type compliance and value-width alignment.

## Parameters

- `$input` : [SignedDelegation](/base-types/SignedDelegation) : The `SignedDelegation` data to encode.

## Returns

- `$packetHash` : `bytes32` : The packet hash of the encoded [SignedDelegation](/base-types/SignedDelegation) data.

## Onchain Implementation

::: code-group

``` solidity [Types.sol:getSignedDelegationHash]
function getSignedDelegationHash(
	SignedDelegation memory $input
) public pure virtual returns (bytes32 $packetHash) {
	$packetHash = keccak256(abi.encode(
		SIGNED_DELEGATION_TYPEHASH,
		getDelegationHash($input.delegation),
		keccak256($input.signature)
	));
}
``` 

:::