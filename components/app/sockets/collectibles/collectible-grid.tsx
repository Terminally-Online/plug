import { FC } from "react"

import { motion } from "framer-motion"

import { SocketCollectibleItem } from "./collectible-item"

type Props = { expanded?: boolean }

export const SocketCollectibleGrid: FC<Props> = ({ expanded = false }) => {
	const balances = undefined

	return (
		<motion.div
			className="mb-24 grid gap-2"
			style={{
				gridTemplateColumns: `repeat(auto-fill, minmax(200px, 1fr))`
			}}
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
			{(balances || Array(6).fill(undefined)).map(
				(collectible, index) => (
					<SocketCollectibleItem key={index} />
				)
			)}
		</motion.div>
	)
}
