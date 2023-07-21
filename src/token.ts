import { AddressLike } from '.'
import { Create2Intent, Intent, LazyMintIntent } from './intents'

export type CollectionModules = [
    AddressLike[], // Transfer hooks
    AddressLike[], // Mint hooks
    AddressLike[], // Burn hooks
]

export type Collection<TIntent extends Create2Intent | LazyMintIntent> = {
    // * The contract may or may not have been declared with an intent
    intent?: Intent<TIntent>

    // * If the collection was deployed through an intent, this was acquired
    //   from in-database otherwise it was acquired from the contract.
    name: 'Name'
    description: 'Description'
    image: 'ipfs://'
    external_url: 'https://example.com/'

    // TODO: Figure out royalties onchain and offchain.
    royalties: {
        bps: 1000
    }

    // extension enabled functionality
    modules: [AddressLike[]]
}
