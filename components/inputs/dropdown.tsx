import { FC, PropsWithChildren } from "react"

import { ChevronDown } from "lucide-react"

import { cn } from "@/lib/utils"

type Props = {
	icon: JSX.Element
	placeholder: string
	value?: string
	options: Array<{
		label: string
		value: string
	}>
	handleClick?: () => void
} & PropsWithChildren &
	React.HTMLAttributes<HTMLDivElement>

export const Dropdown: FC<Props> = ({
	icon,
	placeholder,
	value,
	handleClick,
	className,
	children
}) => {
	return (
		<>
			<button
				className={cn("flex flex-col gap-2", className)}
				onClick={() => (handleClick ? handleClick() : undefined)}
			>
				<div
					className={cn(
						"flex w-full cursor-pointer flex-row items-center gap-4 bg-grayscale-0 p-4 px-6",
						"rounded-full"
					)}
				>
					{icon}

					<p className="flex gap-2">
						<span className="font-bold">{placeholder}</span>
						<span className="opacity-60">{value}</span>
					</p>

					<ChevronDown size={14} className="ml-auto opacity-60" />

					{children && <div className="ml-auto">{children}</div>}
				</div>
			</button>
		</>
	)
}
