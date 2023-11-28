import type { FC, PropsWithChildren } from 'react'
import { memo } from 'react'

import CanvasStore from '@/lib/store'

export type ScalerProps = PropsWithChildren<
	React.HTMLAttributes<HTMLDivElement>
>

export const Scaler: FC<ScalerProps> = ({ children, ...props }) => {
	return (
		<div
			className="w-screen h-full overscroll-none origin-top-left"
			style={{
				transform: `scale(${
					(CanvasStore.scale.x, CanvasStore.scale.y)
				})`
			}}
			{...props}
		>
			{children}
		</div>
	)
}

export default memo(Scaler)
