# Signed Pairs

At the very base level `Emporium` has been designed to operate on top of [EIP-712](/decoders/eip-712) to enable the signing and verification of structured data.

The framework extends the EIP by introducing the concept of a `SignedPair`. With `SignedPairs` the onchain framework operates with the implicit assumption that all `Intents` are delivered in a predefined and standardized data layout.

## What is a Signed Pair?

A `SignedPair` is a pair of structured data that has been signed by a single EVM address. The first piece of data is the `data` and the second is the `signature`. The `data` is the structured data that has been signed and the `signature` is the signature of the `data` signed by the EVM address.

Let's look at a simple example where we have `Mail` that we want to send from `Alice` to `Bob`. In this example, `Alice` is the `Signer` and `Bob` is the `Recipient`. The first step is to declare the types of our protocol:

::: code-group

```typescript [constants.ts]
const TYPES = {
	Mail: [
		{ name: 'from', type: 'Person' },
		{ name: 'to', type: 'Person' },
		{ name: 'contents', type: 'string' }
	],
	Person: [
		{ name: 'name', type: 'string' },
		{ name: 'wallet', type: 'address' }
	],
	SignedMail: [
		{ name: 'mail', type: 'Mail' },
		{ name: 'signature', type: 'bytes' }
	]
} as const
```

:::

With our types declared we can see that we have:

-   a `Mail` type that has a `from` and `to` field that are both of type `Person` and a `contents` field that is of type `string`.
-   The `Person` type has a `name` field that is of type `string` and a `wallet` field that is of type `address`.
-   Finally, we have a `SignedMail` type that has a `mail` field that is of type `Mail` and a `signature` field that is of type `bytes`.

Of note here is that the `SignedMail` type is a nested `Mail` type with an additional `signature` field. This is the `SignedPair` that we are looking for. This signals:

-   `Mail` is the signed message.
-   `SignedMail` is the type consumed by the onchain function.
-   `signature` is the signature of the `Mail` signed by the `from` field of the `Mail`.

Why is it called a `SignedPair`?

-   The type MUST be prefixed with `Signed`.
-   The type MUST be a nested type of the `TypedData` that is being signed.
-   The type MUST always have a `signature` field that is of type `bytes`.

## The Resulting Assumption

`SignedPairs` enables the onchain protocol of `Emporium` to operate with the assumption if an `Intent` is being used or verified, it is always in the form of the `SignedPair` of the message being executed. In [generic form](https://www.typescriptlang.org/docs/handbook/2/generics.html), this can be thought of as:

::: code-group

```typescript [SignedPair.ts]
export type TypedDataToSignedIntent<K, U> = Record<
	'signature',
	`0x${string}`
> & {
	[TK in K as Lowercase<string & TK>]: U
}
```

:::

This may seem confusing if you are not familiar with `Typescript` so let's break it down:

-   `TypedDataToSignedIntent` is a generic type that takes two arguments: `K` and `U`.
    -   `K` is the type of the `TypedData` that is being signed.
    -   `U` is the type of the `TypedData` that is being consumed by the onchain function.
-   The return type is a `Record` that has a `signature` field that is of type `string` and a `TypedData` field that is of type `U` with a lowercase key.

In the case of our `Mail` example, this results in the `SignedMail` type of:

::: code-group

```typescript [SignedMail.ts]
type SignedMail = {
	signature: string
	mail: {
		from: {
			name: string
			wallet: string
		}
		to: {
			name: string
			wallet: string
		}
		contents: string
	}
}
```

:::

Due to the simple architecture in place you can immediately pop over to `Emporium` with the `SignedMail` type and start using it onchain to verified the signed `Mail` contents of your ecosystem.
