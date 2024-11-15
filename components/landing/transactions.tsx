import { useRef } from "react"

import { motion, useScroll, useTransform } from "framer-motion"

import { Execution, LandingContainer, Recurring, Scheduled } from "@/components"

export const Transactions = () => {
	const containerRef = useRef<HTMLDivElement>(null)
	const { scrollYProgress } = useScroll({
		target: containerRef,
		offset: ["start end", "end start"]
	})
	const pathLength = useTransform(scrollYProgress, [0.2, 0.5], [1, 0])

	return (
		<div
			ref={containerRef}
			className="relative z-[11] mb-[80px] h-full bg-plug-white xl:mb-[160px] xl:mt-[900px] 2xl:mt-[70vw]"
		>
			<svg
				viewBox="0 0 1827 976"
				fill="none"
				className="pointer-events-none absolute inset-0 z-[9] -ml-[15%] -mt-[9%] hidden xl:flex"
			>
				<g clip-path="url(#clip0_4611_7253)">
					<path
						d="M15.7501 1013.5C337.75 381 651.346 608.093 778.25 712.516C1184.75 1047 1590.08 925.892 1672.25 786.5C1789.56 587.5 1549.61 646.725 1737.75 371C1882.75 158.5 1727.75 38.9964 1630.25 39C1386.75 55 1417.75 519.515 1040.25 519.515C759.751 519.515 891.75 285.5 387.75 399C227.25 447.5 -155.101 454.5 -67.5012 134.5"
						stroke="url(#paint0_linear_4611_7253)"
						strokeWidth="60"
					/>
					<motion.path
						d="M15.7501 1013.5C337.75 381 651.346 608.093 778.25 712.516C1184.75 1047 1590.08 925.892 1672.25 786.5C1789.56 587.5 1549.61 646.725 1737.75 371C1882.75 158.5 1727.75 38.9964 1630.25 39C1386.75 55 1417.75 519.515 1040.25 519.515C759.751 519.515 891.75 285.5 387.75 399C227.25 447.5 -155.101 454.5 -67.5012 134.5"
						stroke="#FDFEF6"
						strokeWidth="60"
						stroke-dasharray="4 4"
						animate={{ strokeDashoffset: [0, 60] }}
						transition={{
							duration: 0.5,
							repeat: Infinity,
							ease: "linear"
						}}
					/>
				</g>
				<defs>
					<linearGradient
						id="paint0_linear_4611_7253"
						x1="1362.25"
						y1="232.5"
						x2="467.75"
						y2="504"
						gradientUnits="userSpaceOnUse"
					>
						<stop stop-color="#D2F38A" />
						<stop offset="1" stop-color="#385842" />
					</linearGradient>
					<clipPath id="clip0_4611_7253">
						<rect width="1827" height="976" fill="white" />
					</clipPath>
				</defs>
			</svg>

			<svg
				viewBox="0 0 1827 976"
				fill="none"
				className="pointer-events-none absolute inset-0 z-[9] -ml-[15%] -mt-[9%] hidden xl:flex"
			>
				<g clip-path="url(#clip0_4611_7253)">
					<motion.path
						style={{ pathLength }}
						d="M15.7501 1013.5C337.75 381 651.346 608.093 778.25 712.516C1184.75 1047 1590.08 925.892 1672.25 786.5C1789.56 587.5 1549.61 646.725 1737.75 371C1882.75 158.5 1727.75 38.9964 1630.25 39C1386.75 55 1417.75 519.515 1040.25 519.515C759.751 519.515 891.75 285.5 387.75 399C227.25 447.5 -155.101 454.5 -67.5012 134.5"
						stroke="url(#paint0_linear_4611_7253)"
						strokeWidth="60"
						strokeLinecap="round"
					/>
					<mask
						id="mask0_4611_7253"
						style={{ maskType: "alpha" }}
						maskUnits="userSpaceOnUse"
						x="1542"
						y="162"
						width="285"
						height="645"
					>
						<path d="M1542 162H1827V807H1542V162Z" fill="#D9D9D9" />
					</mask>
					<g mask="url(#mask0_4611_7253)">
						<motion.path
							style={{ pathLength }}
							d="M131 975.5C453 343 651.097 608.093 778.001 712.516C1184.5 1047 1589.83 925.892 1672 786.5C1789.31 587.5 1549.36 646.725 1737.5 371C1882.5 158.5 1727.5 38.9964 1630 39C1386.5 55 1417.5 519.515 1040 519.515C759.501 519.515 891.501 285.5 387.501 399C227.001 447.5 -72.0988 464.5 15.5012 144.5"
							stroke="url(#paint1_linear_4611_7253)"
							strokeWidth="60"
							strokeLinecap="round"
						/>
					</g>
				</g>
				<defs>
					<linearGradient
						id="paint0_linear_4611_7253"
						x1="1362.25"
						y1="232.5"
						x2="467.75"
						y2="504"
						gradientUnits="userSpaceOnUse"
					>
						<stop stop-color="#D2F38A" />
						<stop offset="1" stop-color="#385842" />
					</linearGradient>
					<linearGradient
						id="paint1_linear_4611_7253"
						x1="1362"
						y1="232.5"
						x2="467.5"
						y2="504"
						gradientUnits="userSpaceOnUse"
					>
						<stop stop-color="#D2F38A" />
						<stop offset="1" stop-color="#385842" />
					</linearGradient>
					<clipPath id="clip0_4611_7253">
						<rect width="1827" height="976" fill="white" />
					</clipPath>
				</defs>
			</svg>

			<svg
				viewBox="0 0 1827 976"
				fill="none"
				className="pointer-events-none absolute inset-0 z-[9999] -ml-[15%] -mt-[9%] hidden xl:flex"
			>
				<mask
					id="mask0_4612_25"
					style={{ maskType: "alpha" }}
					maskUnits="userSpaceOnUse"
					x="491"
					y="0"
					width="1336"
					height="976"
				>
					<path fill-rule="evenodd" clip-rule="evenodd" d="M1530 0H1827V976H491V555H1530V0Z" fill="#D9D9D9" />
				</mask>
				<g mask="url(#mask0_4612_25)">
					<motion.path
						style={{ pathLength }}
						d="M15.7501 1013.5C337.75 381 651.346 608.093 778.25 712.516C1184.75 1047 1590.08 925.892 1672.25 786.5C1789.56 587.5 1549.61 646.725 1737.75 371C1882.75 158.5 1727.75 38.9964 1630.25 39C1386.75 55 1417.75 519.515 1040.25 519.515C759.751 519.515 891.75 285.5 387.75 399C227.25 447.5 -155.101 454.5 -67.5012 134.5"
						stroke="url(#paint0_linear_4612_25)"
						strokeWidth="60"
						strokeLinecap="round"
					/>
				</g>
				<defs>
					<linearGradient
						id="paint0_linear_4612_25"
						x1="1362.25"
						y1="232.5"
						x2="467.75"
						y2="504"
						gradientUnits="userSpaceOnUse"
					>
						<stop stop-color="#D2F38A" />
						<stop offset="1" stop-color="#385842" />
					</linearGradient>
				</defs>
			</svg>

			<LandingContainer className="relative grid grid-cols-2 gap-2 xl:grid-cols-6 xl:grid-rows-2">
				<Scheduled />
				<Execution />
				<Recurring />
			</LandingContainer>
		</div>
	)
}
