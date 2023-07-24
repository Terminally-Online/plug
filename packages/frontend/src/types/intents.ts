import { AddressLike } from '.'
import { Token, TokenTypes } from './token'

// TODO: implement intents that support token value, and not just ETH value
export type ContractIntentTypes = 'Create2' | 'MinimalProxy'
export type TokenIntentTypes = 'Mint' | 'Transfer' | 'List' | 'Buy' | 'Offer'

export type Intent<
    TIntentType extends ContractIntentTypes | TokenIntentTypes,
    TIntent,
> = {
    // signer of the intent
    intender: AddressLike
    // chain id to prevent replay attacks
    chainId: number
    // chain based nonce of signer
    nonce: bigint
    // timestamp
    expiration: bigint
    // the declared action to take
    intent: {
        type: TIntentType
    } & TIntent
    // signed details of the full order
    signature: `0x${string}`
}

export type ContractIntent<
    TIntentType extends ContractIntentTypes,
    TIntent = {},
> = Intent<TIntentType, TIntent>

export type TokenIntent<
    TIntentType extends TokenIntentTypes,
    TTokenType,
    TIntent = {},
> = Intent<
    TIntentType,
    {
        // Token to schedule intent for.
        token: Token<TTokenType extends TokenTypes ? TTokenType : never>
        // The value of the intent.
        value: bigint
    } & TIntent
>

// * Enables the creator to delay contract deployment until
//   the first mint is called with Create2.
// ! The use of this lowers the cost of transfers though
//   increases the cost of deployment.
export type Create2Intent = ContractIntent<
    'Create2',
    {
        // The salt used with create2 to determine the address
        salt: `0x${string}`
        // The bytecode hash of the contract code & constructor arguments
        bytecodeHash: `0x${string}`
    }
>

// * Enables the creator to delay contract deployment until
//   the first mint is called with MinimalProxy.
// ! The use of this lowers the cost of deployment though
//   increases the cost of transfers.
export type MinimalProxyIntent = ContractIntent<
    'MinimalProxy',
    {
        // * The base implementation of the contract being cloned.
        implementation: AddressLike
        // * The encoded value of the arguments passed to the
        //   initialization function of the contract.
        structBytes: `0x${string}`
    }
>

// * Enables creators to declare the intent for a token to exist
//   without having to deploy the contract or mint the contract.
// ! A valid MintIntent must be for a contract that already exists,
//   or for an intent that is stored in the DeployIntent.
export type MintIntent<TTokenType> = TokenIntent<
    'Mint',
    TTokenType,
    {
        // The address to transfer the token to.
        to: AddressLike
    }
>

// * Enables creators to declare the intent for a token to be
//   transferred without having to mint the token.
export type TransferIntent<TTokenType> = TokenIntent<'Transfer', TTokenType>

export type Listing = {
    // The address to list the token to for private listings.
    to: AddressLike
    // TODO: Add economic action definition.
    value: bigint
}

// * Enables creators to declare economic actions for tokens before
//   they are deployed or minted.
// ! A valid ListIntent must be for a contract that already exists,
//   or for an intent that is stored in the MintIntent.
export type ListIntent<TTokenType> = TokenIntent<
    'List',
    TTokenType,
    {
        // The address to list the token to for private listings.
        to: AddressLike
        // TODO: Add economic action definition.
        value: bigint
    }
>

// * Enables collectors to declare intent to buy a token without
//   a token having been declared.
// ! A valid BuyIntent must be for a token that has already has
//   a counterpart in the ListIntent.
export type BuyIntent<TTokenType> = TokenIntent<'Buy', TTokenType>

// * Enables collectors to declare intent to make an offer on a token
//   without the order being executed.
export type OfferIntent<TTokenType> = TokenIntent<'Offer', TTokenType>

// * Non-fungible token intents.
export type ERC721MintIntent = MintIntent<'ERC721'>
export type ERC721TransferIntent = TransferIntent<'ERC721'>
export type ERC721ListIntent = ListIntent<'ERC721'>
export type ERC721BuyIntent = BuyIntent<'ERC721'>
export type ERC721OfferIntent = OfferIntent<'ERC721'>

// * Fungible token intents.
export type ERC1155MintIntent = MintIntent<'ERC1155'>
export type ERC1155TransferIntent = TransferIntent<'ERC1155'>
export type ERC1155ListIntent = ListIntent<'ERC1155'>
export type ERC1155BuyIntent = BuyIntent<'ERC1155'>
export type ERC1155OfferIntent = OfferIntent<'ERC1155'>

// * Bundle driven intents.
export type DeployIntents = Create2Intent | MinimalProxyIntent
// TODO: Collection intents (This will set metadata and royalties.)
export type TokenIntents<TTokenType> =
    | MintIntent<TTokenType>
    | TransferIntent<TTokenType>
    | ListIntent<TTokenType>
    | BuyIntent<TTokenType>
    | OfferIntent<TTokenType>
