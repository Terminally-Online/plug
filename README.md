# 🔌 Plug

`@nftchance/plug` powers "**if this, then that**" statements for EVM blockchain transactions. Without having to change the core logic of your protocol you can turbocharge your protocol with declarative transaction execution in just a few seconds.

> [!TIP]
> While Plug is a protocol, it has been designed to first solve for the problems experienced by end-users. Not developers. You can find in-depth documentation [here](https://onplug.io). If you are a developer, builder, creator or investor and want to talk or have questions, please [send a DM on Twitter](https://twitter.com/nftchance).

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

## Licensing

Plug is setup with a very clear line of licensing.

All of the code inside `packages/app` falls under a [BUSL-1.1 license](./packages/app/LICENSE) that requires a license key for business usage and sets forth subsequent requirements of operations. At this time, a license cannot be purchased however this capability will become available to you some time mid-2024. For personal use, you are allowed to run your own local instance without acquiring a license key. Please refer to [the license](./packages/app/LICENSE) for complete detail and coverage of what is and is not allowed.

Code beyond the scope of `packages/app` falls under a [MIT license]() and is free for personal and business use in accordance with the rest of the guidelines established herein.

### A Short Explanation

The license that is applied to open source, and work surrounding EVM blockchains in general is critical to viability of adoption and verifiability. I do not want to release code that has disincentivized the use or consumption of. At the same time, it is of paramount importance that the leading development of Plug has the resources and access to persist in delivering the user experience and interface of the quality that you, the user, deserve.

To make this happen, I have taken a slightly abnormal approach by applying two different licenses throughout the entire codebase of Plug. This results in the ability of having granular control over the pieces of the codebase that assist in the value creation of the platform without putting me or any user in line of newfound and unjustified risk that arises from closed source development.

To do this, licensing is simplified into one simple concept:

-   Does an individual user benefit from having code control over this piece?
    -   If yes, a MIT license is applied,
    -   otherwise a BUSL-1.1 license is applied to the respective code and surrounding area.

This results in a paradigm where automatically, pieces critical to the onchain protocol are made open source and designed to defend against users that do not contribute value back to the larger pie. AGPL-3.0 is not a copyright license, but copyleft. The entire Plug codebase has been designed to provide the option of running your own local instance. I am less worried about general individual use, and more concerned with ensuring that no one unjustly takes this code, claims it as their own and pushes it into a closed source environment. An onchain protocol has no moat beyond the cultural gravity surrounding it. Both figuratively and literally, you are the moat of the Plug protocol.

As for [the browser based node-editor application](./packages/app/), a strict business license has been applied to all pieces driving this specific piece of functionality that are not already covered by a lower-level AGPL-3.0 through environmental inclusion. This has been done for several reasons, but namely that while there is no reason a business should ever be running their own version of the app, it is important for you, the user, to be have full access to read the code that is deployed. The application of [BUSL-1.1](./packages/app/LICENSE) is an effort to maximize transparency without sacrificing the future abilities of the supporting team.

Specifically, the license of each package is as follows:

```ml
packages
├─ app — "BUSL-1.1"
├─ client — "MIT"
├─ core — "MIT"
├─ docs — "MIT"
├─ landing — "BUSL-1.1 -- Soon to be deprecated with rollover to MIT."
├─ server — "MIT"
└─ types — "MIT"
```

Once again, individual and use without distribution is fully permitted. Additionally, users have full capability to direct the data generated from the application to an individually run backend (that is AGPL-3.0) meaning you can effectively gain all the benefits of protocol and app without concern. With all this said, if you are still not satisifed you can reference [./packages/client](./packages/client) to see a raw implementation of a protocol and api consumer that you can fork and change without limit beyond the license applied.

## Using a Package

All `@nftchance/plug-*` packages have been built with one thing in mind:

_using your brain as little as possible._

What this means is that every step of the way, I simply designed around my assumptions as if I was someone with no experience. Due to this, the commands to `build`, `dev`, `lint`, etc. are the same across every package. Of course, some commands are missing from certain packages where they are not relevant. The base set will include:

```ml
├─ build — "Compile all the artifacts for production distribution."
├─ dev — "Run the development version of the stack needed to drive the active package."
└─ generate — "Create, update, and manage the static files generated by Plug."
```

The basic rule is that to build the package you will `cd` into it and run `pnpm dev` to get up and running. For information beyond that please reference the local `package.json`.
