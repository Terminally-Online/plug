import { FC } from "react"

import { Clock } from "lucide-react"

import { Frame } from "@/components/app/frames/base"
import { Button } from "@/components/shared/buttons/button"
import { frequencies } from "@/lib"
import { useColumnStore } from "@/state/columns"

export const RecurringFrame: FC<{ index: number }> = ({ index }) => {
	const { column, isFrame, handle } = useColumnStore(index, "recurring")

	return (
		<Frame
			index={index}
			className="scrollbar-hide z-[2] max-h-[calc(100vh-80px)] overflow-y-auto"
			icon={<Clock size={18} className="opacity-40" />}
			label="Recurring Frequency"
			visible={isFrame}
			handleBack={() => handle.frame("schedule")}
			hasOverlay={true}
		>
			<div className="flex flex-col gap-4">
				{frequencies.map(frequency => (
					<Button
						key={frequency.label}
						variant="secondary"
						className="py-4"
						onClick={() => {
							handle.schedule({ date: column?.schedule?.date, repeats: frequency })
							handle.frame("schedule")
						}}
					>
						{frequency.label}
					</Button>
				))}
			</div>
		</Frame>
	)
}
