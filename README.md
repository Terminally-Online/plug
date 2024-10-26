![Plug solver banner](/plug-solver.png)

This Solver repository functions as the primary backend for onchain transactions and Plug action metadata.

Contained within this repository are the following endpoints:

```ml
solver
├─ intent — "Primary endpoint for building intents & transactions."
│   ├─ GET — "Schema definition of the associated action."
│   └─ POST — "Transaction definition of a bundle of actions."
└─ paymaster — "Convert gas costs to comparable ERC20 balances and values."
```

Powering this a lot of logic is going on behind the scenes. Adjusting any piece of this code is not recommended at this time.

## Generating Contract Bindings

For all contract interactions we build an interface binding that streamlines the act of onchain transaction preparation and execution. For this, we utilize the [abigen](https://github.com/ethereum/go-ethereum/tree/master/cmd/abigen) tool from the Ethereum Go client.

While the bindings are generated for abis that are sitting in the `abis` directory, you can also generate them for any contract that is deployed on the Ethereum mainnet. To do this, you will add a mapping record to the relevant network in `references.go`.

With this, all you need to do is add the abi to the `abis` directory and run `make bindings`. By doing this a related `go` file will be generated that you do not need to manage or manually manipulate.

It is implemented this way to allow for the lowest effort and most expedient way of integrating a new protocol.
