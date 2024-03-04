# plug

Submit a bundle of [LivePlugs](/generated/base-types/LivePlugs) to a [Socket](/core/sockets) to execute the declared intent.

## Parameters

- `$livePlugs`: The signed bundle of [Plugs](/generated/base-types/Plugs) to execute.
- `$livePlugs`:overloaded`: A batch of signed bundles of [Plugs](/generated/base-types/Plugs) to execute.

## Returns

- `$results`: An array of results representing the [Plug](/generated/base-types/Plug) bundle.
- `$results`: `overloaded`: A nested array of results representing the batched [Plug](/generated/base-types/Plug) bundles.

## Onchain Implementation

With `plug` publicly exposed any permissioned [Solver](/core/solvers) can submit and execute bundles of signed [Plugs](/generated/base-types/Plugs). As an overloaded function, a [Solver](/core/solvers) has the option to execute a single bundle of [Plugs](/generated/base-types/Plugs) or settle many in a single batch.

When called, the amount of `gas` the transaction has remaining is snapshotted for future use in compensation. The [Router](/core/routers) then interfaces with the [Socket](/core/sockets). To confirm the bundle is valid the signer is retrieved from the [Socket](/core/sockets) and passed into the [Socket](/core/sockets) [plug](/core/sockets/plug) function for execution.

The simplest interaction method comes in the form of single-intent settlement with:

::: code-group

```solidity [plug (single)]
function plug(
  PlugTypesLib.LivePlugs calldata $livePlugs
) public payable virtual returns (
  bytes[] memory $results
) {
  uint256 gas = gasleft();

  PlugSocketInterface socket =
      PlugSocketInterface($livePlugs.plugs.socket);

  require(
      msg.sender == $livePlugs.plugs.executor
          || $livePlugs.plugs.executor == address(0),
      "Plug:invalid-executor"
  );

  address signer = socket.signer($livePlugs);

  $results = socket.plug($livePlugs.plugs, signer, gas);
}
```

```solidity [plug (batch)]
function plug(
  PlugTypesLib.LivePlugs[] calldata $livePlugs
) public payable virtual returns (
  bytes[][] memory $results
) {
  uint256 i;
  uint256 length = $livePlugs.length;
  $results = new bytes[][](length);

  for (i; i < length; i++) {
    $results[i] = plug($livePlugs[i]);
  }
}
```

:::

In the case a [Solver](/core/solvers) reaches a high level of throughput and wants to settle several bundles at once, the `batch` version of the function can be utilized:

With the access point for a single [LivePlugs](/generated/base-types/LivePlugs) bundle defined above, batching is as simple as passing an array of bundles and getting a batch of `result` arrays in response.
