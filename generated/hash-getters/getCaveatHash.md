# getCaveatHash

Encode [Caveat](/base-types/Caveat) data into a packet hash and verify decoded [Caveat](/base-types/Caveat) data from a hash to verify type compliance and value-width alignment.

## Parameters

- `$input` : [Caveat](/base-types/Caveat) : The `Caveat` data to encode.

## Returns

- `$packetHash` : `bytes32` : The packet hash of the encoded [Caveat](/base-types/Caveat) data.

## Onchain Implementation

::: code-group

``` solidity [Types.sol:getCaveatHash]
function getCaveatHash(
	Caveat memory $input
) public pure virtual returns (bytes32 $packetHash) {
	$packetHash = keccak256(abi.encode(
		CAVEAT_TYPEHASH,
		$input.enforcer,
		keccak256($input.terms)
	));
}
``` 

:::