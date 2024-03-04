# enforceFuse

## Parameters

- `$terms`: The active rules relative to the [Fuse](/core/fuses).
- `$current`: The transaction to execute.
- `$plugsHash`: The unique identifier of the intent signed.

## Returns

- `$through`: The data of the transaction being executed.

## Onchain Implementation

Every [Fuse](/core/fuses) varies in the rules that it applies which means the internal logic is always different. Although that is the case, the top level interface of enforcement enabled through `enforceFuse` is always the same so that the [Socket](/core/sockets) can call it without explicit shape knowledge.

```solidity
function enforceFuse(
  bytes calldata $terms,
  PlugTypesLib.Current calldata $current,
  bytes32
)
  public
  view
  override
  returns (bytes memory $through)
{
  /// @dev Decode the terms to get the logic operator and threshold.
  (uint256 $threshold) = decode($live);

  /// @dev Confirm the intent has not expired.
  require($threshold < block.number, "PlugBlockNumberFuse:expired");

  /// @dev Continue the pass through.
  $through = $current.data;
}
```
