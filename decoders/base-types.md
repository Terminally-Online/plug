# Base Types

`Emporium` is architected on a handful of [Base EIP-712 Types](/decoders/base-types) that are used to build the [Core Framework](/core/framework) and [Enforcers](/core/enforcers). Alone these types are not very useful, but together they form the foundation of the `Emporium` framework.

## Type Hashes

Type hashes are a cornerstone in the EIP-712 standard and by extension, the Emporium framework. They function as unique identifiers for each type of data you're working with. These are not to be confused with instance hashes, which are identifiers for specific instances of data. Type hashes essentially serve as templates or blueprints.

Imagine building a house without a blueprint. You could try to explain what you want, but without a standardized plan, you're bound to run into inconsistencies, misunderstandings, or mistakes. Type hashes act like these blueprints, giving a standardized representation that all parties can understand and agree upon.

For instance, when someone signs a permission, it's not just the data that is signed, but also its type hash. This ensures that the `Signer`, the `Executor`, and the smart contract all agree on the "shape" or structure of the data.

### How to Calculate a Type Hash

The type hash is calculated by taking the `keccak256` hash of the type concatenated with the type hash of each of its children. For example, the type hash of the `EIP712Domain` type can be calculated like:

::: code-group

```typescript [viem.ts]
import { keccak256 } from 'viem'

const TYPE_HASH = keccak256(
	toHex(
		'EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)'
	)
)
```

```typescript [ethers.ts]
import { ethers, TypedDataEncoder } from 'ethers'

import config from './config'

const encoder = new TypedDataEncoder(config.types)

const TYPE_HASH = ethers.keccak256(
	ethers.toUtf8Bytes(encoder.encodeType(typeName))
)
```

```solidity [Verbose.sol]
bytes32 constant TYPE_HASH = keccak256(
    abi.encodePacked(
        "EIP712Domain(",
        "string name,",
        "string version,",
        "uint256 chainId,",
        "address verifyingContract",
        ")"
    )
);
```

```solidity [Inline.sol]
bytes32 constant TYPE_HASH = keccak256(
    'EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)'
');
```

```solidity [Hash.sol]
bytes32 constant TYPE_HASH = 0x0aeb9481a395eb23cf4f23582fc3486e5f39ec614b0fa703eff30fe32245d399
```

:::

## Type Categories

::: tip

If you are using a modern interface library for Ethereum such as `viem`, you will not need to worry about the categories. The library will handle the decoding and remove unused types. In other libraries though, you may need to manually remove unused types. For example, if you are using `ethers` or `web3.js` you will need to manually remove unused types.

:::

### Domain Types

As covered in the [EIP-712](/decoders/eip-712) section, typed signatures include the `domain hash` so that a `Signer` can always be sure the data being signed is not malicious or meant for a different contract.

-   [EIP712Domain](/generated/base-types/EIP712Domain)

By default, most libraries include the `EIP712Domain` type in the signature. This is because the `EIP712Domain` type is used to define every signature and things would not be secure without it. While the types of every protocol vary, they all share the same `EIP712Domain` type.

## Permission Types

Now stepping into the pieces introduced by the `Emporium` framework, the `Permission Types` are used to define the [Permission](/permissions/introduction) that is being given.

When creating new permissions you will utilize:

-   [Caveat](/generated/base-types/Caveat)
-   [Permission](/generated/base-types/Permission)

When distributing, verifying or executing permissions you will utilize:

-   [SignedPermission](/generated/base-types/SignedPermission)

No matter the complexity of the permission, the `Permission Types` are always the same. This is because the `Permission Types` are used to define the `Permission` that is being given, not the logic of the `Permission`.

## Intent Types

The `Intent Types` are used to define the [Intent](/intents/introduction) that is executed.

When creating new intents you will utilize:

-   [Replay Protection](/decoders/base-types/replay-protection)
-   [Transaction](/generated/base-types/Transaction)
-   [Intent](/generated/base-types/Intent)

When distributing new intents you will utilize:

-   [SignedIntent](/generated/base-types/SignedIntent)
