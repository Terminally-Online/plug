import { FC, useEffect, useMemo, useState } from "react"

import { motion } from "framer-motion"

import { SocketTokenItem } from "@/components"
import { useBalances, useSockets } from "@/contexts"
import { getPrices } from "@/lib"

type Props = {
	expanded?: boolean
	handleSelect?: (
		token: NonNullable<ReturnType<typeof useBalances>["balances"]>[number]
	) => void
}

// TODO: Handle expanded.

export const SocketTokenList: FC<Props> = ({ handleSelect }) => {
	const { socket } = useSockets()
	const { balances } = useBalances({
		address: socket?.socketAddress || ""
	})

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

		return balances.flatMap(token =>
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
				{(balances || Array(8).fill(undefined)).map((token, index) => (
					<SocketTokenItem
						key={index}
						token={token}
						priceData={priceData}
						handleSelect={handleSelect}
					/>
				))}
			</motion.div>
		</>
	)
}
