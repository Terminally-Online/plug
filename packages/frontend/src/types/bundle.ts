import { DeployIntents, MinimalProxyIntent, TokenIntents } from './intents'

// * Enables the ability to execute multiple intents in a
//   single transaction.
export type IntentBundle<TTokenType> = {
    intents: (DeployIntents | TokenIntents<TTokenType>)[]
}

export type MinimalMintBundle<TTokenType> = {
    intents: (MinimalProxyIntent | TokenIntents<TTokenType>)[]
}
