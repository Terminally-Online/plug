# Live Pairs

At the very base level `Plug` has been designed to operate on top of [EIP-712](/decoders/eip-712) to enable the signing and verification of structured data.

The framework extends the EIP by introducing the concept of a `LivePair`. With `LivePairs` the onchain framework operates with the implicit assumption that all `Plugs` are delivered in a predefined and standardized data layout.

## What is a Live Pair?

A `LivePair` is a pair of structured data that has been signed by a single EVM address. The first piece of data is the `data` and the second is the `signature`. The `data` is the structured data that has been signed and the `signature` is the signature of the `data` signed by the EVM address.

Let's look at a simple example where we have `Mail` that we want to send from `Alice` to `Bob`. In this example, `Alice` is the `Signer` and `Bob` is the `Recipient`. The first step is to declare the types of our protocol:

::: code-group

```typescript [constants.ts]
export const types = {
  Mail: [
    { name: "from", type: "Person" },
    { name: "to", type: "Person" },
    { name: "contents", type: "string" },
  ],
  Person: [
    { name: "name", type: "string" },
    { name: "wallet", type: "address" },
  ],
  LiveMail: [
    { name: "mail", type: "Mail" },
    { name: "signature", type: "bytes" },
  ],
} as const;
```

:::

Of note here is that the `LiveMail` type is a nested `Mail` type with an additional `signature` field. This is the `LivePair` that we are looking for. This signals:

- `Mail` is the signed message.
- `LiveMail` is the type consumed by the onchain function.
- `signature` is the signature of the `Mail` signed by the `from` field of the `Mail`.

Why is it called a `LivePair`?

- The type MUST be prefixed with `Live`.
- The type MUST be a nested type of the `TypedData` that is being signed.
- The type MUST always have a `signature` field that is of type `bytes`.

## The Resulting Assumption

`LivePairs` enables the onchain protocol of `Plug` to operate with the assumption if an `Plug` is being used or verified, it is always in the form of the `LivePair` of the message being executed. In [generic form](https://www.typescriptlang.org/docs/handbook/2/generics.html), this can be thought of as:

::: code-group

```typescript [LivePair.ts]
export type TypedDataToLivePlug<K, U> = Record<"signature", `0x${string}`> & {
  [TK in K as Lowercase<string & TK>]: U;
};
```

:::

This may seem confusing if you are not familiar with `Typescript` so let's break it down:

- `TypedDataToLivePlug` is a generic type that takes two arguments: `K` and `U`.
  - `K` is the type of the `TypedData` that is being signed.
  - `U` is the type of the `TypedData` that is being consumed by the onchain function.
- The return type is a `Record` that has a `signature` field that is of type `string` and a `TypedData` field that is of type `U` with a lowercase key.

In the case of our `Mail` example, this results in the `LiveMail` type of:

::: code-group

```typescript [LiveMail.ts]
type LiveMail = {
  signature: string;
  mail: {
    from: {
      name: string;
      wallet: string;
    };
    to: {
      name: string;
      wallet: string;
    };
    contents: string;
  };
};
```

:::

Due to the simple architecture in place you can immediately pop over to `Plug` with the `LiveMail` type and start using it onchain with:

::: code-group

```solidity [Types.sol]
function getLiveMailSigner(
    LiveMail calldata $signedMail,
    bytes32 $domainHash
) public view returns (address) {
    return getMailHash($signedMail.mail, $domainHash).recover(
        $signedMail.signature
    );
}
```

:::

With just these few lines of `Solidity` we now have the ability to:

- Securely send and receive Mail.
- Verify that the Mail was sent by the `Signer`.
- Verify that the Mail has not been tampered with.

While this is a simplified example, it is important to understand that this is the foundation of the `Plug` framework. With this simple architecture, the framework can be used to build complex protocols that are secure, modular, and easy to use.
