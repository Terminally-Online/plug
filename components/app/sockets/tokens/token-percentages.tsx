import type { FC } from "react"

import { useBalances } from "@/contexts"

type Props = {
	chains: NonNullable<
		ReturnType<typeof useBalances>["balances"]
	>[number]["chains"]
}

const size = 14
const radius = 50
const circumference = 2 * Math.PI * radius
const viewbox = 120

const getDashArray = (percentage: number) => {
	const filledLength = (percentage / 100) * circumference
	const emptyLength = circumference - filledLength
	return `${filledLength} ${emptyLength}`
}

const getChainColor = (chainId: number) => {
	switch (chainId) {
		case 1:
			return "#393939"
		case 10:
			return "#FF0420"
		case 8453:
			return "#0052FF"
		default:
			return "#393939"
	}
}

export const SocketTokenPercentages: FC<Props> = ({ chains }) => {
	let accumulatedPercentage = 0

	return (
		<svg width={size} height={size} viewBox={`0 0 ${viewbox} ${viewbox}`}>
			<g transform={`translate(${viewbox / 2},${viewbox / 2})`}>
				{chains.map((chain, index) => {
					const dashArray = getDashArray(chain.percentage)
					const rotation = (accumulatedPercentage / 100) * 360

					accumulatedPercentage += chain.percentage

					return (
						<circle
							key={index}
							r={radius}
							fill="transparent"
							stroke={getChainColor(chain.chainId)}
							strokeWidth="20"
							strokeDasharray={dashArray}
							transform={`rotate(${-90 + rotation})`}
							style={{
								transition: "stroke-dasharray 3s ease 0s"
							}}
						/>
					)
				})}
			</g>
		</svg>
	)
}
