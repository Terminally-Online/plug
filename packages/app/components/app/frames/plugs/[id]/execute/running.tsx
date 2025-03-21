import { FC, useEffect } from "react"

import { LoaderCircle } from "lucide-react"

import { useAtom, useAtomValue } from "jotai"

import { Frame } from "@/components/app/frames/base"
import { columnByIndexAtom, isFrameAtom, useColumnActions } from "@/state/columns"
import { plugByIdAtom } from "@/state/plugs"

export const RunningFrame: FC<{ index: number; item: string }> = ({ index, item }) => {
	const [column] = useAtom(columnByIndexAtom(index))
	const frameKey = "running"
	const isFrame = useAtomValue(isFrameAtom)(column, frameKey)
	const { frame } = useColumnActions(index, frameKey)

	const [plug] = useAtom(plugByIdAtom(item))

	// TODO: We un-implemented this when beginning to store frames on columns.
	const prevFrame = "NOT_IMPLEMENTED" as string

	const label = prevFrame ? (prevFrame === "schedule" ? "Signing Intent" : "Building Intent...") : ""

	useEffect(() => {
		if (!isFrame) return

		const timeout = setTimeout(() => frame("ran"), 2500)
		return () => clearTimeout(timeout)
	}, [isFrame, frame])

	if (!plug) return null

	return (
		<Frame
			index={index}
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
						<span className="opacity-60">is currently running.</span>
					</p>
				) : (
					<>
						<p className="leading-6">
							<span className="opacity-60">Your </span>{" "}
							<span
								className="rounded-lg bg-gradient-to-tr px-2 py-1 font-bold text-plug-green"
								style={{
									background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`
								}}
							>
								{plug.name}
							</span>{" "}
							<span className="opacity-60">intent is being packaged up and distributed to solvers.</span>
						</p>
						{/* <p className="leading-6">
							<span className="opacity-60">Go ahead and schedule the execution of</span>{" "}
							<span
								className="rounded-lg bg-gradient-to-tr px-2 py-1 font-bold text-plug-green"
								style={{
									background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`
								}}
							>
								{plug.name}
							</span>{" "}
							<span className="opacity-60">by signing the intent in your wallet now.</span>
						</p> */}
					</>
				)}
			</div>
		</Frame>
	)
}
