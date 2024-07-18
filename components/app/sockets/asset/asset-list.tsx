import { FC, useEffect, useMemo, useState } from "react"

import { motion } from "framer-motion"
import { LoaderCircle } from "lucide-react"

import { SocketAssetItem, TransferFrame } from "@/components"
import { useBalances } from "@/contexts"
import { getPrices } from "@/lib"

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
			<motion.div
				className="flex flex-col gap-2"
				initial="hidden"
				animate="visible"
				variants={{
					hidden: { opacity: 0 },
					visible: {
						opacity: 1,
						transition: {
							staggerChildren: 0.05
						}
					}
				}}
			>
				{balances.map(
					(token, index) =>
						token && (
							<motion.div
								key={index}
								variants={{
									hidden: { opacity: 0, y: 10 },
									visible: {
										opacity: 1,
										y: 0,
										transition: {
											type: "spring",
											stiffness: 100,
											damping: 10
										}
									}
								}}
							>
								<SocketAssetItem
									token={token}
									priceData={priceData}
									handleSelect={handleSelect}
								/>
							</motion.div>
						)
				)}
			</motion.div>

			{hasFrame && <TransferFrame />}
		</>
	)
}
