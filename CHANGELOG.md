# @nftchance/emporium-core

## 0.1.1

### Patch Changes

-   c9a1ab5: fix: better package bundling

## 0.1.0

### Minor Changes

-   699392b: feat: out-of-framework domain inclusion

    Before this update, the `domain` used to sign the message was held relative to the types. Now, with integration of the higher-levels it has become clear that control of the `domain` should have been included from the start.

    After this update you have the ability to run a single Framework instance for the global state rather than one per domain. This is a breaking change for the `domain` property of the `sign` method of the `Intent` class.

## 0.0.1

### Patch Changes

-   714502c: feat: initial package
