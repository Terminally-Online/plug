# Reads

[Reads](#reads) are a critical piece of Plug as they enable embedded reactions when powered by [Coils](./coils).

With reads, During the execution of your intent it is extremely straight forward to prepend an onchain function read or offchain read echo. With this, you can directly utilize the value of an offchain read without needing an asynchronous state update.

Embedded reads are the purest form of interoprability enabled by transaction composability.

## Types of Reads

When building Plugs inside the application there exist two primary types that provide coverage for nearly every possible thing you can think of:

- [Embedded](#embedded): Onchain reads are used as source data for coil updates.
- [Echoes](#echoes): Offchain echoes securely bring offchain data onchain for single transaction use.

### Embedded

In many cases there are actions that depend on using data that is not stale or out of date by even a second. With embedded reads we postpone the onchain read balance until it is actually needed; in this case, the exact same block the rest of your transaction is running. This way, you have zero malformed or stale data risk.

**Why read something offchain when you can read it onchain?** There's no reason when you have the ability to use embedded reads.

A simple model for your understanding is a liquidity in a pool or token balance. This data exists onchain and if you want to use it in your transaction you historically need to make an offchain read before you run the transaction. That's no longer the case.

### Echoes

Data can be brought onchain and made usable by [Coils](./coils) through a process called echoing. When offchain data is echoed, it is posted onchain within the transaction so that it can be referred to by [Coils](./coils).

This is particularly important when enforcing constraints as a piece of information may only be available offchain and must be settled atomically onchain alongside all other aspects of the transaction.

To maintain security and verifiability through all pieces of the stack, every built route and Plug execution is validated by [Attesters](../execution/circuit#attester) operating in [Circuit](../execution/circuit).
