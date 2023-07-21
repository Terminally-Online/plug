import { AddressLike } from '.'

export type Intent<TIntent> = {
    // contract deployment details
    creator: AddressLike
    // chain id to prevent replay attacks
    chainId: 1
    // chain based nonce of signer
    nonce: 1
    // timestamp
    expiration: 0n
    // the declared action to take
    intent: TIntent
    // signed details of the full order
    signature: `0x${string}`
}

// * Enables the creator to delay contract deployment until
//   the first mint is called.
export type Create2Intent = Intent<{
    deployer: AddressLike
    salt: `0x${string}`
    bytecodeHash: `0x${string}`
}>

// * Enables creators to declare the intent for a token to exist
//   without having to deploy the contract or mint the contract.
// ! A valid LazyMintIntent must be for a contract that already exists,
//   or for an intent that is stored in the Create2Intent.
export type LazyMintIntent = Intent<{
    // TODO: Spec. out
}>

// * Enables creators to declare economic actions for tokens before
//   they are deployed or minted.
// ! A valid SellIntent must be for a contract that already exists,
//   or for an intent that is stored in the LazyMintIntent.
export type SellIntent = Intent<{
    // TODO: Spec. out
}>

export type BuyIntent = Intent<{
    // TODO: Spec. out
}>
