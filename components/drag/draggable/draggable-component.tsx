import type { FC } from "react"

import { useDraggable } from "@dnd-kit/core"

import { inBounds } from "@/lib/functions/math-utils"
import CanvasStore from "@/lib/store"

import { Draggable } from "./draggable"

export type DraggableComponentProps = {
	id: string
	top: number
	left: number
	gridSize: number
}

export const DraggableComponent: FC<DraggableComponentProps> = ({
	id,
	top,
	left,
	gridSize
}) => {
	const { attributes, isDragging, listeners, setNodeRef, transform } =
		useDraggable({
			id
		})

	const buttonStyle = {
		marginLeft: gridSize - 20 + 1,
		marginTop: gridSize - 40 + 1,
		width: gridSize * 8 - 1,
		height: gridSize * 3 + 1
	}

	// ? If the element is not within the cameras bounds, do not render it.
	if (
		!inBounds(
			{
				left,
				top,
				width: buttonStyle.width,
				height: buttonStyle.height
			},
			{
				left: CanvasStore.screen.x,
				top: CanvasStore.screen.y,
				width: screen.width,
				height: screen.height
			}
		)
	)
		return null

	return (
		<Draggable
			id={id}
			ref={setNodeRef}
			dragging={isDragging}
			listeners={listeners}
			transform={transform}
			gridSize={gridSize}
			style={{
				position: "absolute",
				alignItems: "flex-start",
				top: top - CanvasStore.screen.y,
				left: left - CanvasStore.screen.x
			}}
			buttonStyle={buttonStyle}
			{...attributes}
		/>
	)
}
