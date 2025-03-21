import { FC } from "react"

import { ShieldX } from "lucide-react"

import { useAtom, useAtomValue } from "jotai"

import { Frame } from "@/components/app/frames/base"
import { columnByIndexAtom, isFrameAtom } from "@/state/columns"

export const ErrorFrame: FC<{ index: number }> = ({ index }) => {
	const [column] = useAtom(columnByIndexAtom(index))
	const frameKey = "error"
	const isFrame = useAtomValue(isFrameAtom)(column, frameKey)

	if (!column) return null

	return (
		<Frame
			index={index}
			className="z-[2]"
			icon={<ShieldX size={18} className="opacity-40" />}
			label="Error Encountered"
			visible={isFrame}
		>
			<div className="flex flex-col gap-4 font-bold">
				<p className="mx-auto max-w-[380px] opacity-40">
					We encountered an error while processing your request. Our team has been automatically notified.
				</p>
			</div>
		</Frame>
	)
}
