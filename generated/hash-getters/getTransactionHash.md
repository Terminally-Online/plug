# getTransactionHash

Encode [Transaction](/base-types/Transaction) data into a packet hash and verify decoded [Transaction](/base-types/Transaction) data from a hash to verify type compliance and value-width alignment.

## Parameters

- `$input` : [Transaction](/base-types/Transaction) : The `Transaction` data to encode.

## Returns

- `$packetHash` : `bytes32` : The packet hash of the encoded [Transaction](/base-types/Transaction) data.

## Onchain Implementation

::: code-group

``` solidity [Types.sol:getTransactionHash]
function getTransactionHash(
	Transaction memory $input
) public pure virtual returns (bytes32 $packetHash) {
	$packetHash = keccak256(abi.encode(
		TRANSACTION_TYPEHASH,
		$input.to,
		$input.gasLimit,
		keccak256($input.data)
	));
}
``` 

:::