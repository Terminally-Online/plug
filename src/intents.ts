import { AddressLike } from '.'
import { Token, TokenTypes } from './token'

// TODO: implement intents that support token value, and not just ETH value

type Intent<
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

type ContractIntentTypes = 'Deploy'
type ContractIntent<TIntentType extends ContractIntentTypes> = Intent<
    TIntentType,
    {
        // The salt used with create2 to determine the address
        salt: `0x${string}`
        // The bytecode hash of the contract code & constructor arguments
        bytecodeHash: `0x${string}`
    }
>

type TokenIntentTypes = 'Mint' | 'Transfer' | 'List' | 'Buy'
type TokenIntent<TIntentType, TTokenType, TIntent = {}> = Intent<
    TIntentType extends TokenIntentTypes ? TIntentType : never,
    {
        // Token to schedule intent for.
        token: Token<TTokenType extends TokenTypes ? TTokenType : never>
        // The value of the intent.
        value: bigint
    } & TIntent
>

// * Enables the creator to delay contract deployment until
//   the first mint is called.
export type DeployIntent = ContractIntent<'Deploy'>

// * Enables creators to declare the intent for a token to exist
//   without having to deploy the contract or mint the contract.
// ! A valid MintIntent must be for a contract that already exists,
//   or for an intent that is stored in the DeployIntent.
type MintIntent<TTokenType> = TokenIntent<
    'Mint',
    TTokenType,
    {
        // The address to transfer the token to.
        to: AddressLike
    }
>

export type ERC721MintIntent = MintIntent<'ERC721'>
export type ERC1155MintIntent = MintIntent<'ERC1155'>

// * Enables creators to declare the intent for a token to be
//   transferred without having to mint the token.
type TransferIntent<TTokenType> = TokenIntent<'Transfer', TTokenType>

export type ERC721TransferIntent = TransferIntent<'ERC721'>
export type ERC1155TransferIntent = TransferIntent<'ERC1155'>

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
    }
>

export type ERC721ListIntent = ListIntent<'ERC721'>
export type ERC1155ListIntent = ListIntent<'ERC1155'>

// * Enables collectors to declare intent to buy a token without
//   a token having been declared.
// ! A valid BuyIntent must be for a token that has already has
//   a counterpart in the ListIntent.
export type BuyIntent<TTokenType> = TokenIntent<'Buy', TTokenType>

export type ERC721BuyIntent = BuyIntent<'ERC721'>
export type ERC1155BuyIntent = BuyIntent<'ERC1155'>
