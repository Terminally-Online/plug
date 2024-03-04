# Sockets

Bringing the key execution capabilities of [Plug](/) are `Sockets`. A `Socket` is a smart contract where [Plugs](generated/base-types/Plugs) are routed to and executed from.

Operationally, `Sockets` mirror the functionality of the concept commonly thought of as blockchain accounts with a few small caveats:

- `Permission Allowance`: Multiple accounts can be granted operational permission.
- `Conditional Enforcement`: Ensure a list of conditions are satisfied before executing.
- `Intent Reuse`: Intents are designed to be reused many times in the future.
- `Native Paymasters`: Compensate external parties for execution.

Combining these capabilities together, `Sockets` introduce the ability to have safe automated transactions that can support effectively every smart contract that exists.

## The Abstraction

In the simplest forms, `Sockets` are designed to provide the core functionality of execution. With that though, there are several different forms of `Sockets` that aim to fulfill very specific needs and wants. To provide this flexibility `Sockets` are designed on a single extendable pattern that provide several opportunities to override the logic ran during simulation and execution.

The methods available for public interfacing are:

- [signer](/core/sockets/signer): Retrieve the signer of a [LivePlugs](/generated/base-types/LivePlugs) bundle.
- [plug](/core/sockets/plug): Execute a [LivePlugs](/generated/base-types/LivePlugs) bundle.

While these are the two key functions for public use, there are also internal functions that meant to be overridden including:

- `_enforceRouter`: Extend the validation logic of [Routers](/core/routers).
- `_enforceSigner`: Extend the validation logic of who can be a signer.
- `_enforceFuse`: Extend the validation logic of [Fuse](/core/fuses) conditions.
- `_enforceCurrent`: Extend the validation logic of the executable [Currents](/generated/base-types/Current).

With these pairings, a `Socket` can be configured to provide the precise functionality desired with ease and without needing to manipulate the core logic of the architecture.
