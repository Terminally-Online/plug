import { chains } from "@/lib"

export type ChainId = (typeof chains)[number]["id"]

interface ProviderRpcError extends Error {
	message: string
	code: number
	data?: unknown
}
interface ProviderMessage {
	type: string
	data: unknown
}
interface ProviderInfo {
	chainId: string
}
declare type ProviderChainId = ProviderInfo["chainId"]
declare type ProviderAccounts = string[]

declare namespace IProviderEvents {
	type Event = "connect" | "disconnect" | "message" | "chainChanged" | "accountsChanged" | "display_uri"
	interface EventArguments {
		connect: ProviderInfo
		disconnect: ProviderRpcError
		message: ProviderMessage
		chainChanged: ProviderChainId
		accountsChanged: ProviderAccounts
		display_uri: string
	}
}

export type WalletConnectProvider = {
	connect(opts?: {
		chains?: number[]
		optionalChains?: number[]
		rpcMap?: {
			[chainId: string]: string
		}
		pairingTopic?: string
	}): Promise<void>
	once: <E extends IProviderEvents.Event>(
		event: E,
		listener: (args: IProviderEvents.EventArguments[E]) => void
	) => WalletConnectProvider
}
