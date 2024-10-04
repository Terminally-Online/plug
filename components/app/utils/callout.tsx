import { FC, HTMLAttributes, PropsWithChildren, ReactNode } from "react"

import { Button } from "@/components/shared"
import { usePlugs } from "@/contexts"
import { cn, greenGradientStyle, MOBILE_INDEX, VIEW_KEYS } from "@/lib"
import { useColumns, useSidebar, useSocket } from "@/state"

const Base: FC<
	PropsWithChildren<Omit<HTMLAttributes<HTMLDivElement>, "title" | "description">> & {
		title: ReactNode | JSX.Element | string
		description?: JSX.Element | string
	}
> = ({ title, description, children, className, ...props }) => (
	<div className={cn("flex h-full flex-col items-center justify-center text-center font-bold", className)} {...props}>
		<p className="max-w-[280px]">{title}</p>
		{description && <p className="mt-2 max-w-[300px] text-sm text-black text-opacity-40">{description}</p>}
		{children && <div className="mt-4 flex flex-row gap-2">{children}</div>}
	</div>
)

const Anonymous: FC<
	Omit<HTMLAttributes<HTMLDivElement>, "id" | "title" | "description"> & {
		index: number
		viewing: string
		isAbsolute?: boolean
	}
> = ({ index, viewing, isAbsolute = false, className, ...props }) => {
	const { isAnonymous } = useSocket()
	const { toggleViewingAs } = useSidebar()
	const { column, isExternal, navigate } = useColumns(index)

	if (isAnonymous === false || isExternal === true) return null

	return (
		<>
			{isAbsolute && (
				<div
					className="pointer-events-none absolute left-0 right-0 top-0 h-full bg-gradient-to-b"
					style={{
						backgroundImage: `linear-gradient(to top, rgba(255,255,255,1), rgba(255,255,255,1), rgba(255,255,255,0.85), rgba(255,255,255,0))`
					}}
				/>
			)}

			<Base
				className={cn(isAbsolute && "absolute bottom-0 left-0 right-0 top-0", className)}
				title="Your are anonymous."
				description={`To view ${viewing} you must authenticate a wallet or select an account to view as.`}
				{...props}
			>
				<Button variant="secondary" sizing="sm" onClick={() => toggleViewingAs()}>
					View As
				</Button>
				<Button sizing="sm" onClick={() => navigate({ index, key: VIEW_KEYS.AUTHENTICATE, from: column?.key })}>
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
			<div
				className="pointer-events-none absolute left-0 right-0 top-0 h-full bg-gradient-to-b"
				style={{
					backgroundImage: `linear-gradient(to top, rgba(255,255,255,1), rgba(255,255,255,1), rgba(255,255,255,0.85), rgba(255,255,255,0))`
				}}
			/>

			<Base
				className={cn("absolute bottom-0 left-0 right-0 top-0", className)}
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
		isViewing: string
		isReceivable: boolean
	}
> = ({ index, isEmpty, isViewing, isReceivable, className, ...props }) => {
	const { toggleViewingAs } = useSidebar()

	if (isEmpty === false) return null

	return (
		<>
			<div
				className="pointer-events-none absolute left-0 right-0 top-0 h-full bg-gradient-to-b"
				style={{
					backgroundImage: `linear-gradient(to top, rgba(255,255,255,1), rgba(255,255,255,1), rgba(255,255,255,0.85), rgba(255,255,255,0))`
				}}
			/>

			<Base
				className={cn("absolute bottom-0 left-0 right-0 top-0", className)}
				title="Nothing to see here, yet."
				description={`When this account has ${isViewing} they will appear here.`}
				{...props}
			>
				{index !== MOBILE_INDEX && (
					<Button variant="secondary" sizing="sm" onClick={() => toggleViewingAs()}>
						View As
					</Button>
				)}
				{isReceivable && (
					<Button sizing="sm" onClick={() => {}}>
						Receive
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
	const { handle } = usePlugs()
	const { toggleViewingAs } = useSidebar()
	const { column } = useColumns(index)

	if (!column || isEmpty === false) return null

	return (
		<>
			<div
				className="pointer-events-none absolute left-0 right-0 top-0 h-full bg-gradient-to-b"
				style={{
					backgroundImage: `linear-gradient(to top, rgba(255,255,255,1), rgba(255,255,255,1), rgba(255,255,255,0.85), rgba(255,255,255,0))`
				}}
			/>

			<Base
				className={cn("absolute bottom-0 left-0 right-0 top-0", className)}
				title="Nothing to see here, yet."
				description={" Go ahead and create a Plug from scratch or view the Plugs of another account."}
				{...props}
			>
				{index !== MOBILE_INDEX && (
					<Button variant="secondary" sizing="sm" onClick={() => toggleViewingAs()}>
						View As
					</Button>
				)}
				<Button sizing="sm" onClick={() => handle.plug.add({ index, from: column.key })}>
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
	const { column } = useColumns(index)

	if (!column || isEmpty === false) return null

	return (
		<Base
			className={cn("my-52", className)}
			title="No actions have been added, yet."
			description="Get started by adding one of the many actions available to your Plug."
			{...props}
		/>
	)
}

const EmptyPage: FC<PropsWithChildren> = () => (
	<Base
		title="Oh no! We could not find the page you were looking for."
		description="Our team has been notified of this error. If you believe this happened by accident, please wait a moment and try again."
		className="flex h-full flex-col items-center justify-center"
	>
		<div className="flex flex-col gap-4">
			<Button className="w-max" sizing="sm" href="/">
				Return Back Home
			</Button>
		</div>
	</Base>
)

export const Callout = Object.assign(Base, { Anonymous, EmptySearch, EmptyAssets, EmptyPlugs, EmptyPlug, EmptyPage })
