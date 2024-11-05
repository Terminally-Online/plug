![Plug solver banner](/plug-solver.png)

This Solver repository functions as the primary backend for onchain transactions and Plug action metadata.

## Prerequisites

- Go 1.23.x or later
- Make
- Environment variables (create a `.env` file in the root directory):
  ```env
  ALCHEMY_API_KEY=your_alchemy_api_key
  ```

## Getting Started

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod download
   ```
3. Create your `.env` file with the required environment variables

## Running Processes

To run any piece of the app you will run `make <command>` where `<command>` is the name of the command you want to run:

```ml
commands
├─ api — "Run the API server that enables endpoint access."
├─ bindings — "Generate Go bindings for integrated contracts."
├─ cron — "Run the regularly scheduled maintenance and simulation jobs."
└─ references — "Retrieve contract ABIs from the block explorer for bindings."
```

## Cron Jobs

The Plug stack relies on several cron jobs to run periodically that keep the system running smoothly and as responsive as possible. They are not detailed in depth as you should not need to make any adjustments, though you can find them in [cron/jobs.go](cron/jobs.go).

A majority of these jobs are scheduled to keep the app running so you should always have the app running if you want to run the solver otherwise there will be no data that the solver can retrieve.

## API Endpoints

Contained within this repository are the following API endpoints:

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
