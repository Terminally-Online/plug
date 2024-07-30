import { FC, useMemo, useState } from "react"

import { motion } from "framer-motion"

import { Button, SocketTokenItem } from "@/components"
import { useBalances } from "@/contexts"
import { RouterOutputs } from "@/server/client"

type Props = {
	expanded?: boolean
	handleSelect?: (
		token: NonNullable<RouterOutputs["socket"]["tokens"]>[number]
	) => void
}

export const SocketTokenList: FC<Props> = ({ expanded, handleSelect }) => {
	const { tokens } = useBalances()

	const visibleTokens = useMemo(() => {
		if (tokens === undefined) return Array(5).fill(undefined)

		if (expanded) return tokens

		return tokens.filter(token => token.totalValue > 1)
	}, [expanded, tokens])

	return (
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
			{visibleTokens.map((token, index) => (
				<SocketTokenItem
					key={index}
					token={token}
					handleSelect={handleSelect}
				/>
			))}
		</motion.div>
	)
}
