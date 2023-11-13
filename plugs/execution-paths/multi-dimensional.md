---
head:
  - - meta
    - property: og:title
      content: Multi-Dimensional Nonces
  - - meta
    - name: description
      content: With multi-dimensional nonces, you can revoke a set of pins without impacting the rest of your queues.
  - - meta
    - property: og:description
      content: With multi-dimensional nonces, you can revoke a set of pins without impacting the rest of your queues.
---

# Multi-Dimensional Nonces

We just covered the details of a [Single Lane](/plugs/execution-paths/single-lane) `nonce`, but what about a multi-dimensional `queue-nonce`? What is it and why is it important?

## Nonce Queues

`Plug` uses a multi-dimensional `queue-nonce` architecture that introduces the ability to have multiple `nonces` for every account on every protocol. In practice this looks like:

:::code-group

```solidity [Nonces.sol]
/// @dev Single lane nonce implementation.
mapping(address sender => uint256 nonce) public senderToNonce;

/// @dev Multi-dimensional nonce implementation.
/// @notice A `uint256` is functionally the same as `bytes32`.
mapping(address sender => mapping(uint256 queue => uint256 nonce)) public senderToQueueToNonce;
```

:::

**Simple, right?** This means instead of using a one-way backroad for replay protection, users have the ability to step on the highway that has multiple `queues` each with their own `nonce` and rate of flow.

With this functionality available, a user maintains the ability to revoke a set of pins and even expire an entire `queue` of pins without impacting every other set of pins and plugs previously signed.

- Submit a transaction in queue `3`, your `queue-nonce` is incremented by 1.
- Submit another transaction in queue `3` with a lower `queue-nonce`, it's rejected.
- Submit a transaction in queue `4` with a lower `nonce` than the previous transaction in queue `3`, it's accepted.

Each `queue` is independent and can be incremented without impacting the others.

## The Technical Benefit

To illustrate the benefit of [nonce queues](#nonce-queues) let's look at a simple example where we'd like to expire an plugs already distributed where:

- The current lane `nonce` is `45`.
- We have `10` active plugs (nonces `45-54`).
- We'd like to expire the `6th` pending intent (nonce `50`).

**With single-lane nonces:** We would increment our `nonce` to `50` and all of our previous `nonces` would be invalidated. This seems great at first, but in reality nonces `45-49` should still be valid and active.

**With multi-dimensional nonces:** If the `queue-nonce` of the order we'd like to expire is `50-1`, we simply roll to `50-2`. Critically, every `queue` outside of `50` is still valid and active completely independent of the changes just made. This means that we can expire a set of pins without impacting the rest of our previous declarations.

Of course, this is a very simple example and the benefit of this functionality entirely depends on the complexity of the protocol and app-level implementation, but the core concept remains.

## The Experience Benefit

The technical benefit of [nonce queues](#nonce-queues) is clear, but the experience benefit is even more important. With multi-dimensional queues users of EVM blockchains unlock embedded access control that is not only more secure, but more flexible and user-friendly resulting in:

- Lower wasted gas money.
- Lower counterparty risk.
- Lower smart contract risk.
- Lower exposure to bad actors.
- Lower chance of human error.

::: tip

It is very important to note that this benefit is really only possible due to the underlying single-lane system of EVM blockchains. While an amazing unlock alone, it is not a replacement for the security of the underlying blockchain.

:::
