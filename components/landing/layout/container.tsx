import { FC, PropsWithChildren } from "react"

type Props = PropsWithChildren & React.HTMLProps<HTMLDivElement>

export const LandingContainer: FC<Props> = ({ children, className }) => (
	<div className={`flex flex-row px-8 lg:px-24 ${className}`}>{children}</div>
)

export default LandingContainer
