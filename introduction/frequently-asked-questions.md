---
head:
  - - meta
    - property: og:title
      content: Frequently Asked Questions
  - - meta
    - name: description
      content: Explore the FAQ of the Plug framework and library.
  - - meta
    - property: og:description
      content: Explore the FAQ of the Plug framework and library.
---

# Frequently Asked Questions

While [Plug](/) is rather simple, it's new and that may lead to many questions that you have never had to ask before. However, it is understandable that you may have some questions. This page will answer some of the most common questions that you may have.

::: info

If you are using the documentation and feel that a question or answer belongs here, please do edit the page and submit a PR. We would love to hear from you! If you are looking for more information or have a question not answered here, please [reach out for help](https://twitter.com/nftchance).

:::

## Why would I want intents?

Native EVM transactions serve as a constant limiter of the throughput for protocols and users. Bound to the execution path of the blockchain, it is expected for users to not only accept, but wait for the less than ideal performance of the EVM. Declarative transactions (intents) allow you to break free of this limitation and access a much more performant experience with scheduled, automated, and conditional tranasctions to enable more effective transaction settlement.

## Can I use Plug with protocols already deployed?

Yes! Plug will work with any contract whether it was deployed years in the past, in the present or even years in the future.

## If I want basic support what do I need?

To submit an intent all you need to do is prepare the declaration and submit it to the pool of [Solvers](/core/solvers). You can choose to submit things through a vault for value-based actions or a generalized and permisionless router.

## Are intents onchain?

Magically, intents are only ever written onchain once they can be submit to the blockchain they declare an action on. As the base declaration happens with a signature there is no immediate or rent-seeking cost. The only time any cost of execution is incurred is when the transactions are actually run onchain.

## Do I have to use a specific API or Relay?

No! While there is a first-party API available to you that streamlines the collection, management and distribution of the incoming plugs related to your protocol, it is not required. You can use any API or Relay that you would like to consume the [LivePlugs](/generated/base-types/LivePlugs) and distribute them to the desired network of [Solvers](/core/solvers).

## Doesn't this just make my protocol more complex?

Quite the opposite! By utilizing [Plug](/) developers are able to remove the management of EVM deficiencies from protocols and provide your users with a much more performant and focused experience. For instance, many protocols require the implementation of time-based limitations. When using [Plug](/) that is no longer the case.

## Does my protocol still support native transactions?

Yes! [Plug](/) is designed to be a solution that allows protocols and users to access the benefits of traditional transactions as well as intents. This allows you to support the best of both worlds and provide your users with the best experience possible in every case.
