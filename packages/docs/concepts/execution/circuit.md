# Circuit

In partnership with [Othentic](https://www.othentic.xyz/) and [EigenLayer](https://www.eigenlayer.xyz/) we have designed Circuit, the AVS powering Plug execution with the key principles of:

- **Event-Driven Architecture:** To respond as quickly as possible to market changes and execute precisely on [schedule](../execution/schedules.md) our system is constantly simulating the state of intents and looking for a valid route.
- **Collaborative Solving:** With the streaming of upcoming executions the Plug Solver works to treat [each action](../concepts/actions.md) as an atomic unit to enable for the most efficient execution of [declarative intents](../concepts/actions.html#declarative) like swaps.
- **Efficient Leader Elections:** Until now most existing intent orderflow has been swap and bridge based making it exceptionally capital dependent. With the introduction of non-swap actions it is more important than ever to have a robust election process for intent solving and execution rights.
- **Massively Redundant Execution Trees:** The worst thing that could happen is having a valid route, but never using it onchain. With a dense pool of Operators everything is in place to ensure your intent is successfully fulfilled.
- **Onchain Validation:** Upon successful onchain execution of a Plug, the Attesters will verify the results onchain and come to shared consensus on the validity of the proof provided by the actively elected Performer.

## Operators

While the mechanisms in place for an AVS are exceptionally important, the single most important thing is the individuals participating in the ecosystem as Operators. Within Circuit there are three key Operator roles:

- 👨‍🌾 [Performer](#performer): Doing work.
- 🕵️ [Attester](#attester): Verifying work done.
- 🧑‍✈️ [Aggregator](#aggregator): Orchestrating work to do and done.

Each fill a vital role, have different requirements, and keep the system running in good health by contributing their own piece.

### Performer

Performer is an AVS Operator that executes a task, provides a Proof of Task, and sends the results to [Attesters](#attester). After successfully executing a task, the Performer publishes an event via peer-to-peer networking for [Attester](#attester) nodes to discover.

The RPC call that the Performer sends to the [Attestor](#attester) Nodes:

```json
{
    "jsonrpc": "2.0",
    "method": "sendTask",
    "params": [<proofOfTask>, <data>, <taskDefinitionId>,  <performerAddress>, <signature>]
}
```

> [!TIP]
> Without Performers you would have to pay for the gas consumed when running your transactions. Instead, by contributing to the ecosystem elsewhere Performers have their own incentive models that allow for near costless blockchain use for users like you.

### Attester

Attesters are AVS Operators' quorum that attests to the validity of the executed task. Each task must be attested as either "valid" or "invalid".

The Operator's voting power is proportional and calculated against the amount of re-stake assets staked on the shared security layer, referred to as “dynamic voting power.” The re-staked effective balance determines each Operator's influence in the consensus process.

If over ⅔ of the quorum's voting power attest "valid", the task is considered approved. If over ⅓ of the quorum's voting power attest "invalid", the task is rejected, and the quorum executes a slashing event to the Performer. The Attesters run the validation logic using a local HTTP request to the AVS WebAPI.

```bash
curl -X POST \
 http://localhost:4002/validate_task \
 -H 'Content-Type: application/json' \
 -d '{
    "proofOfTask": "<transactionHash>",
    "data": "<plugsHash>",
    "taskDefinitionID": "<taskDefinitionId>"
 }'
```

> [!TIP]
> Without Attesters we would not be able to safely use off-chain data as triggers and transaction data due to centralization and inaccurate data response risks. So, when you see an Attester out in public make sure to thank them for their service.

### Aggregator

The Aggregator listens to events from the [Attester](#attester) nodes and monitors the necessary voting power contribution to a certain task. The Aggregator aggregates the signatures of the [Attesters](#attester) into a BLS aggregated signature and submits a transaction to the AttestationCenter smart contract. After successful validation, the [Performer](#performer), [Attesters](#attester), and Aggregator are eligible to claim task rewards.

## Tasks

Understanding who [Operators](./operators.md) are is an important first step. Now, we must cover what work is actually being performed. In Circuit, all atomic units of work are broken down into measurable Tasks.

- Plug execution, a task.
- Onchain state validation, a task.
- Stake weight syncing, a task.
- Operator maintenace, a task.

All work needed from the AVS is streamed to the [Aggregator](./operators/#aggregator) as a Task that is then routed to the appropriate parties ([Performers](./operators/#performer) and/or [Attesters](./operators#attester)) with the needed data.

### Leader Election

In Circuit, a leader is chosen from a distributed pool of Operators to receive execution and validation rights.

- **Today:** Plug operates the primary [Performer](#performer) while allowing open entry into the [Attester](#attester) functions. This means that Plug will handle making sure all the transactions are executed and Attesters will constantly verify that data used to run the transactions and that the transaction successfully executed onchain as expected.
- **Tomorrow:** As our Attester pool grows we aim to enable the ability for others to join as a Performer for significant redundancy, reasonable censorship resistance, and the quickest settlement times possible.

## Enrollment

When joining Circuit as an Operator you can choose to utilize the CLI tool we built to make the process easier for you. Alternatively, you can take care of doing everything yourself manually. In the end, the outcome is the same no matter which method you prefer beyond missing convenience features when done manually.

### Managed CLI Approach

1. Install the latest version of the CLI:

::: code-group

```bash [pnpm]
pnpm i -g @terminallyonline/plug
```

```bash [npm]
npm i -g @terminallyonline/plug
```

```bash [yarn]
yarn i -g @terminallyonline/plug
```

:::

2. Initialize yourself as a new Operator in the network:

::: code-group

```bash [pnpm]
pnpm plug register
```

```bash [npm]
npm plug register
```

```bash [yarn]
yarn plug register
```

:::

While registering you will be prompted to provide data for several fields including:

- `PRIVATE_KEY`
  - The private key to your Operator address (public key).
- `REWARDS_ADDRESS`
  - The address rewards are sent upon task completion.
- `NAME`, `DESCRIPTION`, `WEBSITE`, `LOGO_URL`, `TWITTER_URL`:
  - Metadata for the directory.

1. Run the Operator service you would like to participate in.

::: code-group

```bash [pnpm]
pnpm plug run --type=attest
```

```bash [npm]
npm plug run --type=attest
```

```bash [yarn]
yarn plug run --type=attest
```

:::

### Manual Github Approach

> [!CAUTION]
> Auto-updates are only managed by Plug when created through the CLI. Please be cognizant that there may be updates you need to react to continue functional participation as an Operator in the network.

In the case that you prefer to handle configuration and deployment yourself you may choose to follow the steps below:

1. Clone [this GitHub repository](https://github.com/terminally-online/plug/) and execute the following commands:

```bash
git clone https://github.com/terminally-online/plug.git
```

2. Navigate into the solver and copy `.env.avs.example` into `.env`:

```bash
cd packages/solver && cp .env.avs.example .env
```

3. Edit the `.env` and update the OPERATOR_PRIVATE_KEY with your Operator private key.

```bash
OPERATOR_PRIVATE_KEY=<OPERATOR_PRIVATE_KEY>
```

> [!IMPORTANT]
> This private key is stored locally and never leaves your machine. During execution your key is used to sign a message and create a BLS public key that is used for aggregate signature verification.

4. Build the latest version of the Perform or Attest service scripts so that you have the latest Operator binaries.

```bash
go build ./cmd/avs/main.go avs
```

5. Run the Operator service you would like to participate in.

```bash
./avs --type=attest
```
