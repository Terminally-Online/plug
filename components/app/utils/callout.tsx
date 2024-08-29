import { FC, HTMLAttributes, PropsWithChildren, ReactNode } from "react"

import { Button } from "@/components/shared"
import { useSockets } from "@/contexts"
import { cn, greenGradientStyle } from "@/lib"

const Base: FC<
	PropsWithChildren<Omit<HTMLAttributes<HTMLDivElement>, "title" | "description">> & {
		title: ReactNode | JSX.Element | string
		description?: string
	}
> = ({ title, description, children, className, ...props }) => (
	<div className={cn("flex h-full flex-col items-center justify-center text-center font-bold", className)} {...props}>
		<p>{title}</p>
		{description && <p className="max-w-[320px] opacity-40">{description}</p>}
		{children && <div className="mt-4 flex flex-row gap-2">{children}</div>}
	</div>
)

const Anonymous: FC<Omit<HTMLAttributes<HTMLDivElement>, "title" | "description"> & { viewing: string }> = ({
	viewing,
	...props
}) => {
	const { isAnonymous } = useSockets()

	if (isAnonymous === false) return null

	return (
		<Base
			{...props}
			title="Your are anonymous."
			description={`To view ${viewing} you must authenticate a wallet or select an account to view as.`}
		>
			<Button variant="secondary" sizing="sm" onClick={() => {}}>
				View As
			</Button>
			<Button sizing="sm" onClick={() => {}}>
				Login
			</Button>
		</Base>
	)
}

const EmptySearch: FC<
	Omit<HTMLAttributes<HTMLDivElement>, "title" | "description"> & {
		isEmpty: boolean
		search: string
		handleSearch: (data: string) => void
	}
> = ({ isEmpty, search, handleSearch, ...props }) => {
	if (isEmpty === false) return null

	return (
		<Base
			{...props}
			title={
				<>
					No results for &lsquo;
					<span
						style={{
							...greenGradientStyle
						}}
					>
						{search}
					</span>
					&rsquo;.
				</>
			}
			description="Your search returned no results."
		>
			<Button sizing="sm" onClick={() => handleSearch("")}>
				Reset
			</Button>
		</Base>
	)
}

export const Callout = Object.assign(Base, { Anonymous, EmptySearch })
