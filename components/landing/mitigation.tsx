import { FC, useEffect } from "react"

import { animate, motion, useMotionValue, useTransform } from "framer-motion"

export const Mitigation: FC = () => {
	const patch = useMotionValue(0)
	const patchRounded = useTransform(patch, latest => Math.round(latest))

	useEffect(() => {
		const controls = animate(
			patch,
			[40, 31, 38, 22, 30, 26, 32, 16, 10, 40, 60, 80, 90, 100],
			{
				duration: 4,
				delay: 1,
				repeat: Infinity,
				repeatDelay: 2,
				repeatType: "reverse"
			}
		)
		return controls.stop
	}, [patch])

	return (
		<>
			<div className="mb-[-120px] flex flex-col gap-2 p-10">
				<p className="opacity-60">Health Factor</p>
				<motion.p
					className="text-[32px] font-bold"
					animate={{
						color: [
							"#FF5154",
							"#FF5154",
							"#00E100",
							"#FF5154",
							"#00E100",
							"#FF5154",
							"#00E100",
							"#FF5154",
							"#FF5154",
							"#00E100",
							"#00E100",
							"#00E100",
							"#00E100",
							"#00E100"
						]
					}}
					transition={{
						duration: 4,
						delay: 1,
						ease: "easeInOut",
						repeat: Infinity,
						repeatDelay: 2,
						repeatType: "reverse"
					}}
				>
					<motion.span>{patchRounded}</motion.span>%
				</motion.p>
			</div>

			<svg
				viewBox="0 0 535 201"
				fill="none"
				xmlns="http://www.w3.org/2000/svg"
				className="w-[100%]"
			>
				<motion.path
					d="M1.5 106C1.5 106 27.3926 114.059 44.375 113.438C52.6986 113.133 57.2468 109.378 65.5 110.5C74.8249 111.768 79.5155 115.514 87.25 120.875C96.3002 127.148 96.0692 139.668 107 141C117.225 142.246 120.238 131.203 130.125 128.312C138.184 125.956 143.193 124.777 151.5 126C160.621 127.343 163.956 133.961 173 135.75C190.447 139.202 199.538 120.627 217 124C237.634 127.986 236.517 155.299 257 160C274.057 163.915 283.528 150.49 301 151.5C318.816 152.53 344.5 165.5 344.5 165.5L351 145.5L360.5 193C369.652 166.601 392.002 90.9261 415 56C446.45 8.23862 483.863 3.61732 534 2"
					stroke="#D9D9D9"
					stroke-width="4"
					initial={{ strokeDasharray: 750 }}
					animate={{ strokeDashoffset: [750, 0] }}
					transition={{
						duration: 4,
						delay: 1,
						ease: "easeInOut",
						repeat: Infinity,
						repeatDelay: 2,
						repeatType: "reverse"
					}}
				/>
			</svg>
		</>
	)
}

export default Mitigation
