# Coils

<span style="color: rgba(0,0,0,0.6)">[Coils](/concepts/coils) are a proprietary piece of the protocol that enables a vm-like experience at the smart contract level and during transaction execution. With this in place you and every other user has the freedom to do things not possible anywhere else.</span>

## State of The Tech

### Transactions

On Ethereum there has been several updates to the data contained inside of a transaction, but at all times, past and present, a transaction has been a singular entity solely focused on the effects of itself. At the simplest level, a transaction defines a set of fields like:

```tsx
{
  to: "0x0",
  data: "0x0",
  value: "0x0",
}
```

> The [actions](/concepts/actions) living inside each Plug are defined to mirror the singular definition of a transaction with additional metadata to power specific feature sets.

Because of this, typical EOA experiences have required running multiple transactions any time the user wants to do something more than transferring ETH. Of course, that means upwards of 99% of transactions are multi-step experiences.

### Bundles (Multicall)

Due to the increasing complexity in Ethereum protocols and the improving quality of user experience in applications, Bundlers such as Multicall were created so that users could execute multiple transactions at once. There are a few nuanced limitations with Multicall, but for the most part it is a simple bundler that lets you run multiple transactions at a time.

Now, instead of a user running a single transaction they are really running a set of transactions that can be visualized like:

```tsx
[{
  to: "0x0",
  data: "0x0",
  value: "0x0",
},{
  to: "0x0",
  data: "0x0",
  value: "0x0",
},{
  to: "0x0",
  data: "0x0",
  value: "0x0",
},{
  to: "0x0",
  data: "0x0",
  value: "0x0",
}]
```

You can define a bundle and run it all once. Multicall is great and solves for the basic need of being able to dumbly run multiple transactions at once. There is no cross-transaction communication or consideration.

Atomic execution of a blind bundle has shortcomings that drastically limit what can be done onchain, though. This is because with Multicall you must define all the data of your transaction before you run it. Your transaction cannot react to the current onchain state or utilize the return data of each previous transaction run in a bundle.

This single limitation results in the inability to do simple workflows like:

![Need for coils](/assets/coils-problem.png)

With Multicall this cannot be safely done because we do not know the amount that will be received during the swap. Although Multicall enables simple bundling execution, it has zero support for arbitrary transaction data building and that is a deal breaker.

## Coils

With the introduction of [coils](/concepts/coils) a new paradigm is introduced where transactions can respond in realtime during execution.

Inspired by [Multicall](https://github.com/mds1/multicall3) and [Weiroll](https://github.com/weiroll/weiroll), Coils are designed to slice a piece of source data and insert it into yet to be executed transaction data. This means that instead of having to know every piece of data before executing your transaction with Multicall, the data needed for your transaction can be dynamically filled in as your transaction executes. 

To do this, we use a really simple language-first approach that keeps a sentence as the source-of-truth for every piece of the system that interactions with an action like:

![Anatomy of coils](/assets/coils-anatomy.png)

Instead of thinking about onchain actions as functions, we abstract it into a higher layer that surfaces as a sentence any human with decent literacy can follow and comprehend. There's no confusing hashes. No complex types to manage. Just simple logic statements and actions that can be wired together.

With this action sentence, a user has the ability to read the balance of a token held by a user during the atomic execution of the bundle so that the most up to date data is always used and informs the subsequent transactions like:

![Anatomy of coils](/assets/coils-discharge.png)

Now, when after the first swap and before the second swap, the data of the second swap transaction will be updated to reflect the value retrieved from the balance read without you, the user, having to do anything special.

With just these two simple sentences you already have more functionality than any market-accessible aggregator, transaction builder, or automation service in all of crypto. Coils enable a world where there is no fair comparison to existing competitors because you have been enabled in ways that the industry has never seen.
