import type { FC } from 'react'
import { memo } from 'react'

import { Canvas } from '@prisma/client'

import CanvasPreview from '@/components/canvas/preview'

export type CanvasPreviewGridProps = {
	canvases: Array<Canvas> | undefined
}

export const CanvasPreviewGrid: FC<CanvasPreviewGridProps> = ({ canvases }) => {
	if (!canvases) return null

	return (
		<div className="w-full h-full bg-stone-900">
			{canvases.length === 0 ? (
				<div className="h-full flex items-center justify-center">
					<h1 className="bg-stone-900 text-white text-xl font-bold text-center">
						No canvases found
					</h1>
				</div>
			) : (
				<div className="w-full h-full bg-stone-900 grid grid-cols-3 grid-rows-4">
					{canvases.map(canvas => (
						<CanvasPreview key={canvas.id} canvas={canvas} />
					))}
				</div>
			)}
		</div>
	)
}

export default memo(CanvasPreviewGrid)
