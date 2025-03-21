import { FC, PropsWithChildren } from "react"

type Props = PropsWithChildren & React.HTMLProps<HTMLDivElement>

export const LandingContainer: FC<Props> = ({ children, className }) => (
	<div
		className={`flex w-full flex-row overflow-x-hidden px-4 text-plug-green sm:px-6 md:px-8 lg:px-24 ${className}`}
	>
		{children}
	</div>
)

export default LandingContainer
