---
head:
    - - meta
      - property: og:title
        content: Single Lane Nonces
    - - meta
      - name: description
        content: A single-lane nonce is used to ensure that a reference is only seen once.
    - - meta
      - property: og:description
        content: A single-lane nonce is used to ensure that a reference is only seen once.
---

# Single Lane Nonces

With native accounts and transactions a single-lane nonce is used to ensure that a transaction is only executed once.

## Incremental Transaction Nonces

With native accounts and transactions a single-lane nonce is used to ensure that a transaction is only executed once:

-   Submit a transaction, your nonce is incremented by `1`.
-   Submit another transaction with a lower nonce, it's rejected.

This is an low-overhead way to ensure that a transaction is never replayed.

## The Technical Benefit

Although very simple, a single nonce is extremely effective at ensuring uniqueness while preventing unintended reuse. By consuming an incrementally increasing nonce, a user can be sure that a transaction:

-   **Is never executed out of order (replayed or skipped).**

This simple rule result in a system that makes [double-spending](https://en.wikipedia.org/wiki/Double-spending) and [replay attacks](https://en.wikipedia.org/wiki/Replay_attack) impossible. How? Let's look at a simple example where we have $100 and we'd like to send $50 to `Alice` and $50 to `Bob`:

-   Send $50 to `Alice`, your nonce is incremented by `1`.
-   Send $50 to `Bob`, your nonce is incremented by `1`.

Now, let's say that `Alice` is a bad actor and she tries to replay the transaction that sent her $50. What happens?

-   Your nonce is already incremented by `1`, it's rejected.

The transaction to `Alice` is rejected because the nonce is not the next incrementally increasing number. This is a very simple example, but it illustrates the power of a single-lane nonce.

In practice, the `Solidity` needed to power this is as simple as the concept:

::: code-group

```solidity 2,5 [Nonces.sol]
/// @dev Single lane nonce implementation.
mapping(address sender => uint256 nonce) public senderToNonce;

function execute(uint256 $nonce, bytes calldata $message) external {
    require($nonce == ++senderToNonce[msg.sender], "Invalid nonce.");
    /// Your replay-protected logic here.
}
```

Just two lines of code and you have a system that ensures that a transaction is never replayed or skipped. This is the core of the `Emporium` permissioning system and is extended to provide a powerful and flexible permissioning system.

:::

::: tip

The detail of validators and the underlying blockchain is not important to understand the benefit of a single-lane nonce. The only thing that matters is that the nonce is incremented by `1` for every transaction.

:::

## The Experience Benefit

Before the advent of blockchain technology, double spending was a significant challenge in digital transactions. Traditional systems like banks, PayPal, or credit cards relied heavily on centralized authorities to prevent double spending. These systems require users to trust a third party to validate transactions and ensure that the same money isn't spent twice.

While single-lane nonces have proven effective in combating double spending in decentralized systems, traditional national currencies and banking systems have historically been vulnerable to various forms of this problem. In the context of national economies, double spending often manifests not as an individual attempting to spend the same dollar bill twice, but more broadly as systemic issues that have a profound economic impact.

The fractional reserve banking system, used globally, allows banks to lend more money than they actually have in reserve. In a way, this can be viewed as a form of institutionalized double spending. While this system is regulated and generally stable, it can lead to a cascade of problems if trust in the banking sector erodes, such as during the financial crisis of 2007-2008.

National governments themselves can engage in a form of double spending by printing more money than is backed by their reserves or economic output, risking hyperinflation. Countries like Zimbabwe and Venezuela have experienced the dire economic consequences of such policies.

In both the banking sector and governmental monetary policy, the risk of double spending and the broader systemic risks stem from centralized control. The lack of a foolproof, trustless system like blockchain's single-lane nonces means that citizens must rely on these centralized institutions to act responsibly, a trust that has been broken numerous times throughout history.

::: info

The risk of double spending in traditional financial systems reveals their central points of failure and highlights the need for more secure, decentralized alternatives. It underscores the significance of innovations like single-lane nonces in creating a more robust financial infrastructure.

:::

In digital currencies without a central authority, the risk of double spending increases significantly. If there were no mechanisms like a single-lane nonce, a malicious actor could attempt to spend the same digital coin in more than one transaction. This would severely undermine the integrity and trust in the system.

The use of single-lane nonces in blockchain architectures directly tackles this problem. It not only prevents a transaction from being processed more than once but also thwarts any attempt to double spend. This is because each transaction is uniquely identified and ordered through its nonce, rendering previous transactions obsolete for replay purposes allowing individuals to transact with confidence.
