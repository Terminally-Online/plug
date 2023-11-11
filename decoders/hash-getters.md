# Hash Getters

The term [Hash Getters](/decoders/hash-getters) might sound a bit arcane, but it's a concept vital to the `Plug` framework.

This page aims to demystify hashGetters, delve into the importance of EIP-712 type hashes, and provide you with the knowledge required to understand the Hash Getters feature in Plug.

## What Are Hashes?

Let's start with the basics. A hash is a unique representation of data that is generated through a mathematical function.

Let's look at a very simple hashing function:

```typescript
function hash(data: string): string {
  return data
    .split("")
    .map((char) => char.charCodeAt(0))
    .join(".");
}
```

If we call the `hash` function in this example with the string `hello`, we will get the following result:

```typescript
// * We can now have a hash reference for the string 'hello'
hash("hello") == "104.101.108.108.111";
// * Hello and goodbye have different hashes
hash("goodbye") == "103.111.111.100.98.121.101";
// * The hash function returns the same hash for the same data.
hash("hello") == "104.101.108.108.111";
```

This is a very simple example, but it illustrates the basic concept of a hash. The function takes in a string and returns a unique representation of that string.

In the realm of blockchain and cryptographic technologies, hashes serve to protect data integrity. **In simpler terms, if we hash different data we will get different hashes.** If we hash the same data, we will always get the same hash. This is the basis of the `Plug` framework.

::: tip

When building a dApp or API that consumes typed data, you should store intent declarations by their reference hash while only storing one instance of the decoded data itself.

This way, instead of storing a massive amount of payloads in a database or in memory, it's just a matter of storing a reference to the hash and lazy-loading the data when needed.

:::

## EIP-712 Type Hashes

When we talk about [EIP-712](/decoders/eip-712), we are referring to [Ethereum Improvement Proposal 712](https://eips.ethereum.org/EIPS/eip-712), a standard that specifies a methodology for creating typed, structured data hashes.

A type hash extends the idea of the basic hashes we just discussed by introducing the ability to recover the original data from the hash. This is done by using a hashing function that can be both encoded and decoded.

The EIP-712 standard is crucial because it provides a clear, user-friendly way to understand and confirm what you are signing or authorizing. With a simple architecture EVM developers and users gain access to:

- **Human-Readable Information:** The structured data in EIP-712 is both machine and human-readable, offering a transparent representation of the transaction.

- **Secure & Verifiable:** It helps to create a unique and tamper-proof signature, ensuring secure transactions.

- **Standardization:** By adhering to a common standard, it ensures compatibility and interoperability among various Ethereum-based platforms and applications.

## Onchain Getters

Hash Getters are specialized functions within the `Plug` framework that facilitate the retrieval of message hashes.

What many don't fully comprehend at first, is that when you are working with a blockchain signtaure you have to be able to verify any of the data that is provided in the message.

This means, that if we have a signature that says 'Spongebob aired at 4pm yesterday,' but have no way to prove it onchain, then the signature is useless. In effect, it represents data that could be false even though it's been signed.

Hash Getters solve this by enabling smart contracts to look at the message and signature provided and confirm that everything can be verified. With a unique hash for each message and type of data, smart contracts can verify the integrity of the data provided in the message in just a couple lines of `Solidity`.
