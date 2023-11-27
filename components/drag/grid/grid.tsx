import React, { FC, PropsWithChildren } from 'react'

import classNames from 'classnames'

import CanvasStore from '@/lib/store'
import { cn } from '@/lib/utils'

import styles from './Grid.module.css'

export interface GridProps {
	size: number
	step?: number
	onSizeChange(size: number): void
}

export const Grid: FC<PropsWithChildren<GridProps>> = ({ size, children }) => {
	return (
		<>
			<div
				className={cn(
					'bg-stone-900 absolute z-[-1] pointer-events-none',
					classNames(styles.Grid)
				)}
				style={
					{
						'--grid-size': `${size}px`,
						left: -1 * CanvasStore.screen.x,
						top: -1 * CanvasStore.screen.y,
						width: 10000,
						height: 10000
					} as React.CSSProperties
				}
			/>

			{children}
		</>
	)
}
