import { FC, PropsWithChildren } from "react"

type Props = PropsWithChildren

export const Footer: FC<Props> = ({ children }) => {
	return (
		<div className="fixed bottom-0 left-0 right-0 z-[1] min-w-24 bg-gradient-to-t from-[#FFFFFF] to-[#FFFFFF]/0 p-4 py-8">
			{children}
		</div>
	)
}
