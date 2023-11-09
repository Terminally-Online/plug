# @nftchance/emporium

`Emporium` powers "**if this, then that**" statements for EVM blockchains. Without having to change the core logic of your protocol you can turbocharge your protocol with declarative transaction execution in just a few seconds. 

The key functionality is powered by the following packages (each have their own function):

```ml
packages
├─ types — "Automatically generate the types and decoders of your intent framework."
├─ core — "Intent framework smart contracts and management utilities."
└─ docs — "In-depth documentation for both end-users and developers."
└─ core — "Intent framework smart contracts and management utilities."
```

> **Note**
> While you can browse the codebase and read the sparse implementation/maintenance focused documentation, this is likely not what you are looking for unless you intend of making a contribution directly to `emporium`.

## Using a Package

All `Emporium` packages have been built with one thing in mind: using your brain as little as possible. 

What this means is that every step of the way, I simply designed around my assumptions as if I was someone with no experience. Due to this, the commands to `build`, `run`, `lint`, etc. are the same acros every package. Of course, some commands are missing from certain packages where they are not relevant, but you can always check the `package.json` of the package you are working within.

The basic rule is that to build the package you will `cd` into it and run `pnpm build` and to get up and running you just need to use `pnpm dev`.

Additionally, in all packages that are released for independent consumption `changesets/cli` is used to power a very simple package building and distribution pipeline. When making a change that requires a release simply run `pnpm changeset add` and configure the preferred type of update whether it be `major`, `minor` or `patch`.

