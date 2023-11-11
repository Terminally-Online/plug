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

Getting started with Plug is quite simple however it is understandable that you may have some questions. This page will attempt to answer some of the most common questions that you may have.

- [Why would I want my protocol to support declarative transactions?](#why-would-i-want-declarative-transactions)
- [Can I use Plug with smart contracts that are already deployed?](#can-i-use-plug-with-protocols-already-deployed)
- [If I want basic support do I need to implement custom types?](#if-i-want-basic-support-do-i-need-custom-types)
- [Do I need to deploy my own set of Enforcers?](#do-i-need-to-deploy-my-own-set-of-enforcers)
- [Is there a fee?](#is-there-a-fee)
- [Do I have to use a specific API or Relay?](#do-i-have-to-use-a-specific-api-or-relay)
- [Doesn't this just make my protocol more complex?](#doesn-t-this-just-make-my-protocol-more-complex)
- [Does my protocol still support native transactions?](#does-my-protocol-still-support-native-transactions)
- [Bro... These docs are huge! Can I get a TL;DR?](#bro-these-docs-are-huge-can-i-get-a-tl-dr)

::: info

If you are using the documentation and feel that a question or answer belongs here, please do edit the page and submit a PR. We would love to hear from you! Otherwise, this page is regularly updated with new questions and answers so check back often.

:::

## Why would I want declarative transactions?

Native EVM transactions serve as a constant limiter of the throughput for your protocol. Bound to the execution path of the blockchain it is deployed on your users are forced to not only accept, but wait for the less than ideal performance of the EVM. Declarative transactions allow you to break free of this limitation and provide your users with a much more performant experience to prevent significant fund loss and waste.

## Can I use Plug with protocols already deployed?

While the architecture is designed to support this functionality, a first-party implementation has not yet been made public. Plug attempts to make integration as seamless as possible by avoiding the inclusion of opinion and a `Relay-like` implementation is in the works, however, it is not yet ready for public use.

## If I want basic support do I need custom types?

By default the Plug framework has been packaged with a pre-built set of types and smart contracts that are needed to consume the framework. If you are looking to build a protocol that is not supported by the pre-built types, you will need to implement your own custom types by utilizing [`@nftchance/plug-types`](https://www.npmjs.com/package/@nftchance/plug-types?activeTab=readme). Otherwise, you can use the pre-built types and not worry about the declaration or consuming implementation of types.

## Do I need to deploy my own set of Enforcers?

_Generally, no._ The base set of [Enforcers](/enforcers) have been designed to be consumed by deployed instances of the `Plug` framework. If you are using a third-party implementation of the core framework or an `Enforcer` you will need to read the logic and determine if it is suitable for your use case as there is no expected or enforced standard.

## Is there a fee?

Within the base framework there is no fee declared as you have full access to the benefits of declarative transactions. While there is no fee baked in it is very simple to implement a fee structure designed to fit and benefit the rest of your protocol.

::: tip

If you are not sure how to implement a fee structure, please [reach out for help](https://twitter.com/nftchance). I would love to help you get started and moving in the right direction!

:::

## Do I have to use a specific API or Relay?

No, in fact using your own is highly recommended! While there is a first-party API available to you that streamlines the collection, management and distribution of the incoming intents related to your protocol, it is not required. You can use any API or Relay that you would like to consume the intents and distribute them to the appropriate [Executors](/executors).

## Doesn't this just make my protocol more complex?

Quite the opposite! By implementing Plug you are able to remove the complexity of the EVM from your protocol and provide your users with a much more performant experience. This allows you to focus on the core functionality of your protocol and not worry about designing around the limitations of the EVM.

## Does my protocol still support native transactions?

Yes! Plug is designed to be a drop-in solution that allows you to support both native and declarative transactions. This allows you to support the best of both worlds and provide your users with the best experience possible.

## Bro... These docs are huge! Can I get a TL;DR?

Sure! Here is a quick summary of the most important parts of the Plug framework:

- **Declarative Transactions** - Plug allows you to break free of the limitations of the EVM and provide your users with a much more performant experience.

- **Permission-Stack** - Plug allows you to build a permission-stack that is as simple or complex as you need it to be. This allows you to build a permission that is as simple as a single signature or as complex as a multi-signature with a time-lock.

- **Enforcers** - Plug allows you to build a permission-stack that is as simple or complex as you need it to be. This allows you to build a permission that is as simple as a single signature or as complex as a multi-signature with a time-lock.

That's it! If you are looking for more information, please check out the rest of the documentation.

::: tip

If you are still looking for more information, please [reach out for help](https://twitter.com/nftchance). I would love to help you get started and moving in the right direction!

:::
