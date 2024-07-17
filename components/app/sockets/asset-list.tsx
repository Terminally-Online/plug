import { type FC, useCallback, useEffect, useMemo, useState } from "react"

import { LoaderCircle } from "lucide-react"

import { useBalances } from "@/contexts"
import { getPrices } from "@/lib/functions/llama/price"

import { TransferFrame } from "../frames/transfer"
import { SocketAssetItem } from "./asset-item"

type Props = {
	balances: ReturnType<typeof useBalances>["balances"]
	hasFrame?: boolean
	handleSelect?: (
		token: NonNullable<ReturnType<typeof useBalances>["balances"]>[number]
	) => void
}

export const SocketAssetList: FC<Props> = ({
	balances,
	hasFrame = true,
	handleSelect
}) => {
	const [priceData, setPriceData] = useState<
		| undefined
		| Record<
				`${string}:${string}`,
				{
					decimals: number
					symbol: string
					price: number
					timestamp: number
					confidence: number
					change: number | undefined
				}
		  >
	>()

	const coinsKey = useMemo(() => {
		if (balances === undefined) return []

		return balances.flatMap((token, index) =>
			token.chains.flatMap(
				chain => `${chain.chainName.toLowerCase()}:${chain.address}`
			)
		)
	}, [balances])

	useEffect(() => {
		if (coinsKey === undefined) return

		getPrices(coinsKey.join(",")).then(setPriceData)
	}, [coinsKey])

	useEffect(() => console.log(priceData), [priceData])

	if (balances === undefined)
		return (
			<div className="my-8 flex items-center justify-center">
				<p className="flex flex-row items-center gap-2">
					<span className="origin-center animate-spin opacity-40">
						<LoaderCircle size={14} />
					</span>
					<span className="opacity-60">Loading...</span>
				</p>
			</div>
		)

	return (
		<>
			<div className="flex flex-col gap-2">
				{balances.map(
					(token, index) =>
						token && (
							<SocketAssetItem
								key={index}
								token={token}
								priceData={priceData}
								handleSelect={handleSelect}
							/>
						)
				)}
			</div>

			{hasFrame && <TransferFrame />}
		</>
	)
}
