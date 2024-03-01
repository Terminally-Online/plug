---
head:
  - - meta
    - property: og:title
      content: Fuse | Limited Calls
  - - meta
    - name: description
      content: declare the amount of times a single intent can be reused.
  - - meta
    - property: og:description
      content: Declare a schedule for which your intents can be executed on.
---

# Limited Calls Fuse

In Plug, all intents operate as constant processes unless specified otherwise. The use of a `nonce` is optional to prevent the need of signing `N` orders when intents are to be executed many times. This simple [Fuse](core/fuses) enables the ability to have non-ordered global execution while maintaining the linear state of a single intent.

## Logic

Any intent using the `Limited Calls Fuse` has the ability to declare:

- Call Count: How many times this specific intent can be called.

The limit imposed by this Fuse can be completely bypassed by not including it. The limit can be set to 1 or even 100.

[Plug](/) is designed around consumption of resources (and permission) rather than forced non-incremental nonces. In reality, not all intents should expire unless directly terminated.
