import { useEffect, useMemo } from "react"

import { animate, motion, useMotionValue, useTransform } from "framer-motion"
import { FileCog, FileTerminal } from "lucide-react"

import { Counter } from "@/components/utils/Counter"
import { formatNumber } from "@/lib/functions"
import { cn } from "@/lib/utils"

import { Header } from "../header"
import { ActivityList } from "./activity-list"

export const SocketActivity = () => {
	const activity = Array.from({ length: 7 }, () => Math.random())

	const [start, end] = useMemo(
		() => [
			new Date(Date.now() - 7 * 24 * 60 * 60 * 1000).toLocaleDateString(
				"en-US",
				{
					month: "numeric",
					day: "numeric"
				}
			),
			new Date().toLocaleDateString("en-US", {
				month: "numeric",
				day: "numeric"
			})
		],
		[]
	)

	return (
		<>
			<div className="flex flex-col gap-4">
				<div className="flex flex-col gap-2">
					<div className="ml-auto mt-auto flex min-h-[124px] w-full flex-row gap-2">
						{activity.map((activity, index) => (
							<motion.div
								key={index}
								className={cn(
									"mt-auto w-full rounded-md",
									index === 6
										? "bg-gradient-to-tr from-plug-green to-plug-yellow"
										: "bg-grayscale-0"
								)}
								initial={{
									height: "10px"
								}}
								animate={{
									height: `${Math.max(10, activity * 96)}px`
								}}
							/>
						))}
					</div>

					<div className="flex flex-row gap-2 text-sm tabular-nums">
						<p className="opacity-60">{start}</p>
						<p className="ml-auto opacity-60">{end}</p>
					</div>
				</div>

				<div className="flex flex-row gap-2">
					<div className="w-full rounded-[16px] bg-grayscale-0 p-4">
						<Counter
							className="mr-auto w-max text-2xl font-bold"
							count={201}
							decimals={0}
						/>
						<p className="font-bold opacity-40">Pending</p>
					</div>
					<div className="w-full rounded-[16px] bg-grayscale-0 p-4">
						<Counter
							className="mr-auto w-max text-2xl font-bold"
							count={9351}
							decimals={0}
						/>
						<p className="font-bold opacity-40">Run</p>
					</div>
				</div>
			</div>

			<Header size="md" icon={<FileCog size={14} />} label="Runs" />

			<ActivityList />
		</>
	)
}
