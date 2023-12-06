# 🔌 Plug

`@nftchance/plug` powers "**if this, then that**" statements for EVM blockchain transactions. Without having to change the core logic of your protocol you can turbocharge your protocol with declarative transaction execution in just a few seconds.

The key functionality is powered by the following packages (each have their own function):

```ml
packages
├─ app — "Browser based node-editor application to interact with the Plug protocol."
├─ client — "Raw playground to build, sign and save Plug objects."
├─ core — "Plug framework smart contracts and management utilities."
├─ docs — "In-depth documentation for both end-users and developers."
├─ landing — "Marketing landing page for the Plug ecosystem."
├─ server — "API backend that powers the server, client interface and sdk when needed."
└─ types — "Automatically generate the types and decoders of your intent framework."
```

> [!NOTE]
> While Plug is a protocol, it has been designed to first solve for the problems experienced by end-users. Not developers. You can find in-depth documentation [here](https://onplug.io). If you are a developer, builder, creator or investor and want to talk or have questions, please [send a DM on Twitter](https://twitter.com/nftchance).

## Using a Package

All `@nftchance/plug-*` packages have been built with one thing in mind: _using your brain as little as possible._

What this means is that every step of the way, I simply designed around my assumptions as if I was someone with no experience. Due to this, the commands to `build`, `run`, `lint`, etc. are the same across every package. Of course, some commands are missing from certain packages where they are not relevant, but you can always check the `package.json` of the package you are working within.

The basic rule is that to build the package you will `cd` into it and run `pnpm dev` to get up and running.
