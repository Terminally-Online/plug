import { FC, useEffect } from "react"

import { animate, motion, useMotionValue, useTransform } from "framer-motion"

import { greenGradientStyle } from "@/lib/constants"

export const Version: FC = () => {
	const major = useMotionValue(0)
	const majorRounded = useTransform(major, latest => Math.round(latest))

	const minor = useMotionValue(0)
	const minorRounded = useTransform(minor, latest => Math.round(latest))

	const patch = useMotionValue(0)
	const patchRounded = useTransform(patch, latest => Math.round(latest))

	const duration = 3
	const delay = 2
	const repeatType = "reverse"
	const repeatDelay = 1

	useEffect(() => {
		const controls = animate(major, 9, {
			duration,
			delay,
			repeat: Infinity,
			repeatType,
			repeatDelay
		})
		return controls.stop
	}, [])

	useEffect(() => {
		const controls = animate(minor, 45, {
			duration,
			delay,
			repeat: Infinity,
			repeatType,
			repeatDelay
		})
		return controls.stop
	}, [])

	useEffect(() => {
		const controls = animate(patch, 23, {
			duration,
			delay,
			repeat: Infinity,
			repeatType,
			repeatDelay
		})
		return controls.stop
	}, [])

	return (
		<div className="w-full px-8 pt-8 text-right">
			<h3 className="ml-auto text-[72px] font-bold tabular-nums text-[rgba(217,217,217,0.4)]">
				v
				<span>
					<motion.span style={{ ...greenGradientStyle }}>
						{majorRounded}
					</motion.span>
					.
					<motion.span style={{ ...greenGradientStyle }}>
						{minorRounded}
					</motion.span>
					.
					<motion.span style={{ ...greenGradientStyle }}>
						{patchRounded}
					</motion.span>
				</span>
			</h3>
		</div>
	)
}

export default Version
