import { FC } from "react"

import { Clock } from "lucide-react"

import { Frame } from "@/components/app/frames/base"
import { Button } from "@/components/shared/buttons/button"
import { frequencies } from "@/lib"
import { columnByIndexAtom, isFrameAtom, useColumnActions } from "@/state/columns"
import { useAtom, useAtomValue } from "jotai"

export const RecurringFrame: FC<{ index: number }> = ({ index }) => {
	const [column] = useAtom(columnByIndexAtom(index))
	const frameKey = "recurring"
	const isFrame = useAtomValue(isFrameAtom)(column, frameKey)
	const { frame, schedule } = useColumnActions(index, frameKey)

	return (
		<Frame
			index={index}
			className="scrollbar-hide z-[2] max-h-[calc(100vh-80px)] overflow-y-auto"
			icon={<Clock size={18} className="opacity-40" />}
			label="Recurring Frequency"
			visible={isFrame}
			handleBack={() => frame("schedule")}
			hasOverlay={true}
		>
			<div className="flex flex-col gap-2">
				{frequencies.map(frequency => (
					<Button
						key={frequency.label}
						variant="secondary"
						className="py-4"
						onClick={() => {
							schedule({ date: column?.schedule?.date, repeats: frequency })
							frame("schedule")
						}}
					>
						{frequency.label}
					</Button>
				))}
			</div>
		</Frame>
	)
}
