# getReplayProtectionHash

Encode [ReplayProtection](/base-types/ReplayProtection) data into a packet hash and verify decoded [ReplayProtection](/base-types/ReplayProtection) data from a hash to verify type compliance and value-width alignment.

## Parameters

- `$input` : [ReplayProtection](/base-types/ReplayProtection) : The `ReplayProtection` data to encode.

## Returns

- `$packetHash` : `bytes32` : The packet hash of the encoded [ReplayProtection](/base-types/ReplayProtection) data.

## Onchain Implementation

::: code-group

``` solidity [Types.sol:getReplayProtectionHash]
function getReplayProtectionHash(
	ReplayProtection memory $input
) public pure virtual returns (bytes32 $packetHash) {
	$packetHash = keccak256(abi.encode(
		REPLAY_PROTECTION_TYPEHASH,
		$input.nonce,
		$input.queue
	));
}
``` 

:::