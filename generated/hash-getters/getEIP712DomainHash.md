# getEIP712DomainHash

Encode [EIP712Domain](/base-types/EIP712Domain) data into a packet hash and verify decoded [EIP712Domain](/base-types/EIP712Domain) data from a hash to verify type compliance and value-width alignment.

## Parameters

- `$input` : [EIP712Domain](/base-types/EIP712Domain) : The `EIP712Domain` data to encode.

## Returns

- `$packetHash` : `bytes32` : The packet hash of the encoded [EIP712Domain](/base-types/EIP712Domain) data.

## Onchain Implementation

::: code-group

``` solidity [Types.sol:getEIP712DomainHash]
function getEIP712DomainHash(
	EIP712Domain memory $input
) public pure virtual returns (bytes32 $packetHash) {
	$packetHash = keccak256(abi.encode(
		EIP712_DOMAIN_TYPEHASH,
		keccak256(bytes($input.name)),
		keccak256(bytes($input.version)),
		$input.chainId,
		$input.verifyingContract
	));
}
``` 

:::