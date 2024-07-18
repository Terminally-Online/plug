import { FC, PropsWithChildren } from "react"

type Props = PropsWithChildren & React.HTMLProps<HTMLDivElement>

export const LandingContainer: FC<Props> = ({ children, className }) => (
	<div className={`mx-4 flex flex-row lg:mx-24 ${className}`}>{children}</div>
)

export default LandingContainer
