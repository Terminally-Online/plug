---
head:
    - - meta
      - property: og:title
        content: Declarative Messages
    - - meta
      - name: description
        content: A brief breakdown of why intents are important.
    - - meta
      - property: og:description
        content: A brief breakdown of why intents are important.
---

# Declarative Messages

Unlike [Imperative Transactions](/intents/imperative-transactions), `Declarative Messages` don't perform an action directly. Instead, they express an intent to perform an action under specific conditions.

Think of it like placing an order for a customized car. You aren't buying the car right then and there; you're specifying what you want (color, features, etc.), and the action (you receiving the car) happens later, once the car is built to your specifications.

## Characteristics of Declarative Messages

### Flexibility

Declarative messages offer a level of flexibility that is unmatched by their imperative counterparts. You can specify multiple conditions under which an action should be taken.

For instance, in a multi-signature wallet, a transaction could be executed only if a certain number of authorized signatures are gathered.

More interestingly, imagine that an employee is paid in an ERC20 that fluctuates in price, but compensation in USD denomenation. While they earn $10 USD they may earn 5000 $TOKEN. A declarative message could be used to automatically convert the defined USD amount to the ERC20 aprice at the time of payment without the employee having to do anything.

### Conditional Execution

The hallmark of a declarative message is that it sets conditions for execution. Think of it like an "_If, Then_" statement. If all specified conditions are met, then the transaction will execute. Of course, if a bad actor was to run the declared transaction before the conditions were met, the transaction would fail because the '_If_' declared has not been satisfied.

### Compositionality

Declarative messages can be composed together to form more complex transaction templates. In the blockchain ecosystem this is commonly referred to as transaction 'batching' or 'bundling'.

For instance, a single declarative message could involve a token swap, followed by staking the received tokens, but only if the initial token's price reaches a certain level.

::: tip

There is another common implementation of batch settlement known as 'chaining' where responses from previous functions are used to inform following ones however Emporium does not currently support this.

:::

### Predictability

Despite their complexity, declarative messages are designed to be predictable. Each condition within the message is verifiable, ensuring that participants can anticipate the outcome if all conditions are met.

With declarative messages, you can be confident that the transaction will execute as intended, even if it's a complex one.

If something can change to negatively impact the outcomes of your transaction, conditions can be set to protect against it.

## Examples of Declarative Messages

### Conditional Payments

Imagine you want to pay someone only if they complete a specific task by a certain date. You could create a declarative message that releases the funds only if proof of the completed task is submitted and verified on the blockchain by that date.

### Automated Trading

In decentralized finance (DeFi), traders often set conditions for buying or selling assets. For instance, you could specify that you want to sell a token if its price reaches a certain high, and buy it back if it drops to a specific low.

### Escrow Services

Using declarative messages, you can create a smart contract that serves as an escrow. The contract could release funds to a seller only after verifying that the buyer has received the shipped goods.

### Governance Proposals

In a DAO, you could use declarative messages to propose changes that only go into effect if they receive enough votes and meet any other specified conditions.

## Why Declarative Messages are Vital

Declarative messages serve as the building blocks for more complex and conditional transactions, enabling functionalities that go beyond the capabilities of basic, imperative transactions. This added layer of flexibility is critical for the development of decentralized applications (dApps) and smart contracts that require a broad range of user interactions and system states.

For instance, they enable the creation of decentralized autonomous organizations (DAOs), conditional payments in DeFi, and much more. This makes declarative messages an essential tool in achieving the full potential of what blockchain technology promises: a more decentralized, transparent, and programmable world of transactions and agreements.

In essence, while imperative transactions are akin to individual brush strokes, declarative messages are the palettes and techniques that allow for the creation of a full-fledged painting. They add depth, context, and sophistication to the blockchain canvas, allowing for a more intricate and detailed representation of agreements and actions.
