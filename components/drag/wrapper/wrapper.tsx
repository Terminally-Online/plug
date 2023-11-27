import React, { FC, memo, PropsWithChildren } from 'react'

import { cn } from '@/lib/utils'

export type WrapperProps = {
	center?: boolean
} & PropsWithChildren<React.HTMLAttributes<HTMLDivElement>>

export const Wrapper: FC<WrapperProps> = ({ children, center, className }) => {
	return (
		<div
			className={cn(
				'h-full flex w-full justify-start',
				center ? 'items-center' : 'items-start',
				className
			)}
		>
			{children}
		</div>
	)
}

export default memo(Wrapper)
