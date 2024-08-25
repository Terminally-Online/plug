import { FC, PropsWithChildren } from "react"

export const StatCard: FC<PropsWithChildren> = ({ children, ...props }) => {
	return (
		<div className="w-full rounded-[16px] bg-gradient-to-tr from-grayscale-100 to-grayscale-0 p-4" {...props}>
			{children}
		</div>
	)
}
