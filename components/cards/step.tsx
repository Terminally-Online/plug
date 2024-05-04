import type { FC, PropsWithChildren } from "react"

import { motion, MotionProps } from "framer-motion"

import { greenGradientStyle } from "@/lib/constants"

type Props = PropsWithChildren & {
	index: number
	title: string
	description: string
} & MotionProps

export const StepCard: FC<Props> = ({
	children,
	index,
	title,
	description,
	...props
}) => {
	return (
		<motion.div
			className="relative flex flex-row items-center gap-8 rounded-lg bg-gradient-to-r from-[#d9d9d9]/0 to-[#D9D9D9]/20 px-[20px] py-[30px] lg:px-[40px] lg:py-[60px]"
			style={{
				background:
					"linear-gradient(45deg, rgba(217,217,217,.1), rgba(217,217,217,.1), rgba(217,217,217,.4))"
			}}
			{...props}
		>
			<div className="absolute left-0 top-[25%] ml-4 lg:top-[2%] lg:ml-8">
				<p
					className="absolute text-[72px] font-bold blur-[20px] filter lg:text-[148px]"
					style={{
						...greenGradientStyle
					}}
				>
					{index}
				</p>
				<p
					className="absolute text-[72px] font-bold lg:text-[148px]"
					style={{
						...greenGradientStyle
					}}
				>
					{index}
				</p>
			</div>
			<div className="relative z-[10] ml-12 flex flex-col gap-2 lg:ml-24">
				<h2 className="flex flex-wrap items-center gap-4 text-2xl font-bold">
					<span className="opacity-40">{children}</span>
					{title}
				</h2>
				<p className="text-black/65">{description}</p>
			</div>
		</motion.div>
	)
}

export default StepCard
