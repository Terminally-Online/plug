export type Processes = Record<string, any>

export type Retries = {
    retries: number
    delay: number
}

export type NetworkBase = {
    key: string
    rpc: `wss://${string}` | `https://${string}`
    explorer: string
    explorerHasApiKey: boolean
}

export type NetworkReferences = Partial<{
    explorerApiKey: string
}> & {
    artifacts: string
    // references: References
}

export type NetworkConfig = {
}

export type Network = NetworkBase & NetworkReferences
