import { FC, HTMLAttributes, PropsWithChildren, ReactNode } from "react"
import { Button } from "@/components/shared"
import { cn, greenGradientStyle } from "@/lib"
import { COLUMNS, useColumnData, useColumnStore, usePlugStore, useSidebar, useSocket } from "@/state"

const GradientOverlay = () => (
	<div
		className="pointer-events-none absolute inset-0 z-[5] bg-gradient-to-b md:z-[3]"
		style={{
			backgroundImage: `linear-gradient(
				to top,
				rgb(253, 255, 247) 0%,
				rgb(253, 255, 247) 30%,
				rgba(253, 255, 247, 0.85) 70%,
				rgba(253, 255, 247, 0) 100%
			)`
		}}
	/>
)

const Base: FC<
	PropsWithChildren<Omit<HTMLAttributes<HTMLDivElement>, "title" | "description">> & {
		title: ReactNode | JSX.Element | string
		description?: JSX.Element | string
	}
> = ({ title, description, children, className, ...props }) => (
	<div
		className={cn(
			"absolute inset-0 z-[7] flex flex-col items-center justify-center text-center md:z-[5]",
			"p-4 font-bold md:p-0",
			"gap-2 md:gap-3",
			className
		)}
		{...props}
	>
		<p className="max-w-[280px] text-base leading-snug md:text-lg">{title}</p>
		{description && <p className="max-w-[300px] text-xs leading-relaxed text-black/40 md:text-sm">{description}</p>}
		{children && <div className="mt-3 flex flex-row gap-2 md:mt-4">{children}</div>}
	</div>
)

const Anonymous: FC<
	Omit<HTMLAttributes<HTMLDivElement>, "id" | "title" | "description"> & {
		index: number
		viewing: string
		isAbsolute?: boolean
	}
> = ({ index, viewing, isAbsolute = false, className, ...props }) => {
	const { handleSidebar } = useSidebar()
	const { isAnonymous } = useSocket()
	if (isAnonymous === false) return null
	return (
		<>
			{isAbsolute && <GradientOverlay />}
			<Base
				className={cn("z-[99999]", className)}
				title="Your are anonymous."
				description={`To view ${viewing} you must authenticate a wallet or select an account to view as.`}
				{...props}
			>
				<Button sizing="sm" onClick={() => handleSidebar("authenticating")} className="relative z-[999999]">
					Login
				</Button>
			</Base>
		</>
	)
}

const EmptySearch: FC<
	Omit<HTMLAttributes<HTMLDivElement>, "title" | "description"> & {
		isEmpty: boolean
		search: string
		handleSearch: (data: string) => void
	}
> = ({ isEmpty, search, handleSearch, className, ...props }) => {
	if (isEmpty === false) return null
	return (
		<>
			<GradientOverlay />
			<Base
				className={cn(className)}
				title={
					<>
						No results for &lsquo;
						<span style={{ ...greenGradientStyle }}>{search}</span>
						&rsquo;.
					</>
				}
				description="Your search returned no results try something else or reset your search."
				{...props}
			>
				<Button sizing="sm" onClick={() => handleSearch("")}>
					Reset
				</Button>
			</Base>
		</>
	)
}

const EmptyAssets: FC<
	Omit<HTMLAttributes<HTMLDivElement>, "title" | "description"> & {
		index: number
		isEmpty: boolean
		isViewing?: string
		isReceivable?: boolean
	}
> = ({ index, isEmpty, isViewing = "assets", isReceivable = false, className, ...props }) => {
	if (isEmpty === false) return null
	return (
		<>
			<GradientOverlay />
			<Base
				className={cn(className)}
				title="Nothing to see here, yet."
				description={`When this account has ${isViewing} they will appear here.`}
				{...props}
			>
				{isReceivable && (
					<Button sizing="sm" onClick={() => {}}>
						Deposit
					</Button>
				)}
			</Base>
		</>
	)
}

const EmptyPlugs: FC<
	Omit<HTMLAttributes<HTMLDivElement>, "title" | "description"> & {
		index: number
		isEmpty: boolean
	}
> = ({ index, isEmpty, className, ...props }) => {
	const { column } = useColumnData(index)
	const { handle } = usePlugStore()
	if (!column || isEmpty === false) return null
	return (
		<>
			<GradientOverlay />
			<Base
				className={cn(className)}
				title="Nothing to see here, yet."
				description="Go ahead and create a Plug from scratch or view the Plugs of another account."
				{...props}
			>
				<Button
					className="relative z-[12] md:z-[10]"
					sizing="sm"
					onClick={() => handle.plug.add({ index, from: column.key })}
				>
					Create
				</Button>
			</Base>
		</>
	)
}

const EmptyPlug: FC<
	Omit<HTMLAttributes<HTMLDivElement>, "title" | "description"> & {
		index: number
		isEmpty: boolean
	}
> = ({ index, isEmpty, className, ...props }) => {
	const { column } = useColumnData(index)
	if (!column || isEmpty === false) return null
	return (
		<div className="relative h-[300px]">
			<GradientOverlay />
			<Base
				className={cn(className)}
				title="No actions added, yet."
				description="Get started by adding one of the many actions available to your Plug."
				{...props}
			/>
		</div>
	)
}

const EmptyPage: FC<PropsWithChildren> = () => (
	<Base
		title="Oh no! We could not find the page you were looking for."
		description="Our team has been notified of this error. If you believe this happened by accident, please wait a moment and try again."
		className="z-[7] flex h-full flex-col items-center justify-center md:z-[5]"
	>
		<div className="flex flex-col gap-4">
			<Button className="w-max" sizing="sm" href="/">
				Return Back Home
			</Button>
		</div>
	</Base>
)

const EmptyActivity: FC<
	Omit<HTMLAttributes<HTMLDivElement>, "title" | "description"> & {
		index: number
		isEmpty: boolean
	}
> = ({ index, isEmpty, className, ...props }) => {
	const { handle } = useColumnStore(index)
	if (isEmpty === false) return null
	return (
		<>
			<GradientOverlay />
			<Base
				className={cn(className)}
				title="No activity to show yet."
				description="When you create and run Plugs, their activity will appear here."
				{...props}
			>
				<Button sizing="sm" onClick={() => handle.navigate({ index, key: COLUMNS.KEYS.DISCOVER })}>
					Discover
				</Button>
			</Base>
		</>
	)
}

export const Callout = Object.assign(Base, {
	Anonymous,
	EmptySearch,
	EmptyAssets,
	EmptyPlugs,
	EmptyPlug,
	EmptyPage,
	EmptyActivity
})