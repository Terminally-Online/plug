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

Getting started with [Plug](/) is quite simple. However, it is understandable that you may have some questions. This page will answer some of the most common questions that you may have.

- [Why would I want intents?](#why-would-i-want-intents)
- [Can I use Plug with smart contracts that are already deployed?](#can-i-use-plug-with-protocols-already-deployed)
- [If I want basic support what do I need?](#if-i-want-basic-support-what-do-i-need)
- [Is there a fee?](#is-there-a-fee)
- [Do I have to use a specific API or Relay?](#do-i-have-to-use-a-specific-api-or-relay)
- [Doesn't this just make my protocol more complex?](#doesn-t-this-just-make-my-protocol-more-complex)
- [Does my protocol still support native transactions?](#does-my-protocol-still-support-native-transactions)
- [Bro... These docs are huge! Can I get a TL;DR?](#bro-these-docs-are-huge-can-i-get-a-tl-dr)

::: info

If you are using the documentation and feel that a question or answer belongs here, please do edit the page and submit a PR. We would love to hear from you! Otherwise, this page is regularly updated with new questions and answers so check back often.

:::

## Why would I want intents?

Native EVM transactions serve as a constant limiter of the throughput for your protocol. Bound to the execution path of the blockchain it is deployed on your users are forced to not only accept, but wait for the less than ideal performance of the EVM. Declarative transactions allow you to break free of this limitation and provide your users with a much more performant experience to prevent significant fund loss and waste.

## Can I use Plug with protocols already deployed?

Yes! Plug will work with any contract whether it was deployed years in the past, in the present or even years in the future.

## If I want basic support what do I need?

To submit an intent all you need to do is prepare the declaration and submit it to the pool of Executors. You can choose to execute things through a vault for value-based actions or a generalized and permisionless router.

## Is there a fee?

The fee of your intent declaration is mostly up to the user. The execution of intents is incentivized with this fee giving Executors a reason to execute your transaction as close to the time that all the conditions have been met.

## Do I have to use a specific API or Relay?

No, in fact using your own is highly recommended! While there is a first-party API available to you that streamlines the collection, management and distribution of the incoming plugs related to your protocol, it is not required. You can use any API or Relay that you would like to consume the plugs and distribute them to the appropriate network of Executors.

## Doesn't this just make my protocol more complex?

Quite the opposite! By implementing Plug you are able to remove the complexity of the EVM from your protocol and provide your users with a much more performant experience. This allows you to focus on the core functionality of your protocol and not worry about designing around the limitations of the EVM.

## Does my protocol still support native transactions?

Yes! Plug is designed to be a drop-in solution that allows you to support both imperative and declarative transactions. This allows you to support the best of both worlds and provide your users with the best experience possible.

## Bro... These docs are huge! Can I get a TL;DR?

Sure! Here is a quick summary of the most important parts of the Plug framework:

- **Declarative Transactions** - Plug allows you to break free of the limitations of the EVM and provide your users with a much more performant and secure experience.

- **Plugs** - Plug allows you to build a bundle of Plugs that is as simple or complex as you need it to be. This allows you to build a Plug that is as simple as a single signature or as complex as a multi-signature with a time-lock. The possibilities are endless.

- **Fuses** - Fuses provide the enforcement of a users conditional declaration. From execution time, to allowed methods, to active state of the protocol being interacted with. Simply, Fuses check that something is in the state desired and reverts the transaction otherwise.

That's it! If you are looking for more information, please check out the rest of the documentation.

::: tip

If you are still looking for more information, please [reach out for help](https://twitter.com/nftchance). I would love to help you get started and moving in the right direction!

:::
