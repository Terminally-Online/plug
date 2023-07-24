import { AddressLike } from '.'
import { Create2Intent, MinimalProxyIntent } from './intents'

// TODO: Figure out how to get token type on collection.
// TODO: Figure out royalties onchain and offchain.

type DeployIntent = Create2Intent | MinimalProxyIntent

export type Collection<TIntent = undefined> = {
    // * The contract may or may not have been declared with an intent.
    intent?: TIntent extends DeployIntent ? TIntent : undefined

    deployer: AddressLike
    owner: AddressLike

    // * If the collection was deployed through an intent, this was acquired
    //   from in-database otherwise it was acquired from the contract.
    name: 'Name'
    description: 'Description'
    image: 'ipfs://'
    external_url: 'https://example.com/'

    // * Definition of extension enabled functionality.
    extensions: AddressLike[]
    modules: [
        AddressLike[], // Transfer hooks
        AddressLike[], // Mint hooks
        AddressLike[], // Burn hooks
    ]
}
