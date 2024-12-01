---
head:
    - - meta
      - property: og:title
        content: getLivePlugsHash
    - - meta
      - name: description
        content: Encode a LivePlugs into a hash and verify the decoded data to verify type compliance.
    - - meta
      - property: og:description
        content: Encode a LivePlugs into a hash and verify the decoded data to verify type compliance.
notes:
    - - author: Auto generated by @nftchance/plug-types/cli
---
        
# getLivePlugsHash

Encode a [LivePlugs](/generated/base-types/LivePlugs) into a hash and verify the decoded [LivePlugs](/generated/base-types/LivePlugs) data from a hash to verify type compliance.

## Parameters

- `$input` : [LivePlugs](/generated/base-types/LivePlugs) : The `LivePlugs` data to encode.

## Returns

- `$typeHash` : `bytes32` : The packet hash of the encoded [LivePlugs](/generated/base-types/LivePlugs) data.

## Onchain Implementation

With `getLivePlugsHash` you can call the function as a `read` and get the encoded data back as a hash. 
        
This is helpful in times when you need to build a message hash without tracking down all the types as well as when you need to verify a signed message hash containing a `LivePlugs` data type.

::: code-group

``` solidity [Types.sol:getLivePlugsHash]
function getLivePlugsHash(
	TypesLib.LivePlugs memory $input
) public pure virtual returns (bytes32 $typeHash) {
	$typeHash = keccak256(abi.encode(
		LIVE_PLUGS_TYPEHASH,
		getPlugsHash($input.plugs),
	keccak256($input.signature)
	));
}
``` 

:::