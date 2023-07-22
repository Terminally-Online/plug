import { AddressLike } from '.'
import { DeployIntent } from './intents'

// TODO: Figure out how to get token type on collection
// TODO: Figure out royalties onchain and offchain.

export type Collection<TIntent> = {
    // * The contract may or may not have been declared with an intent
    intent?: TIntent extends DeployIntent ? TIntent : never

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
