# Enrollment

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
