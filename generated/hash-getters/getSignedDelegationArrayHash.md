# getSignedDelegationArrayHash

Encode [SignedDelegation[]](/base-types/SignedDelegation[]) data into a packet hash and verify decoded [SignedDelegation[]](/base-types/SignedDelegation[]) data from a hash to verify type compliance and value-width alignment.

## Parameters

- `$input` : [SignedDelegation[]](/base-types/SignedDelegation[]) : The `SignedDelegation[]` data to encode.

## Returns

- `$packetHash` : `bytes32` : The packet hash of the encoded [SignedDelegation[]](/base-types/SignedDelegation[]) data.

## Onchain Implementation

::: code-group

``` solidity [Types.sol:getSignedDelegationArrayHash]
function getSignedDelegationArrayHash(
	SignedDelegation[] memory $input
)  public pure virtual returns (bytes32 $packetHash) {
	bytes memory encoded;

	uint256 i;
	uint256 length = $input.length;

	for (i; i < length;) {
		encoded = bytes.concat(
			encoded,
			getSignedDelegationHash($input[i])
		);

		unchecked { i++; }
	}
	
	$packetHash = keccak256(encoded);
}
``` 

:::