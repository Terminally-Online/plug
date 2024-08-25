import { FC, HTMLAttributes, useState } from "react"

import { motion } from "framer-motion"

import { Counter } from "@/components/shared"
import { cn } from "@/lib"

export const SocketEarningsChartItem: FC<
	HTMLAttributes<HTMLDivElement> & {
		padded: boolean
		forks: number
		runs: number
		active: boolean
	}
> = ({ padded, forks, runs, active, ...props }) => {
	return (
		<div className={cn("mt-auto flex h-full w-full flex-col", padded && "pl-2")} {...props}>
			<div className={cn("mx-auto flex-col transition-all duration-200 ease-in-out", active ? "opacity-100" : "opacity-0")}>
				<p className="mx-auto text-sm font-bold opacity-40">
					<Counter count={active ? 912 : 0} />
				</p>
				<div className="mx-auto h-3 w-1 bg-grayscale-0" />
			</div>

			<motion.div
				className="rounded-t-md bg-gradient-to-tr transition-all duration-200 ease-in-out"
				initial={{
					height: "10px"
				}}
				animate={{
					height: `${Math.max(10, forks * 96)}px`,
					background: active ? "linear-gradient(to right, #00E100, #A3F700)" : "linear-gradient(to right, #EBECEC, #f8f8f8)"
				}}
			/>

			<motion.div
				className="mt-1 rounded-b-md bg-gradient-to-tr transition-all duration-200 ease-in-out"
				initial={{
					height: "10px"
				}}
				animate={{
					height: `${Math.max(10, runs * 96)}px`,
					background: active ? "linear-gradient(to right, #FFA800, #FAFF00)" : "linear-gradient(to right, #EBECEC, #f8f8f8)"
				}}
			/>
		</div>
	)
}
