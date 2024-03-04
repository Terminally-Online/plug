# solve

With a generalized [Solver](/core/solvers) one has the capability to settle intents on the [Router](/core/routers) with multiple accounts that are working to power their own key pieces of functionality. This means, that instead of a single monolithic algorithm to solve for every edge case there are multiple variants running to power each key piece of logic.

With this though, the core onchain functionality of a [Solver](/core/solvers) stems from the [Router](/core/routers) itself.

## Parameters

- `$livePlugs`: A batch of signed bundles of [Plugs](/generated/base-types/Plugs) to execute.

## Returns

- `$results`: A nested array of results representing the batched [Plug](/generated/base-types/Plug) bundles.

## Onchain Implementation

Importantly, the implementation of [Solvers](/core/solvers) is abstracted to enable additional pre-execution logic to ensure that everything runs smoothly at all times. If [Solvers](/core/solvers) were forced to be EOAs, then the ability to interact would be significantly hampered as well as risk of access-loss would grow exponentially.

To power this, a [Solver](/core/solvers) has the ability to declare multiple internal executors as well as additional internal logic to smoothen the process of simulation and execution. At it's core though, the pattern follows a simple pattern to forward the calls with:

```solidity
function solve(
  uint256 $version,
  PlugTypesLib.LivePlugs[] calldata $livePlugs
) public payable virtual onlyExecutor(msg.sender) returns (
  bytes[][] memory $results
) {
  routers[$version].plug($livePlugs);
}
```

Realistically, most users will never need to run their own [Solver](/core/solvers). Though, in the case that it is desired one can do so by referencing the canonical reference.
