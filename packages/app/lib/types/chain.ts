import { Chain as WagmiChain } from "wagmi/chains"

export enum RPCType {
	Public = "public",
	Private = "private",
	PublicAlt = "public_alternative",
	AppOnly = "app"
}

type ChainRpcUrls = {
	http: readonly string[]
	webSocket?: readonly string[] | undefined
}

export interface Chain extends WagmiChain {
	color: `#${string}` | `rgb(${number}, ${number}, ${number})` | `rgba(${number}, ${number}, ${number}, ${number})`
	logo: `${string}.png`
	blockExplorers?: WagmiChain["blockExplorers"]
	rpcUrls: WagmiChain["rpcUrls"] & { [RPCType.AppOnly]: ChainRpcUrls }
}
