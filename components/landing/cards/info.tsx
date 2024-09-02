import { FC, HTMLAttributes, PropsWithChildren, ReactNode } from "react"

import { motion, MotionProps } from "framer-motion"

import { cn } from "@/lib/utils"

export const InfoCard: FC<
	HTMLAttributes<HTMLDivElement> &
		MotionProps &
		PropsWithChildren<{
			icon: ReactNode
			text: string | React.ReactNode
			description?: string
		}>
> = ({ children, icon, text, description, className, ...props }) => {
	return (
		<motion.div
			className={cn(
				"relative flex min-h-[240px] flex-row items-end gap-8 rounded-xl bg-grayscale-0 p-8",
				className
			)}
			initial={{ transform: "translateY(20px)", opacity: 0 }}
			whileInView={{
				transform: ["translateY(0px)", "translateY(20px)"],
				opacity: [0, 1]
			}}
			transition={{ duration: 0.3 }}
			{...props}
		>
			<div className="absolute bottom-0 left-0 right-0 top-0 overflow-hidden rounded-xl">{children}</div>

			<div className="flex-rows flex items-center gap-8">
				<div className="mb-auto mt-1">{icon}</div>
				<div className="z-[10] flex flex-col gap-2 font-bold">
					<h2 className="flex items-center gap-4 text-lg lg:text-2xl">{text}</h2>
					{description && <p className="max-w-[480px] text-black/40">{description}</p>}
				</div>
			</div>
		</motion.div>
	)
}

export default InfoCard
