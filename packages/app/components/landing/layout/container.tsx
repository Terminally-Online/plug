import { FC, PropsWithChildren } from "react"

type Props = PropsWithChildren & React.HTMLProps<HTMLDivElement>

export const LandingContainer: FC<Props> = ({ children, className }) => (
	<div className={`flex flex-row w-full overflow-x-hidden px-4 sm:px-6 md:px-8 lg:px-24 text-plug-green ${className}`}>{children}</div>
)

export default LandingContainer

