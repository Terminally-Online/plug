import type { FC, PropsWithChildren } from "react"

type Props = PropsWithChildren & React.HTMLProps<HTMLDivElement>

export const Container: FC<Props> = ({ children, className }) => (
	<div className={`mx-4 flex flex-col overflow-x-hidden ${className}`}>
		{children}
	</div>
)

export default Container
