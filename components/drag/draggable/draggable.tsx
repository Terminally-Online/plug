import type { CSSProperties, FC, PropsWithChildren } from "react"
import React, { forwardRef, useEffect, useMemo, useState } from "react"

import classNames from "classnames"

import type { DragEndEvent, DraggableSyntheticListeners } from "@dnd-kit/core"
import {
	DndContext,
	KeyboardSensor,
	MouseSensor,
	PointerActivationConstraint,
	TouchSensor,
	useDraggable,
	useSensor,
	useSensors
} from "@dnd-kit/core"
import { createSnapModifier } from "@dnd-kit/modifiers"
import type { Transform } from "@dnd-kit/utilities"
import { Component } from "@prisma/client"

import Plug from "@/components/canvas/plug/plug"
import { api } from "@/lib/api"
import { inBounds } from "@/lib/functions/math-utils"
import CanvasStore from "@/lib/store"
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
