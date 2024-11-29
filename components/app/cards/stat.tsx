import { FC, PropsWithChildren } from "react"

export const StatCard: FC<PropsWithChildren> = ({ children, ...props }) => {
	return (
		<div className="w-full rounded-[16px] bg-gradient-to-tr from-plug-green/10 to-plug-green/5 p-4" {...props}>
			{children}
		</div>
	)
}
