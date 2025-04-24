import { RouterOutputs } from "@/server/client"

import { getChainId } from "../blockchain"
import { zerionChains } from "./addresses"

export * from "./positions"
export * from "./collectible"
export * from "./tokens"

export type ZerionPositions = NonNullable<RouterOutputs["service"]["zerion"]["wallet"]["positions"]["data"]>
export type ZerionPosition = NonNullable<ZerionPositions>[number]

export type ZerionFungible = NonNullable<RouterOutputs["service"]["zerion"]["fungibles"]["list"]["data"]>[number]

export const getZerionTokenIconUrl = (token: ZerionPosition | string | undefined) => {
	if (!token) return ""

	let url
	if (typeof token === "string") url = token
	else url = token.attributes.fungible_info.icon?.url

	if (url) return url
	if (typeof token === "string") return ""

	const implementations = token.attributes.fungible_info.implementations

	if (!implementations || implementations.length === 0) return ""

	const chainId = getChainId(implementations[0].chain_id)
	const address = implementations[0].address
	return `https://token-icons.llamao.fi/icons/tokens/${chainId}/${address}?h=240&w=240`
}

export const getZerionChainIconUrl = (chainId: string) => {
	return zerionChains[chainId as keyof typeof zerionChains]?.attributes.icon?.url
}
