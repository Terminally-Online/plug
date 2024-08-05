import { FC, HTMLAttributes, useMemo, useState } from "react"

import { motion, MotionProps } from "framer-motion"

import { SocketTokenItem } from "@/components"
import { useBalances } from "@/contexts"
import { cn } from "@/lib"
import { RouterOutputs } from "@/server/client"

export const SocketTokenList: FC<
	HTMLAttributes<HTMLDivElement> &
		MotionProps & {
			expanded?: boolean
			handleSelect?: (
				token: NonNullable<
					RouterOutputs["socket"]["balances"]["tokens"]
				>[number]
			) => void
		}
> = ({ expanded, handleSelect, className, ...props }) => {
	const { tokens } = useBalances()

	const visibleTokens = useMemo(() => {
		if (tokens === undefined) return Array(5).fill(undefined)

		if (expanded) return tokens

		return tokens.filter(token => token.value > 1)
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
