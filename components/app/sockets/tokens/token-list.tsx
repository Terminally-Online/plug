import { FC, HTMLAttributes, useMemo, useState } from "react"

import { motion, MotionProps } from "framer-motion"

import { Button, SocketTokenItem } from "@/components"
import { useBalances } from "@/contexts"
import { cn } from "@/lib"
import { RouterOutputs } from "@/server/client"

type Props = {
	expanded?: boolean
	handleSelect?: (
		token: NonNullable<RouterOutputs["socket"]["tokens"]>[number]
	) => void
} & MotionProps &
	HTMLAttributes<HTMLDivElement>

export const SocketTokenList: FC<Props> = ({
	expanded,
	handleSelect,
	className,
	...props
}) => {
	const { tokens } = useBalances()

	const visibleTokens = useMemo(() => {
		if (tokens === undefined) return Array(5).fill(undefined)

		if (expanded) return tokens

		return tokens.filter(token => token.totalValue > 1)
	}, [expanded, tokens])

	return (
		<motion.div
			className={cn("flex flex-col gap-2", className)}
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
			{...(props as MotionProps)}
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
