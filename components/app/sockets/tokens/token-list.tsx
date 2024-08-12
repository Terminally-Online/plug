import { FC, HTMLAttributes, useMemo } from "react"

import { motion, MotionProps } from "framer-motion"

import { SocketTokenItem } from "@/components"
import { useBalances } from "@/contexts"
import { cn } from "@/lib"
import { RouterOutputs } from "@/server/client"

export const SocketTokenList: FC<
	HTMLAttributes<HTMLDivElement> &
		MotionProps & {
			id: string
			expanded?: boolean
			handleSelect?: (
				token: NonNullable<
					RouterOutputs["socket"]["balances"]["tokens"]
				>[number]
			) => void
		}
> = ({ id, expanded, handleSelect, className, ...props }) => {
	const { tokens } = useBalances()

	const visibleTokens = useMemo(() => {
		if (tokens === undefined) return Array(5).fill(undefined)

		if (expanded) return tokens

		return tokens.slice(0, 5)
	}, [expanded, tokens])

	return (
		<motion.div
			className={cn("mb-4 flex flex-col gap-2", className)}
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
					id={id}
					token={token}
					handleSelect={handleSelect}
				/>
			))}
		</motion.div>
	)
}
