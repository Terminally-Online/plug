import { FC } from "react"

import { motion } from "framer-motion"

import { SocketTokenItem } from "@/components"
import { useBalances } from "@/contexts"
import { RouterOutputs } from "@/server/client"

type Props = {
	expanded?: boolean
	handleSelect?: (
		token: NonNullable<RouterOutputs["socket"]["tokens"]>[number]
	) => void
}

// TODO: Fix whatever is wrong with drakes token retrieval.
// 		NOTES: Gut feeling is that it has nothing to do with the amount and some request is throwing for some reason and it is not properly handled.
// TODO: Implement see all that minimizes the list down to tokens than have > $5 dollars.

export const SocketTokenList: FC<Props> = ({ handleSelect }) => {
	const { tokens } = useBalances()

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
				{(tokens || Array(5).fill(undefined)).map((token, index) => (
					<SocketTokenItem
						key={index}
						token={token}
						handleSelect={handleSelect}
					/>
				))}
			</motion.div>
		</>
	)
}
