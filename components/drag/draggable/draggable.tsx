import type { CSSProperties } from "react"
import React, { forwardRef } from "react"

import classNames from "classnames"

import type { DraggableSyntheticListeners } from "@dnd-kit/core"
import type { Transform } from "@dnd-kit/utilities"

import Plug from "@/components/canvas/plug/plug"
import { cn } from "@/lib/utils"

import styles from "./draggable.module.css"

export type DraggableProps = {
	id: string
	dragOverlay?: boolean
	dragging?: boolean
	listeners?: DraggableSyntheticListeners
	style?: CSSProperties
	buttonStyle?: CSSProperties
	transform?: Transform | null
	gridSize: number
}

export const Draggable = forwardRef<HTMLButtonElement, DraggableProps>(
	function Draggable(
		{
			id,
			dragOverlay,
			dragging,
			listeners,
			transform,
			style,
			buttonStyle,
			gridSize,
			...props
		},
		ref
	) {
		return (
			<div
				id={id}
				className={cn(
					"w-max items-start",
					classNames(
						styles.Draggable,
						dragOverlay && styles.dragOverlay,
						dragging && styles.dragging
					)
				)}
				style={
					{
						...style,
						"--translate-x": `${transform?.x ?? 0}px`,
						"--translate-y": `${transform?.y ?? 0}px`
					} as CSSProperties
				}
			>
				<button
					{...props}
					{...listeners}
					ref={ref}
					style={buttonStyle}
					className="absolute flex appearance-none content-center items-center border-none outline-none"
					aria-label="Draggable"
					data-cypress="draggable-item"
				>
					<Plug id={id} gridSize={gridSize} selecting={null}>
						{"[]"}
					</Plug>
				</button>
			</div>
		)
	}
)
