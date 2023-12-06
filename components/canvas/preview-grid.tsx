import type { FC } from "react"

import { Canvas } from "@prisma/client"

import CanvasPreview from "@/components/canvas/preview"

export type CanvasPreviewGridProps = {
	canvases: Array<Canvas>
}

export const CanvasPreviewGrid: FC<CanvasPreviewGridProps> = ({ canvases }) => (
	<div className="relative grid h-full grid-cols-3 gap-[1px] bg-stone-900">
		{canvases && canvases.length > 0 ? (
			<>
				{canvases.map(canvas => (
					<CanvasPreview key={canvas.id} canvas={canvas} />
				))}{" "}
			</>
		) : (
			<div className="col-span-3 row-span-4 flex h-full flex-col items-center justify-center gap-2 border-b-[1px] border-stone-950 text-white">
				<h1 className="text-2xl">No Results</h1>
				<p className="text-sm opacity-60">
					We couldn{"'"}t find any canvases matching your search.
				</p>
			</div>
		)}
	</div>
)

export default CanvasPreviewGrid
