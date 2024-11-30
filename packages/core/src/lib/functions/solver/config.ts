import { author, maintainers } from "@/package.json"
import {
	BaseConfig,
	Config,
	DEFAULT_NETWORKS,
	DEFAULT_NETWORK_CONFIG,
	DEFAULT_NETWORK_REFERENCES,
	DEFAULT_NETWORK_RETRIES,
	DEFAULT_SCHEMA
} from "@/src/lib"

export type MaybeArray<T> = T | T[]

export const name = "plug" as const

export const defineConfig = (base: BaseConfig): Config => {
    // * Destructure the network configuration from the base configuration.
    let { networks, retries, delay, contract, out, dangerous, types } = base

    types = types
        ? {
            ...DEFAULT_SCHEMA.types,
            ...types
        }
        : DEFAULT_SCHEMA.types

    const config: Config = {
        // * Overwrite the default retry configuration with the provided
        //   values when they are not undefined.
        retries: retries ?? DEFAULT_NETWORK_RETRIES.retries,
        delay: delay ?? DEFAULT_NETWORK_RETRIES.delay,
        /// * Set the default networks for the Engine.
        networks: {},
        out: {
            ...DEFAULT_SCHEMA.config.out,
            ...out
        },
        contract: {
            ...DEFAULT_SCHEMA.config.contract,
            ...contract,
            authors: [author, ...maintainers]
                .concat(contract?.authors ?? [])
                .map(author => ` * @author ${author}`)
                .join("\n")
        },
        dangerous: {
            ...DEFAULT_SCHEMA.config.dangerous,
            ...dangerous
        },
        types
    }

    // * While we have a set of default networks, the configuration only
    //   contains the chains that the user has actually configured.
    for (const networkId in networks) {
        if (networks[networkId] === undefined) continue

        const network = {
            // * Append the default values to the network so that the user
            //   can operate with the default RPC and Explorer URLs when
            //   they are not provided.
            ...DEFAULT_NETWORKS[networkId],
            // * Set the default values of the references for the network.
            // ! Realistically, most implementations of the Engine configuration
            //   will have at least some amount of contract references, but
            //   we can't assume that they will be provided at time of loading.
            ...DEFAULT_NETWORK_REFERENCES,
            // * Set the operational pieces of the Engine as null
            //   and then overwrite them with the provided values
            //   when they are not undefined.
            ...DEFAULT_NETWORK_CONFIG,
            // * Finally set all the fields that the user provided in
            //   their configuration instantiation.
            ...networks[networkId]
        }

        config.networks[networkId] = network
    }

    if (Object.keys(config).length === 0) {
        console.warn("No networks were configured.")
    }

    return config
}

export async function configs() {
    return [defineConfig({ networks: {} })]
}
