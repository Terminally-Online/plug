---
head:
  - - meta
    - property: og:title
      content: Revocation
  - - meta
    - name: description
      content: Details of pin revocation and why it's important.
  - - meta
    - property: og:description
      content: Details of pin revocation and why it's important.
---

# Revocation

The concept of raw **revocation** is extremely simple: **_If you can give, you can take away._** But, when it comes to transaction blockchains that are immutable, this becomes a bit more complex.

With typical native EVM transactions on Ethereum, there is no way to revoke a transaction. Once it's sent, it's sent. If you change your mind, you can't take it back. If the conditions change and you don't want to experience the consequences of the transaction, you can't take it back.

With `Plug`, you can with ease.

Importantly, this is enabled by precisely the same pin-stack as every other execution condition. Due to the modular design of `Plug`, unlocking this ability for the users of your protocol is extremely simple and only requires app-level logic.

## How does it work?

As revocation is built on the same pin-stack as every other execution condition, it is extremely simple to implement. In fact, it is so simple that it is only a two-step process:

1. Scope a [Revocation Enforcer](/core/fuses) at the time of **giving** the pins.
2. Call the `revoke` function on the [Enforcer](/core/fuses) **originally declared** in the pins.

::: info

If you have been reading the documentation from top to bottom, you may not have gotten to the [Enforcers](/core/fuses) section yet. If this is the case, you may want to read that section before continuing. For now, the simple explanation is that an [Enforcer](/core/fuses) is what powers the "_if this_" part of the "_if this, then that_" logic of a pin.

:::

### Giving Revocable Pins

The first step is to scope a [Revocation Enforcer](/core/fuses) at the time of declaring the pin. This is done by adding the `RevocationEnforcer` to the `enforcers` array of the `Pin` struct.

```typescript
const pins = {
  delegate: "<the EVM address of the delegate>",
  authority: bytes32(0),
  fuses: [
    {
      enforcer: RevocationEnforcer.address,
      terms: bytes(0),
    },
  ],
  salt: bytes32(Date.now().toString()),
};
```

With the `RevocationEnforcer` scoped, the pins can now be signed and given to the delegate. If the pins given are ever used, first the `RevocationEnforcer` will be called to ensure that the they have not been revoked.

### Revoking Pins

With the pins given, let's look at the implementation of the `revoke` function in the `Plug` framework contract in chunks to better understand how this works:

```solidity
function revoke(LivePins calldata $signedPins, bytes32 $domainHash) public
```

Notably, the `revoke` function takes two arguments: the `LivePins` and the `domainHash`.

- The `LivePins` is the same as the `LivePins` that was originally declared.
- The `domainHash` is the same as the `domainHash` of the intent target (you may give the same pins for two different contracts).

```solidity
require(
    getLivePinsSigner($signedPins, $domainHash) == _msgSender(),
    'RevocationEnforcer:InvalidRevoker'
);
```

Now inside the function, the logic starts by ensuring that the `Sender` of the `revoke`transaction is the same as the`Signer`of the pins. This is important because it ensures that only the original`Signer` of the pins can revoke them.

```solidity
bytes32 pinsHash = getLivePinsHash($signedPins);
```

Now that we know the caller is the signer of the pins, we need to determine the hash of the pins that are being revoked.

```solidity
require(isRevoked[pinsHash] == false, 'RevocationEnforcer:AlreadyRevoked');
```

Everything is going great, but before we can revoke the pins, we need to ensure that they have not already been revoked.

```solidity
isRevoked[pinsHash] = true;
```

Finally, with all that work done, we can revoke the pins by setting the `isRevoked` mapping to `true`. All in all, not too complicated.

::: tip

It is very important to understand that when you revoke pins, it does so based on the hash of the pins given. If the user attempts to use that same set of pins (that encode to the same hash) again in the future, it will be rejected.

Due to this, it is recommended to use `salt` as an encoded timestamp to ensure that the hash is unique each time a set of pins are given.

Before moving on it is worth noting that the `LivePins` remain localized to the `domainHash` of the intent target. This means that if you revoke pins for one `domainHash`, it will not affect the pins for any other `domainHash`.

:::
