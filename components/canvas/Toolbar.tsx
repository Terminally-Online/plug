import { memo } from 'react'

export const Toolbar = () => (
	<div
		className="text-white fixed bottom-4 left-[50%]"
		style={{
			transform: 'translate(-50%, -50%)'
		}}
	>
		<p className="text-xs opacity-60">
			Tip: Double click anywhere to start a new plug
		</p>
	</div>
)

export default memo(Toolbar)
