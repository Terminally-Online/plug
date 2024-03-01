---
head:
  - - meta
    - property: og:title
      content: Addresses
  - - meta
    - name: description
      content: Deploy a piece of the Plug Framework on any chain with the help of CREATE2.
  - - meta
    - property: og:description
      content: Deploy a piece of the Plug Framework on any chain with the help of CREATE2.
---

# Deterministic

All pieces of Plug can be deployed to their respective canonical address on EVM chains using 0age's keyless CREATE2 factory (`0x0000000000FFe8B47B3e2130213B802212439497`). By doing this, the addresses referenced are constant, docs are easier to follow, and users can verify that the contract being interacted with is truly the contract that has been designed.

## Mining an Efficient Address

Even with standard verification practices there is nothing that stops developers from adding malicious code and verifying the code on Etherscan to provide a false-security that the contract is the same across all places.

With this approach, all users and developers can rest assured, trust, and verify themselves that the deployed source is accurate simply by confirming the address. This means that instead of submitting a PR that introduces new overhead or relying on the Plug team to deploy each instance, you can just deploy the instance yourself. Further, this means that integration across any chain can be high-speed as incongruencies between address interaction does not exist. A `Fuse` deployed on one chain should be found at the same address on all others.

Additionally, if you choose to deploy your own instance with the method provided below you do not even need to open a PR. Our system will automatically detect when new deployments take place.

Finally, a major reason to follow this method of deployment is that with the use of CREATE2 one can mine efficient addresses that save users real money. In one-off instances the savings is quite small (~<100 gas) however when you accumulate thousands of calls and transactions this savings become significant rather quickly.

## Keyless CREATE2 Factory

The factory to deploy contracts with CREATE2 is rather simple. The address of the contract being deployed is determined by:

```solidity
address deployment = address(
  uint160(
    uint256(
      keccak256(
        abi.encodePacked(
          hex"ff",
          address(this),
          salt,
          keccak256(
            abi.encodePacked(
              initCode
            )
          )
        )
      )
    )
  )
);
```

Simplified, you can imagine this as:

```solidity
address deployment = hash(0xFF, $sender, $salt, $bytecode);
```

- In this case the sender is, `address(this)`, the address of the factory.
- The salt is user-chosen and a result of your mining efforts.
- The bytecode can be retrieved by building your contract and retrieving the value from the generated artifacts or simply logging it from a foundry script with `type(YourContract).creationCode`.

With just these pieces of data you have everthing needed to deploy an instance of the protocol.

::: info

Note, the address of the CREATE2 factory and all supporting pieces are constant across most (it would be all, but certain chains like EVMOS have broken their compatibility) EVM chains. It is unlikely that support is built for chains that have chosen to break this standard expectation.

:::

### Mining an Address

To mine an address for a contract can be confusing at first, but it's really rather simple:

1. **Acquire GPU access. Whether local or in the cloud does not matter.**

2. **Access your terminal or SSH in and install `Rust` as well as [create2crunch](https://github.com/0age/create2crunch).**

```bash
sudo apt install build-essential -y; \
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs
  | sh -s -- -y; source "$HOME/.cargo/env"; \
git clone https://github.com/0age/create2crunch; \
cd create2crunch; sed -i 's/0x4/0x40/g' src/lib.rs
```

This command contains the reference usage of linux systems. In the case that you are doing so locally it is rather likely that you will need to install a few system level dependencies and potentially even update the drives of your graphics card. If you run into issues, just ask ChatGPT.

3. **Acquire the initcode hash of your contract.**

However you choose to do it is fine. For reference, in Plug we solve for it with:

```typescript
import { keccak256 } from "viem";

const initCode = "abcd...1234";
const initCodeHash = keccak256(`0x${initCode}`);
```

For all contracts of Plug these artifacts are automatically generated when building and can be found in `artifacts/YourContract.sol/YourContract.initcode.json`.

4. **Mine the address within your specification.**

```bash
export FACTORY="0x0000000000ffe8b47b3e2130213b802212439497"; \
export CALLER="0x0000000000000000000000000000000000000000"; \
export INIT_CODE_HASH="<INSERT_YOUR_HASH_HERE>"; \
export LEADING=5; \
export TOTAL=7; \
cargo run --release $FACTORY $CALLER $INIT_CODE_HASH 0 $LEADING $TOTAL
```

Because you are deploying with the keyless CREATE2 factory you will leave `CALLER` as the zero address. Do not use the address of the account you will use to make the call to the factory. For `LEADING` and `TOTAL` you will set these to the numbers that you choose. The higher the number, the longer it will take to find a match.

To determine how long it will take you can apply the simple formula of:

```typescript
hours = 256 ** LEADING / (ITERATIONS_PER_SECOND * 3600);
```

Finally, it is worth noting that hexadecimal characters are not one-to-one with the visual representation. If you set `LEADING` to 5, then it will have 10 leading zeroes.

As the script runs, the outputs will be saved to `efficient_addresses.txt` along with results that closely match your specification being logged to the console. When it reaches a point that you are happy with you can choose the address and salt that is most to your liking.
