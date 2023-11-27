'use client'

import { memo, useEffect } from 'react'
import type { CSSProperties, FC, PropsWithChildren } from 'react'
import React from 'react'

import { useDrag } from 'react-dnd'
import type { DragSourceMonitor } from 'react-dnd'
import { getEmptyImage } from 'react-dnd-html5-backend'

import { useDraggable } from '@dnd-kit/core'

import { ItemTypes } from '@/lib/constants'

function getStyles(isDragging: boolean): CSSProperties {
	return {
		opacity: isDragging ? 0 : 1,
		height: isDragging ? 0 : ''
	}
}

export type DraggableProps = {
	id: string
	// role: string
	// left: number
	// top: number
}

export const Draggable: FC<PropsWithChildren<DraggableProps>> = ({
	id,
	children
}) => {
	const { attributes, listeners, setNodeRef, transform } = useDraggable({
		id: 'draggable'
	})

	const style = transform
		? {
				transform: `translate3d(${transform.x}px, ${transform.y}px, 0)`
		  }
		: undefined

	return (
		<button
			ref={setNodeRef}
			id={id}
			style={style}
			{...listeners}
			{...attributes}
		>
			{children}
		</button>
	)
}

// export const Draggable: FC<PropsWithChildren<DraggableProps>> = memo(
// 	function Draggable({ id, role, left, top, children }) {
// 		// const [{ isDragging }, drag, preview] = useDrag(
// 		// 	() => ({
// 		// 		type: ItemTypes.Box,
// 		// 		item: { id, left, top, children },
// 		// 		collect: (monitor: DragSourceMonitor) => ({
// 		// 			isDragging: monitor.isDragging()
// 		// 		})
// 		// 	}),
// 		// 	[id, left, top, children]
// 		// )
// 		//
// 		// useEffect(() => {
// 		// 	preview(getEmptyImage(), { captureDraggingState: true })
// 		// }, [preview])
//
// 		return (
// 			<>{children}</>
// 			// <div ref={drag} style={getStyles(isDragging)} role={role}>
// 			// 	{children}
// 			// </div>
// 		)
// 	}
// )
