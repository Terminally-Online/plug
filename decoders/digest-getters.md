# Digest Getters

In Ethereum Virtual Machine (EVM) frameworks, a crucial aspect of ensuring the security and integrity of signature validations involves what is known as the "digest." This component plays a pivotal role by appending the domain of declaration to a raw signed message, thereby enhancing the uniqueness and security of the transaction.

## Understanding Digests

Following the insights from the [Hash Getters](/decoders/hash-getters) section, it's clear that during the signature process, the contents of a message are hashed to generate a unique representation. However, this approach encounters a significant challenge: if only the content hash is relied upon, the resulting hash remains identical across different contracts. This similarity poses a risk, as it could allow protocols to maliciously or inadvertently accept transactions intended for another contract.

### The Role of Domain in Digests

The introduction of the [domain](/generated/base-types/EIP712Domain) component addresses this vulnerability by adding an additional layer of specificity and security to the hashing process. The [domain](/generated/base-types/EIP712Domain) — which includes details such as the contract's `name`, `version`, `chainId` and the `address` it's deployed to — is also hashed alongside the message content.

This dual-hash mechanism ensures that the digest encapsulates both the unique content of the message and its specific operational context.

To do this, the hash is created with the onchain code living inside the [Plug](/) onchain protocol like:

```solidity
function getPlugsDigest(
  PlugTypesLib.Plugs memory $input
) public view virtual returns (
  bytes32 $digest
) {
  $digest = keccak256(
    bytes.concat("\x19\x01", domainHash, getPlugsHash($input))
  );
}
```

With this simple piece of code, we can make an offchain call to the contract and verify that everything is safe and functions as expected. Notably, the same onchain code is run when an intent is simulated and executed removing the need for an implementation onchain and offchain. Everything runs smoothly off this single point of consumption.

## Onchain Getters

Onchain getters play a crucial role in the practical application of digest getters within smart contracts. These functions allow contracts to retrieve and utilize digests dynamically, facilitating secure and context-specific transactions. By leveraging onchain getters, smart contracts can implement robust mechanisms for signature validation, ensuring that each action taken is both authenticated and explicitly intended for the contract in question.

The integration of digest getters into the smart contract ecosystem represents a significant advancement in the security and integrity of blockchain transactions. By ensuring that each transaction is uniquely tied to its domain, developers can build more reliable, secure, and user-trusted applications on the blockchain.
