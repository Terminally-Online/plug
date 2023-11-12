# @nftchance/plug

`Plug` powers "**if this, then that**" statements for EVM blockchains. Without having to change the core logic of your protocol you can turbocharge your protocol with declarative transaction execution in just a few seconds.

The key functionality is powered by the following packages (each have their own function):

```ml
packages
├─ client — "Raw playground to build, sign and save Plug objects."
├─ core — "Intent framework smart contracts and management utilities."
├─ docs — "In-depth documentation for both end-users and developers."
├─ landing — "Marketing landing page for the Plug ecosystem."
├─ server — "API backend that powers the server, client interface and sdk when needed."
└─ types — "Automatically generate the types and decoders of your intent framework."
```
> **Important**
> While Plug is a protocol, it has been designed to first solve for the problems experienced by end-users, not developers. You can find in-depth documentation [here](https://onplug.io).

## Using a Package

All `Plug` packages have been built with one thing in mind: using your brain as little as possible.

What this means is that every step of the way, I simply designed around my assumptions as if I was someone with no experience. Due to this, the commands to `build`, `run`, `lint`, etc. are the same acros every package. Of course, some commands are missing from certain packages where they are not relevant, but you can always check the `package.json` of the package you are working within.

The basic rule is that to build the package you will `cd` into it and run `pnpm build` and to get up and running you just need to use `pnpm dev`.

Additionally, in all packages that are released for independent consumption `changesets/cli` is used to power a very simple package building and distribution pipeline. When making a change that requires a release simply run `pnpm changeset add` and configure the preferred type of update whether it be `major`, `minor` or `patch`.
