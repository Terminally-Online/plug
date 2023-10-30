---
head:
    - - meta
      - property: og:title
        content: Imperative Transactions in Blockchain
    - - meta
      - name: description
        content: Understanding the immediate and deterministic nature of Imperative Transactions in blockchain.
    - - meta
      - property: og:description
        content: Understanding the immediate and deterministic nature of Imperative Transactions in blockchain.
---

# Imperative Transactions

Imperative transactions are the basic actions that make things happen on a blockchain.

Think of them like sending an email; once you hit send, it's gone, and the message is delivered. These are the transactions that most people are familiar with in the blockchain world, like sending or receiving cryptocurrency. They are straightforward but rigid, performing a single action as soon as they're confirmed by the network. You can't change them or add conditions to them once they're sent. They're your straightforward, no-nonsense way to interact with a blockchain.

Now, take a second to think about how many times you have pressed that `Send` button only to scream out '**NoOoOoO!!!**' a few seconds later. That's the state of typical blockchain user experience even when it's just $5 dollars.

## Characteristics of Imperative Transactions

### Immediacy

The first thing to know about imperative transactions is that they happen quickly. Once the network confirms your transaction, it's done. The blockchain updates, and that's it. This quick action is useful for straightforward tasks like transferring money, but it means you can't take it back. Mistakes are permanent, so you need to be careful.

### Determinism

Determinism in this context means that the result of your transaction is predictable. If you're sending someone 10 coins, you know that once the transaction goes through:

-   They'll have 10 more coins,
-   and you'll have 10 less.

This is good for everyone; it keeps things transparent and easy to understand. However, it also means that these types of transactions lack flexibility. What you see is what you get.

### Fees and Gas

Sending transactions isn't free; you have to pay for the computing power that makes it happen. These fees vary depending on the blockchain you're using and how busy it is. The fees serve two purposes.

-   First, they discourage spam transactions that could clog the network.
-   Second, they incentivize the people running the network to include your transaction.

So, you'll want to set your fees carefully.

## Examples of Imperative Transactions

### Value Transfer

The most basic transaction is just sending money from one person to another. It's straightforward, immediate, and irreversible. Whether you're sending funds to a friend or paying for goods and services, value transfers are the simplest form of blockchain transactions.

### Minting Tokens

When you mint new tokens in a blockchain like Ethereum, you're executing an imperative transaction. You call a function on a smart contract to create a specific number of tokens. Once minted, these tokens exist and are usually sent to a particular address.

### Voting in a DAO

DAOs (Decentralized Autonomous Organizations) often require members to vote on proposals. Casting your vote is an imperative transaction. Once you vote 'yes' or 'no' on a proposal, your vote is recorded on the blockchain, and you can't change it later.

### Changing Ownership of a Digital Asset

When you sell or transfer a digital asset like an NFT (Non-Fungible Token), you're executing an imperative transaction. You update the owner field in the smart contract, transferring ownership rights to the new owner. Like all imperative transactions, this is irreversible.

### Creating a Smart Contract

When you deploy a new smart contract to a blockchain, that's an imperative transaction. The code is sent in a transaction, and once confirmed, the smart contract exists on the blockchain at a specific address.

## Why Imperative Transactions are Vital

Imperative transactions are like the bread and butter of blockchain. They're the simplest, most direct way to make something happen. They're the starting point for anyone new to blockchain and will likely remain the most commonly used form of transaction for at least the next few years. However, they aren't perfect. They lack the flexibility and conditional logic that other, more advanced types of transactions offer. That's where things like [Intents](/intents/introduction) come in, adding layers of complexity and options on top of these basic transactions.
