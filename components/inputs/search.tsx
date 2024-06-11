import type { FC, PropsWithChildren, RefObject } from "react"
import { useRef } from "react"

import { cn } from "@/lib/utils"

type Props = {
	ref?: RefObject<HTMLInputElement>
	icon: JSX.Element
	placeholder: string
	search?: string
	handleSearch?: (search: string) => void
	handleOnClick?: () => void
	hasClear?: boolean
} & PropsWithChildren &
	React.HTMLAttributes<HTMLDivElement>

export const Search: FC<Props> = ({
	ref,
	icon,
	placeholder,
	search,
	handleSearch,
	handleOnClick,
	className,
	children
}) => {
	const searchRef = useRef<HTMLInputElement>(null)
	// NOTE: Accept an external `ref` that is managed by an external Frame trigger.
	ref = ref ? ref : searchRef

	return (
		<div className={cn("flex flex-col gap-2", className)}>
			<div
				className="flex w-full cursor-pointer items-center gap-4 rounded-full bg-grayscale-0 p-4 px-6"
				onClick={
					handleOnClick
						? () => handleOnClick()
						: () => ref.current?.focus()
				}
			>
				<div className="w-max">{icon}</div>
				<input
					ref={ref}
					type="text"
					placeholder={placeholder}
					className="w-full cursor-pointer bg-transparent outline-none"
					value={search}
					onChange={e =>
						handleSearch ? handleSearch(e.target.value) : null
					}
				/>

				<div className="ml-auto">{children}</div>
			</div>
		</div>
	)
}
