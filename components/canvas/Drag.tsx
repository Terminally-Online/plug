'use client'

import type { CSSProperties, FC } from 'react'

import type { XYCoord } from 'react-dnd'
import { useDragLayer } from 'react-dnd'

import { snapToGrid } from '@/lib/functions/snap-to-grid'

import { Preview } from './Preview'

const layerStyles: CSSProperties = {
	position: 'fixed',
	pointerEvents: 'none',
	zIndex: 100,
	left: 0,
	top: 0,
	width: '100%',
	height: '100%'
}

function getItemStyles(
	initialOffset: XYCoord | null,
	currentOffset: XYCoord | null,
	isSnapToGrid: boolean
) {
	if (!initialOffset || !currentOffset) {
		return {
			display: 'none'
		}
	}

	let { x, y } = currentOffset

	if (isSnapToGrid) {
		x -= initialOffset.x
		y -= initialOffset.y
		;[x, y] = snapToGrid(x, y)
		x += initialOffset.x
		y += initialOffset.y
	}

	const transform = `translate(${x}px, ${y}px)`
	return {
		transform,
		WebkitTransform: transform
	}
}

export interface CustomDragLayerProps {
	snapToGrid?: boolean
}

export const Drag: FC<CustomDragLayerProps> = props => {
	// const { isdragging, item, initialoffset, currentoffset } = usedraglayer(
	// 	monitor => ({
	// 		item: monitor.getitem(),
	// 		itemtype: monitor.getitemtype(),
	// 		initialoffset: monitor.getinitialsourceclientoffset(),
	// 		currentoffset: monitor.getsourceclientoffset(),
	// 		isdragging: monitor.isdragging()
	// 	})
	// )
	//
	// function renderItem() {
	// 	if (!item || !item.children) return null
	//
	// 	return <Preview>{item.children}</Preview>
	// }
	//
	// if (!isDragging) {
	// 	return null
	// }

	// return (
	// 	<div style={layerStyles}>
	// 		<div
	// 			style={getItemStyles(
	// 				initialOffset,
	// 				currentOffset,
	// 				props.snapToGrid || true
	// 			)}
	// 		>
	// 			{renderItem()}
	// 		</div>
	// 	</div>
	// )
	//
	return <>Drag</>
}
