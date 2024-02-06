---
head:
  - - meta
    - property: og:title
      content: Transaction Types
  - - meta
    - name: description
      content: A brief breakdown of why plugs are important.
  - - meta
    - property: og:description
      content: A brief breakdown of why plugs are important.
---

# Transaction Types

For the longest time interaction with the blockchain has been limited to basic actions without conditions. A user clicks a button or two and a transaction is submit and executed in response. With the introduction of [Plug](/plugs/introduction) there is a new paradigm where the user gains newfound control to append conditions that must be met in order to execute the transaction.

[Imperative transactions](#imperative-transactions) are the bread and butter of blockchain. They're the simplest, most direct way to make something happen. They're the starting point for anyone new to blockchain and will likely remain the most commonly used form of transaction for at least the next few years. However, they aren't perfect.

They lack the flexibility and conditional logic resulting in less than ideal outcomes very often. That's where [Plug](/plugs/introduction) comes in. By adding layers of conditions on top of these basic transactions users are capable of securing better results with lower risk a higher percent of the time.

We call a transaction with the appendage of conditions [declarative transactions](#declarative-transactions). For instance, they enable the automation of any transaction, conditional payments in DeFi, and much more. This higher level of control is an essential tool in achieving the full potential of what blockchain technology promises: **a more decentralized, transparent, and programmable world of transactions and agreements**.

## Imperative Transactions

[Imperative transactions](#imperative-transactions) are the basic actions that make things happen on a blockchain.

Think of them like sending an email; once you hit send, it's gone, and the message is delivered. These are the transactions that most people are familiar with in the blockchain world, like sending or receiving cryptocurrency. They are straightforward but rigid, performing a single action as soon as they're confirmed by the network. You can't change them or add conditions to them once they're sent. They're your simple, no-nonsense way to interact with a blockchain.

Now, take a second to think about how many times you have pressed that `Send` button only to scream out '**NoOoOoO!!!**' a few seconds later. That's the state of typical blockchain user experience even when it's just $5 dollars.

### Immediacy

The first thing to know about [imperative transactions](#imperative-transactions) is that they happen quickly. Once the network accepts your transaction, there is no changing it. The blockchain updates, and that's it. This speed is useful and valued for straightforward tasks like transferring money, but it means you can't take it back. Mistakes are permanent, so you need to be careful.

### Determinism

Determinism in this context means that the result of your transaction is predictable. If you're sending someone 10 coins, you know that once the transaction goes through:

- They'll have 10 more coins,
- and you'll have 10 less.

This is good for everyone; it keeps things transparent and easy to understand. However, it also means that these types of transactions lack flexibility. What you see is what you get.

### Fees and Gas

Sending transactions isn't free; you have to pay for the computing power that makes it happen. These fees vary depending on the blockchain you're using and how busy it is. The fees serve two purposes.

- First, they discourage spam transactions that could clog the network.
- Second, they incentivize the people running the network to include your transaction.

So, you'll want to set your fees carefully or you could end up wasting large amounts of money in seconds.

## Declarative Transactions

Unlike [imperative transactions](#imperative-transactions), [declarative transactions](#declarative-transactions) don't perform an action directly or immediately. Instead, they express an **intent** to perform an action under specific conditions.

Think of it like placing an order for a customized car. You aren't buying the car right then and there; you're specifying what you want (color, features, etc.), and the action (you receiving the car) happens later, once the car is built and delivered to your specifications.

### Flexibility

[Declarative transactions](#declarative-transactions) offer a level of flexibility that is unmatched by their imperative counterparts. You can specify multiple conditions under which an action should be taken.

For instance, in a multi-sig wallet, a transaction could be executed only if a certain number of authorized signatures are gathered. More interestingly, imagine that an employee is paid in an ERC20 that fluctuates in price, but compensation in USD denomenation. While they earn $10 USD they may earn 5000 $TOKEN.

A [declarative transaction](#declarative-transactions) may be used to automatically convert the defined USD amount to the ERC20 token at the time of payment without the employee having to do anything.

### Conditional Execution

The hallmark of a [declarative transaction](#declarative-transactions) is that it sets conditions for execution. Think of it like an "_If This, Then That_" statement.

If all specified conditions are met, then the transaction will execute. Otherwise, when someone attempts running the declared transaction before the conditions were met, the transaction would fail because the '_If_' added has not been satisfied.

### Composability

[Declarative transactions](#declarative-transactions) can be composed together to form more complex transaction templates. In the blockchain ecosystem this is commonly referred to as transaction bundling.

For instance, this could be seen in practice with a token swap, followed by staking the received tokens, but only if the initial token's price reaches a certain level.

### Predictability

Despite their complexity, [declarative transactions](#declarative-transactions) are designed to be predictable. Each condition within the message is verifiable, ensuring that participants can anticipate the outcome when all conditions are met.

With [declarative transactions](#declarative-transactions), you can be confident that the transaction will execute as intended, even if it's a complex one.

If something can change to negatively impact the outcomes of your transaction, conditions can be set to protect against it.
