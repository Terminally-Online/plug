import { FC } from "react"

export const Dot: FC<{
	length: number
	color: string
	cx?: number
	cy?: number
	index?: number
}> = ({ length, color, cx, cy, index }) => {
	if (index === length - 1) {
		return (
			<g>
				<circle cx={cx} cy={cy} r={12} fill={color}>
					<animate
						attributeName="r"
						values="8;5;8"
						dur="1s"
						repeatCount="indefinite"
						keyTimes="0;0.5;1"
						keySplines="0.1 0.8 0.2 1; 0.1 0.8 0.2 1"
						calcMode="spline"
						begin="0s;animate2.end+1s"
					/>
					<animate
						id="animate2"
						attributeName="opacity"
						values="0;0.5;0"
						dur="1s"
						repeatCount="indefinite"
						keyTimes="0;0.5;1"
						keySplines="0.4 0 0.6 1; 0.4 0 0.6 1"
						calcMode="spline"
					/>
				</circle>
				<circle cx={cx} cy={cy} r={5} fill={color} />
			</g>
		)
	}

	return null
}

export const ActiveDot: FC<{
	color: string
	cx?: number
	cy?: number
}> = ({ color, cx, cy }) => <circle cx={cx} cy={cy} r={5} fill={color} />
