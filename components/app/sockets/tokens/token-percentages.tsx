import { FC } from "react"

import { RouterOutputs } from "@/server/client"

import { chains, getChainId } from "@/lib"

const SIZE = 14
const RADIUS = 50
const CIRCUMFERENCE = 2 * Math.PI * RADIUS
const VIEWBOX = 120

export const SocketTokenPercentages: FC<{
	implementations: RouterOutputs["socket"]["balances"]["positions"]["tokens"][number]["implementations"]
}> = ({ implementations }) => {
	let accumulatedPercentage = 0

	const getDashArray = (percentage: number) => {
		const filledLength = (percentage / 100) * CIRCUMFERENCE
		const emptyLength = CIRCUMFERENCE - filledLength
		return `${filledLength} ${emptyLength}`
	}

	return (
		<svg width={SIZE} height={SIZE} viewBox={`0 0 ${VIEWBOX} ${VIEWBOX}`}>
			<g transform={`translate(${VIEWBOX / 2},${VIEWBOX / 2})`}>
				{implementations.map((implementation, index) => {
					const dashArray = getDashArray(implementation.percentage)
					const rotation = (accumulatedPercentage / 100) * 360
					accumulatedPercentage += implementation.percentage

					return (
						<circle
							key={index}
							r={RADIUS}
							fill="transparent"
							stroke={chains[getChainId(implementation.chain)].color}
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
