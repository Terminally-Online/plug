import { useMemo } from "react"

import { motion } from "framer-motion"
import { FileCog, FileTerminal } from "lucide-react"

import { formatNumber } from "@/lib/functions"
import { cn } from "@/lib/utils"

import { Header } from "../header"

export const SocketActivity = () => {
	// TODO: Implement the backend functionality for this.
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
										? "bg-gradient-to-tr from-[#00E100] to-[#A3F700]"
										: "bg-grayscale-100"
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

					<div className="flex flex-row gap-2">
						<p className="opacity-40">{start}</p>
						<p className="ml-auto opacity-40">{end}</p>
					</div>
				</div>

				<div className="flex flex-row gap-2">
					<div className="w-full rounded-lg bg-grayscale-100 p-4">
						<h4 className="text-2xl font-bold">
							{formatNumber(200)}
						</h4>
						<p className="opacity-40">Pending</p>
					</div>
					<div className="w-full rounded-lg bg-grayscale-100 p-4">
						<h4 className="text-2xl font-bold">
							{formatNumber(1900)}
						</h4>
						<p className="opacity-40">Run</p>
					</div>
				</div>
			</div>

			<Header
				size="md"
				icon={<FileCog size={14} className="opacity-60" />}
				label="Runs"
			/>

			<p className="opacity-60">
				Pending and completed Plug runs in your Socket will appear
				here...
			</p>
		</>
	)
}
