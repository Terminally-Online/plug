---
head:
  - - meta
    - property: og:title
      content: Revocation
  - - meta
    - name: description
      content: Details of permission revocation and why it's important.
  - - meta
    - property: og:description
      content: Details of permission revocation and why it's important.
---

# Revocation

The concept of raw **revocation** is extremely simple: **_If you can give, you can take away._** But, when it comes to transaction blockchains that are immutable, this becomes a bit more complex.

With typical native EVM transactions on Ethereum, there is no way to revoke a transaction. Once it's sent, it's sent. If you change your mind, you can't take it back. If the conditions change and you don't want to experience the consequences of the transaction, you can't take it back.

With `Plug`, you can with ease.

Importantly, this is enabled by precisely the same permission-stack as every other execution condition. Due to the modular design of `Plug`, unlocking this ability for the users of your protocol is extremely simple and only requires app-level logic.

## How does it work?

As revocation is built on the same permission-stack as every other execution condition, it is extremely simple to implement. In fact, it is so simple that it is only a two-step process:

1. Scope a [Revocation Enforcer](/core/enforcers) at the time of **giving** the permissions.
2. Call the `revoke` function on the [Enforcer](/core/enforcers) **originally declared** in the permissions.

::: info

If you have been reading the documentation from top to bottom, you may not have gotten to the [Enforcers](/core/enforcers) section yet. If this is the case, you may want to read that section before continuing. For now, the simple explanation is that an [Enforcer](/core/enforcers) is what powers the "_if this_" part of the "_if this, then that_" logic of a permission.

:::

### Giving Revocable Permissions

The first step is to scope a [Revocation Enforcer](/core/enforcers) at the time of declaring the permission. This is done by adding the `RevocationEnforcer` to the `enforcers` array of the `Permission` struct.

```typescript
const permissions = {
  delegate: "<the EVM address of the delegate>",
  authority: bytes32(0),
  caveats: [
    {
      enforcer: RevocationEnforcer.address,
      terms: bytes(0),
    },
  ],
  salt: bytes32(Date.now().toString()),
};
```

With the `RevocationEnforcer` scoped, the permissions can now be signed and given to the delegate. If the permissions given are ever used, first the `RevocationEnforcer` will be called to ensure that the they have not been revoked.

### Revoking Permissions

With the permissions given, let's look at the implementation of the `revoke` function in the `Plug` framework contract in chunks to better understand how this works:

```solidity
function revoke(SignedPermissions calldata $signedPermissions, bytes32 $domainHash) public
```

Notably, the `revoke` function takes two arguments: the `SignedPermissions` and the `domainHash`.

- The `SignedPermissions` is the same as the `SignedPermissions` that was originally declared.
- The `domainHash` is the same as the `domainHash` of the intent target (you may give the same permissions for two different contracts).

```solidity
require(
    getSignedPermissionsSigner($signedPermissions, $domainHash) == _msgSender(),
    'RevocationEnforcer:InvalidRevoker'
);
```

Now inside the function, the logic starts by ensuring that the `Sender` of the `revoke`transaction is the same as the`Signer`of the permissions. This is important because it ensures that only the original`Signer` of the permissions can revoke them.

```solidity
bytes32 permissionsHash = getSignedPermissionsHash($signedPermissions);
```

Now that we know the caller is the signer of the permissions, we need to determine the hash of the permissions that are being revoked.

```solidity
require(isRevoked[permissionsHash] == false, 'RevocationEnforcer:AlreadyRevoked');
```

Everything is going great, but before we can revoke the permissions, we need to ensure that they have not already been revoked.

```solidity
isRevoked[permissionsHash] = true;
```

Finally, with all that work done, we can revoke the permissions by setting the `isRevoked` mapping to `true`. All in all, not too complicated.

::: tip

It is very important to understand that when you revoke permissions, it does so based on the hash of the permissions given. If the user attempts to use that same set of permissions (that encode to the same hash) again in the future, it will be rejected.

Due to this, it is recommended to use `salt` as an encoded timestamp to ensure that the hash is unique each time a set of permissions are given.

Before moving on it is worth noting that the `SignedPermissions` remain localized to the `domainHash` of the intent target. This means that if you revoke permissions for one `domainHash`, it will not affect the permissions for any other `domainHash`.

:::
