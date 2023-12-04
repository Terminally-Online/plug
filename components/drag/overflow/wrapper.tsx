import React from "react"

interface OverflowWrapperProps {
	children: React.ReactNode
}

export function OverflowWrapper({ children }: OverflowWrapperProps) {
	return (
		<div className="absolute bottom-0 left-0 right-0 top-0 overflow-hidden">
			{children}
		</div>
	)
}
