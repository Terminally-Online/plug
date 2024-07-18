import { FC, PropsWithChildren } from "react"

import { motion, MotionProps } from "framer-motion"

import { greenGradientStyle } from "@/lib"

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
			className="relative flex flex-row items-center gap-8 rounded-lg bg-gradient-to-r from-[#d9d9d9]/0 to-[#D9D9D9]/20 px-[20px] py-[30px] lg:py-[60px] xl:px-[30px]"
			style={{
				background:
					"linear-gradient(30deg, rgba(217,217,217,.1), rgba(217,217,217,.1), rgba(217,217,217,.4))"
			}}
			{...props}
		>
			<div className="mx-auto">
				<div className="left-[20%] top-[25%] my-auto md:ml-[-10%] lg:absolute lg:top-0 lg:ml-8 xl:left-0">
					<p
						className="absolute mt-[-2%] text-[72px] font-bold blur-[20px] filter md:mt-0 lg:text-[148px]"
						style={{
							...greenGradientStyle
						}}
					>
						{index}
					</p>
					<p
						className="absolute mt-[-2%] text-[72px] font-bold md:mt-[-1%] lg:mt-0 lg:text-[148px]"
						style={{
							...greenGradientStyle
						}}
					>
						{index}
					</p>
				</div>
				<div className="relative z-[10] ml-16 flex flex-col gap-2 xl:ml-28">
					<h2 className="flex flex-wrap items-center gap-4 text-2xl font-bold">
						<span className="opacity-40">{children}</span>
						{title}
					</h2>
					<p className="text-black/65 sm:max-w-[380px] lg:max-w-[240px]">
						{description}
					</p>
				</div>
			</div>
		</motion.div>
	)
}

export default StepCard
