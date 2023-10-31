# getDelegationHash

Encode [Delegation](/base-types/Delegation) data into a packet hash and verify decoded [Delegation](/base-types/Delegation) data from a hash to verify type compliance and value-width alignment.

## Parameters

- `$input` : [Delegation](/base-types/Delegation) : The `Delegation` data to encode.

## Returns

- `$packetHash` : `bytes32` : The packet hash of the encoded [Delegation](/base-types/Delegation) data.

## Onchain Implementation

::: code-group

``` solidity [Types.sol:getDelegationHash]
function getDelegationHash(
	Delegation memory $input
) public pure virtual returns (bytes32 $packetHash) {
	$packetHash = keccak256(abi.encode(
		DELEGATION_TYPEHASH,
		$input.delegate,
		$input.authority,
		getCaveatArrayHash($input.caveats),
		$input.salt
	));
}
``` 

:::