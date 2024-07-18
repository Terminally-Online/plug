import { useEffect } from "react"

import { LoaderCircle } from "lucide-react"

import { Frame } from "@/components"
import { useFrame, usePlugs } from "@/contexts"

export const RunningFrame = () => {
	const { frameVisible, handleFrameVisible } = useFrame()
	const { plug } = usePlugs()

	const isFrame = frameVisible
		? frameVisible.split("-")[0] === "running"
		: false

	const prevFrame = frameVisible ? frameVisible.split("-")[1] : undefined

	const label = prevFrame
		? prevFrame === "schedule"
			? "Signing Intent"
			: "Transaction Running"
		: ""

	useEffect(() => {
		if (isFrame) {
			const timeout = setTimeout(
				() => handleFrameVisible(`ran-${prevFrame}`),
				3000
			)

			return () => clearTimeout(timeout)
		}
	}, [isFrame, prevFrame, handleFrameVisible])

	if (!plug) return null

	return (
		<Frame
			className="z-[2]"
			icon={<LoaderCircle size={18} className="animate-spin" />}
			label={label}
			visible={isFrame}
		>
			<div className="flex flex-col gap-8">
				{prevFrame === "run" ? (
					<p className="leading-6">
						<span className="opacity-60">The execution of</span>{" "}
						<span
							className="rounded-lg bg-gradient-to-tr px-2 py-1 font-bold text-plug-green"
							style={{
								background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`
							}}
						>
							{plug.name}
						</span>{" "}
						<span className="opacity-60">
							is currently running.
						</span>
					</p>
				) : (
					<>
						<p className="leading-6">
							<span className="opacity-60">
								Go ahead and schedule the execution of
							</span>{" "}
							<span
								className="rounded-lg bg-gradient-to-tr px-2 py-1 font-bold text-plug-green"
								style={{
									background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`
								}}
							>
								{plug.name}
							</span>{" "}
							<span className="opacity-60">
								by signing the intent in your wallet now.
							</span>
						</p>
					</>
				)}
			</div>
		</Frame>
	)
}
