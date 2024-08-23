# @nftchance/plug-core

## 0.3.1

### Patch Changes

-   ee54d0f: feat: enable users to control one clickers

## 0.3.0

### Minor Changes

-   86dae2d: feat: deprecate vault as execution entry points and upgrade to router instances
-   53d6f52: feat: simplify sdk and deprecate custom type acceptance
-   142f8cf: feat: flatten architecture

### Patch Changes

-   b9eede4: chore: update to support abitypes breaking changes
-   3350d4f: fix: out of form error in PlugTypes
-   dd48d07: feat: vault deployment without explicitly provided salt
-   3350d4f: feat: improved automatic mining and etching
-   f252b33: chore: update protocol types to current version
-   3350d4f: feat: automatic deployment script generation
-   a36fa68: chore: deprecate the use of submodules as dependencies
-   3350d4f: feat: include generated contracts in package
-   3350d4f: feat: automatic wagmi hook generation
-   b37be3a: feat: add BaseFee fuse for gas fee control

## 0.2.6

### Patch Changes

-   a9d04c4: feat: simpler simulation interface
-   52dbda5: feat: deprecate use of hardhat-viem
-   33618b7: feat: fuse pass through lane

## 0.2.5

### Patch Changes

-   3d33571: feat: add artifacts export path

## 0.2.4

### Patch Changes

-   013ac9d: feat: add artifacts

## 0.2.3

### Patch Changes

-   b6c40d4: fix: src/ as key entrypoint for solidity

## 0.2.2

### Patch Changes

-   ca17edc: feat: include

## 0.2.1

### Patch Changes

-   445e859: feat: add protocol contracts to package

## 0.2.0

### Minor Changes

-   5d07ade: feat: plug-focused nomenclature

### Patch Changes

-   681380c: feat: add nouns + small cleanup

## 0.1.10

### Patch Changes

-   17e6a98: feat: no more codename

## 0.1.9

### Patch Changes

-   faf6324: feat: types version bump

## 0.1.8

### Patch Changes

-   9c17064: feat: unsigned pairs

## 0.1.7

### Patch Changes

-   5470937: feat: package exports

## 0.1.6

### Patch Changes

-   bc7f7b4: feat: roll ahead

## 0.1.5

### Patch Changes

-   f9f870f: feat: new nomenclature
-   6b9b199: feat: nomenclature spread
-   e82fe3a: feat: move up

## 0.1.4

### Patch Changes

-   827cea6: feat: optional contract inclusion

## 0.1.3

### Patch Changes

-   a566c59: fix: resolve type conflicts

## 0.1.2

### Patch Changes

-   115abf9: fix: exports

## 0.1.1

### Patch Changes

-   c9a1ab5: fix: better package bundling

## 0.1.0

### Minor Changes

-   699392b: feat: out-of-framework domain inclusion

    Before this update, the `domain` used to sign the message was held relative to the types. Now, with integration of the higher-levels it has become clear that control of the `domain` should have been included from the start.

    After this update you have the ability to run a single Plug instance for the global state rather than one per domain. This is a breaking change for the `domain` property of the `sign` method of the `Plug` class.

## 0.0.1

### Patch Changes

-   714502c: feat: initial package
