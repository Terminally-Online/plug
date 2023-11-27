import React from 'react'

import classNames from 'classnames'

import { cn } from '@/lib/utils'

import styles from './Grid.module.css'

export interface GridProps {
	size: number
	step?: number
	onSizeChange(size: number): void
}

export function Grid({ size }: GridProps) {
	return (
		<div
			className={cn(
				'bg-stone-900 w-full h-full z-[-1] pointer-events-none',
				classNames(styles.Grid)
			)}
			style={
				{
					'--grid-size': `${size}px`
				} as React.CSSProperties
			}
		/>
	)
}
