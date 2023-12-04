import type { FC } from "react"
import { memo, useMemo } from "react"

import Link from "next/link"

import { Canvas } from "@prisma/client"

export type CanvasPreviewProps = {
	canvas: Canvas
}

export const CanvasPreview: FC<CanvasPreviewProps> = ({ canvas }) => {
	const { id, name, color, updatedAt } = canvas

	const durationDisplay = useMemo(() => {
		const duration = new Date().getTime() - updatedAt.getTime()

		const seconds = Math.floor(duration / 1000)
		const minutes = Math.floor(seconds / 60)
		const hours = Math.floor(minutes / 60)
		const days = Math.floor(hours / 24)

		if (days > 0) return `${days} day${days > 1 ? "s" : ""}`
		if (hours > 0) return `${hours} hour${hours > 1 ? "s" : ""}`
		if (minutes > 0) return `${minutes} minute${minutes > 1 ? "s" : ""}`
		if (seconds > 0) return `${seconds} second${seconds > 1 ? "s" : ""}`

		return "just now"
	}, [updatedAt])

	return (
		<Link
			href={{
				pathname: "/canvas/[id]",
				query: { id }
			}}
			className="flex min-h-[220px] flex-row items-end border-[1px] border-l-[0px] border-t-[0px] border-stone-950 bg-stone-900 p-4 text-white transition-all duration-200 ease-in-out hover:bg-stone-950 hover:text-white"
		>
			<div className="flex flex-col gap-2">
				<h1 className="flex flex-row items-center gap-4 text-lg font-bold">
					<div
						className="h-4 w-4 rounded-full border-[1px] border-stone-950"
						style={{ background: color }}
					/>
					{name}
				</h1>
				<p className="text-sm opacity-60">
					Edited {durationDisplay} ago
				</p>
			</div>
		</Link>
	)
}

export default memo(CanvasPreview)
