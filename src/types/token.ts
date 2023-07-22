import { AddressLike } from '.'

export type TokenTypes = 'ERC721' | 'ERC1155'

export type Token<TTokenType> = {
    type: TTokenType extends TokenTypes ? TTokenType : never
    // Contract address to mint.
    contractAddress: AddressLike
    // The tokenId to mint.
    tokenId: bigint
} & (TTokenType extends 'ERC1155'
    ? {
          // The amount of tokens being interacted with.
          amount: bigint
      }
    : {})
