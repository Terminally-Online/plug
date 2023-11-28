import React, { forwardRef } from 'react'

import classNames from 'classnames'

import type { DraggableSyntheticListeners } from '@dnd-kit/core'
import type { Transform } from '@dnd-kit/utilities'

import { cn } from '@/lib/utils'

import { Handle } from '../item/handle'
import {
	draggable,
	draggableHorizontal,
	draggableVertical
} from './draggable-svg'
import styles from './draggable.module.css'

export enum Axis {
	All,
	Vertical,
	Horizontal
}

interface Props {
	id: string
	axis?: Axis
	dragOverlay?: boolean
	dragging?: boolean
	handle?: boolean
	listeners?: DraggableSyntheticListeners
	style?: React.CSSProperties
	buttonStyle?: React.CSSProperties
	transform?: Transform | null
}

export const Draggable = forwardRef<HTMLButtonElement, Props>(
	function Draggable(
		{
			id,
			axis,
			dragOverlay,
			dragging,
			handle,
			listeners,
			transform,
			style,
			buttonStyle,
			...props
		},
		ref
	) {
		return (
			<div
				id={id}
				className={cn(
					classNames(
						styles.Draggable,
						dragOverlay && styles.dragOverlay,
						dragging && styles.dragging,
						handle && styles.handle
					)
				)}
				style={
					{
						...style,
						'--translate-x': `${transform?.x ?? 0}px`,
						'--translate-y': `${transform?.y ?? 0}px`
					} as React.CSSProperties
				}
			>
				<button
					{...props}
					aria-label="Draggable"
					data-cypress="draggable-item"
					{...(handle ? {} : listeners)}
					tabIndex={handle ? -1 : undefined}
					ref={ref}
					style={buttonStyle}
				>
					{axis === Axis.Vertical
						? draggableVertical
						: axis === Axis.Horizontal
						  ? draggableHorizontal
						  : draggable}
					{handle ? <Handle {...(handle ? listeners : {})} /> : null}
				</button>
			</div>
		)
	}
)
