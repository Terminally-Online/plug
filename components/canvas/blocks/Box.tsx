import type { FC, PropsWithChildren } from 'react'
import { memo } from 'react'

export interface BoxProps {
	preview?: boolean
}

export const Box: FC<PropsWithChildren<BoxProps>> = memo(function Box({
	children,
	preview
}) {
	return (
		<div
			className="bg-white dark:bg-black cursor-move p-2 px-4 border-[1px] border-gray-200"
			role={preview ? 'BoxPreview' : 'Box'}
		>
			{children}
		</div>
	)
})
