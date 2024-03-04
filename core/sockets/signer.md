# signer

Retrieves and validates the `signer` of a signed [Plugs](/generated/base-types/Plugs) bundle is accepted as a permission granter of the executing [Socket](/core/sockets).

## Parameters

- `$livePlugs`: The [LivePlugs](/generated/base-types/LivePlugs) bundle that contains the Plugs data and signature.

## Returns

- `$signer`: The address of the account that signed the message.

## Onchain Implementation

With the call of `signer` validation can take place both offchain and onchain. Offchain, this simple function provides the ability to not even bother simulating invalid intents. Meanwhile, when called onchain the [Socket](/core/sockets) will automatically verify that the individual that signed the bundle is allowed to execute calls on behalf of the [Socket](/core/sockets).

```solidity [./Plug.Socket.sol]
function signer(PlugTypesLib.LivePlugs calldata $livePlugs)
  external
  view
  returns (address $signer)
{
  /// @dev Determine the address that signed the Plug bundle.
  $signer = getLivePlugsSigner($livePlugs);

  /// @dev Confirm the signer of the bundle is valid.
  require(_enforceSigner($signer), "Plug:invalid-signer");
}
```

When calling this, if the signer is not allowed the call will simply revert.
