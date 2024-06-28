import { useEffect, useState } from "react"

import { motion } from "framer-motion"
import { ArrowRight, CalendarPlus, ChevronDown } from "lucide-react"

import { Button } from "@/components/buttons"
import { Checkbox } from "@/components/inputs"
import { useFrame, usePlugs, useSockets } from "@/contexts"
import { cn } from "@/lib"

import { Frame } from "../../../base"

export const ScheduleFrame = () => {
	const { frameVisible, handleFrameVisible } = useFrame()
	const { sockets } = useSockets()
	const { chainsAvailable } = usePlugs()

	const [advanced, setAdvanced] = useState(false)
	const [allDay, setAllDay] = useState(false)

	const isFrame = frameVisible === "schedule"

	const handleBack =
		chainsAvailable.length === 1
			? sockets && sockets.length === 1
				? undefined
				: () => handleFrameVisible("socket-schedule")
			: () => handleFrameVisible("chain-schedule")

	useEffect(() => setAdvanced(false), [frameVisible])

	return (
		<Frame
			className="z-[2]"
			handleBack={handleBack}
			icon={<CalendarPlus size={18} className="opacity-60" />}
			label="Choose Schedule"
			visible={isFrame}
			hasOverlay={true}
		>
			<div className="flex flex-col gap-4">
				<div className="flex w-full flex-row items-center justify-between">
					<button
						className="rounded-md bg-grayscale-100 p-1 px-2 font-bold text-plug-green"
						style={{
							background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`
						}}
					>
						Wed Apr 24
					</button>
					<ArrowRight size={14} className="opacity-60" />
					<button
						className="rounded-md bg-grayscale-100 p-1 px-2 font-bold text-plug-green"
						style={{
							background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`
						}}
					>
						Wed Apr 24
					</button>
				</div>

				<button
					className="group mt-4 flex w-full cursor-pointer flex-row gap-2 font-bold opacity-40"
					onClick={() => setAdvanced(!advanced)}
				>
					<span className="mr-auto">Advanced</span>
					<Button
						variant="secondary"
						className={cn(
							"p-1 transition-all duration-200 ease-in-out group-hover:bg-grayscale-100",
							advanced && "bg-grayscale-100"
						)}
						onClick={() => setAdvanced(!advanced)}
					>
						<motion.div
							animate={{ rotate: advanced ? 180 : 0 }}
							transition={{ duration: 0.2 }}
						>
							<ChevronDown size={14} />
						</motion.div>
					</Button>
				</button>

				{advanced && (
					<>
						<div className="flex flex-row items-center gap-2">
							<p className="mr-auto font-bold">All-day</p>

							<Checkbox
								checked={allDay}
								handleChange={setAllDay}
							/>
						</div>
					</>
				)}

				<Button
					className="mt-4 w-full"
					onClick={() => handleFrameVisible("run-schedule")}
				>
					Next
				</Button>
			</div>
		</Frame>
	)
}
