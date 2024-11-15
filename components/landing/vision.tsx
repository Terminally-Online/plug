import { FC, useRef } from "react"

import { motion, useScroll, useTransform } from "framer-motion"

import { BookProfit, LandingContainer, Routing } from "@/components"

import { Underperforming } from "./underperforming"

export const Vision: FC = () => {
	const containerRef = useRef<HTMLDivElement>(null)
	const { scrollYProgress } = useScroll({
		target: containerRef,
		offset: ["start end", "end start"]
	})
	const pathLength = useTransform(scrollYProgress, [0.2, 0.7], [0, 1])

	return (
		<div className="relative z-[11] mb-[80px] h-full bg-plug-white" ref={containerRef}>
			<svg
				viewBox="0 0 1827 976"
				fill="none"
				className="pointer-events-none absolute inset-0 z-[2] -mr-[10%] -mt-[10%] hidden xl:flex"
			>
				<g clip-path="url(#clip0_4612_71)">
					<path
						d="M1798.8 -0.000377674C1476.8 632.5 1272.62 374.395 1151.8 262.984C866.618 -0.000367615 664.618 129.999 556.618 246.499C399.571 415.907 189.936 272.5 74.6184 441.5C-40.6992 610.5 86 932 327 882.5C550.338 836.628 509.184 536.485 886.682 536.485C1167.18 536.485 1126 737 1539.18 657C1699.68 608.5 2001.9 511 1914.3 831"
						stroke="url(#paint0_linear_4612_71)"
						stroke-width="60"
					/>
					<motion.path
						d="M1798.8 -0.000377674C1476.8 632.5 1272.62 374.395 1151.8 262.984C866.618 -0.000367615 664.618 129.999 556.618 246.499C399.571 415.907 189.936 272.5 74.6184 441.5C-40.6992 610.5 86 932 327 882.5C550.338 836.628 509.184 536.485 886.682 536.485C1167.18 536.485 1126 737 1539.18 657C1699.68 608.5 2001.9 511 1914.3 831"
						stroke="#FEFFF7"
						stroke-width="60"
						stroke-dasharray="4 4"
						animate={{ strokeDashoffset: [60, 0] }}
						transition={{
							duration: 0.5,
							repeat: Infinity,
							ease: "linear"
						}}
					/>
				</g>
				<defs>
					<linearGradient
						id="paint0_linear_4612_71"
						x1="567.805"
						y1="743"
						x2="1462.3"
						y2="471.5"
						gradientUnits="userSpaceOnUse"
					>
						<stop stop-color="#D2F38A" />
						<stop offset="1" stop-color="#385842" />
					</linearGradient>
					<clipPath id="clip0_4612_71">
						<rect width="1827" height="976" fill="white" />
					</clipPath>
				</defs>
			</svg>

			<svg
				viewBox="0 0 1827 976"
				fill="none"
				xmlns="http://www.w3.org/2000/svg"
				className="pointer-events-none absolute inset-0 z-[2] -mr-[10%] -mt-[10%] hidden xl:flex"
			>
				<g clip-path="url(#clip0_4612_71)">
					<motion.path
						d="M1798.8 -0.000377674C1476.8 632.5 1272.62 374.395 1151.8 262.984C866.618 -0.000367615 664.618 129.999 556.618 246.499C399.571 415.907 189.936 272.5 74.6184 441.5C-40.6992 610.5 86 932 327 882.5C550.338 836.628 509.184 536.485 886.682 536.485C1167.18 536.485 1126 737 1539.18 657C1699.68 608.5 2001.9 511 1914.3 831"
						stroke="url(#paint0_linear_4612_71)"
						stroke-width="60"
						style={{ pathLength }}
						strokeLinecap="round"
					/>
				</g>
				<defs>
					<linearGradient
						id="paint0_linear_4612_71"
						x1="567.805"
						y1="743"
						x2="1462.3"
						y2="471.5"
						gradientUnits="userSpaceOnUse"
					>
						<stop stop-color="#D2F38A" />
						<stop offset="1" stop-color="#385842" />
					</linearGradient>
					<clipPath id="clip0_4612_71">
						<rect width="1827" height="976" fill="white" />
					</clipPath>
				</defs>
			</svg>

			<svg
				viewBox="0 0 1827 976"
				fill="none"
				className="pointer-events-none absolute inset-0 z-[99] -mr-[10%] -mt-[10%] hidden xl:flex"
			>
				<mask
					id="mask0_4612_101"
					style={{ maskType: "alpha" }}
					maskUnits="userSpaceOnUse"
					x="0"
					y="59"
					width="844"
					height="858"
				>
					<path
						fill-rule="evenodd"
						clip-rule="evenodd"
						d="M843.998 59H0V917H370.95V488H642.357V325H843.998V59Z"
						fill="#D9D9D9"
					/>
				</mask>
				<g mask="url(#mask0_4612_101)">
					<motion.path
						d="M1798.8 -0.000377674C1476.8 632.5 1272.62 374.395 1151.8 262.984C866.618 -0.000367615 664.618 129.999 556.618 246.499C399.571 415.907 189.936 272.5 74.6184 441.5C-40.6992 610.5 86 932 327 882.5C550.338 836.628 509.184 536.485 886.682 536.485C1167.18 536.485 1126 737 1539.18 657C1699.68 608.5 2001.9 511 1914.3 831"
						stroke="url(#paint0_linear_4612_101)"
						stroke-width="60"
						style={{ pathLength }}
						strokeLinecap="round"
					/>
				</g>
				<defs>
					<linearGradient
						id="paint0_linear_4612_101"
						x1="567.805"
						y1="743"
						x2="1462.3"
						y2="471.5"
						gradientUnits="userSpaceOnUse"
					>
						<stop stop-color="#D2F38A" />
						<stop offset="1" stop-color="#385842" />
					</linearGradient>
				</defs>
			</svg>

			<LandingContainer className="relative z-[9] mb-[80px] grid grid-cols-2 gap-2 xl:grid-cols-6 xl:grid-rows-2">
				<BookProfit />
				<Routing />
				<Underperforming />
			</LandingContainer>
		</div>
	)
}
