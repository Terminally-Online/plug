![Plug solver banner](/plug-solver.png)

## Generating Contract Bindings

For all contract interactions we build an interface binding that streamlines the act of onchain transaction preparation and execution. For this, we utilize the [abigen](https://github.com/ethereum/go-ethereum/tree/master/cmd/abigen) tool from the Ethereum Go client.

While the bindings are generated for abis that are sitting in the `abis` directory, you can also generate them for any contract that is deployed on the Ethereum mainnet. To do this, you will add a mapping record to the relevant network in `references.go`.

With this, all you need to do is add the abi to the `abis` directory and run `make bindings`. By doing this a related `go` file will be generated that you do not need to manage or manually manipulate.

It is implemented this way to allow for the lowest effort and most expedient way of integrating a new protocol.
