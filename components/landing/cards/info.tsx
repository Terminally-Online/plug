import { FC, PropsWithChildren } from "react"

import { motion, MotionProps } from "framer-motion"

import { cn } from "@/lib/utils"

type Props = PropsWithChildren & {
	text: string | React.ReactNode
	description?: string
} & React.HTMLProps<HTMLDivElement> &
	MotionProps

export const InfoCard: FC<Props> = ({
	children,
	text,
	description,
	className,
	...props
}) => {
	const base =
		"relative flex flex-row items-center gap-8 rounded-xl bg-[#FBFBFB] p-8 items-end min-h-[240px]"

	return (
		<motion.div className={cn(base, className)} {...props}>
			<div className="absolute bottom-0 left-0 right-0 top-0 overflow-hidden rounded-xl">
				{children}
			</div>

			<div className="z-[10] flex flex-col gap-2">
				<h2 className="flex flex-wrap items-center gap-4 text-lg font-bold lg:text-2xl">
					{text}
				</h2>
				{description && (
					<p className="max-w-[640px] text-black/65">{description}</p>
				)}
			</div>
		</motion.div>
	)
}

export default InfoCard
