---
head:
  - - meta
    - property: og:title
      content: Vaults
  - - meta
    - name: description
      content: Vaults on Plug are super charged smart contract accounts that power blockchain transactions with conditional execution.
  - - meta
    - property: og:description
      content: Vaults on Plug are super charged smart contract accounts that power blockchain transactions with conditional execution.
---

# Vaults

`Vaults` are the key interaction point for end-users of [Plug](/). Building on top of the concept that many are already familiar with, `Vaults` are a smart contract that operates much like a typical account.

The divergence from familiarity comes from the fact that users are empowered with the ability to append additional logic, set required conditions for all transactions, and schedule automated execution with ease.

Additionally, [Plug](/) holds the ability to wrap an existing smart contract account on top of the `Vault`. So, both [Externally Owned Accounts](https://ethereum.org/en/developers/docs/accounts) (EOAs) and existing Smart Contract Accounts ([Safe](https://safe.global/) & [4337](https://eips.ethereum.org/EIPS/eip-4337) accounts) are immediately supported. Nothing special has to be done. Simply, there are no limitations that keep any blockchain user or protocol from unlocking the power of generalized intents. All configuration and setup is automatically handled by the protocol and accompanying application used to interface with the onchain protocol.

Notably, a user of Plug can even bypass the need of deploying a Vault with the one significant caveat that no transactions with associated value can be run. To run a transaction that transfers an asset or pays the native asset such as ETH, a Vault must be created first.

## Upgradeable Beacons

An important piece of Vaults deployed by Plug is that they are upgradeable to the newest version, but only at the request of the user. This means, that the Plug team can continue building without negatively impacting Vaults that are already deployed. You do not need to create and migrate to a new Vault every time there is an update to the capabilities of the system. Instead, users have the ability to opt in and choose to upgrade to the newest version.

At no time, can Plug or any of our team members unilaterally upgrade the state of your Vault. Upgrades always happen with the direct permission and expression of the user(s) that own and control the Vault.
