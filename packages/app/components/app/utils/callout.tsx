import { FC, HTMLAttributes, PropsWithChildren, ReactNode } from "react"

import { HTMLMotionProps, motion } from "framer-motion"
import { Loader } from "lucide-react"

import { Button } from "@/components/shared/buttons/button"
import { cn, greenGradientStyle } from "@/lib"
import { useSocket } from "@/state/authentication"
import { COLUMNS, useColumnData, useColumnStore } from "@/state/columns"
import { usePlugStore } from "@/state/plugs"
import { useSidebar } from "@/state/sidebar"

const Base: FC<
	PropsWithChildren & {
		title: ReactNode | JSX.Element | string
		description?: JSX.Element | string
		className?: string
	}
> = ({ title, description, children, className }) => (
	<motion.div
		className={cn("flex h-full flex-col items-center justify-center gap-2 text-center font-bold", className)}
		initial={{ opacity: 0 }}
		animate={{ opacity: 1 }}
		transition={{ duration: 0.2 }}
	>
		<p className="max-w-[280px] text-xl">{title}</p>
		{description && <p className="max-w-[300px] text-sm text-black text-opacity-40">{description}</p>}
		{children && <div className="mt-4 flex flex-row gap-2">{children}</div>}
	</motion.div>
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
			{isAbsolute && (
				<div
					className="pointer-events-none absolute left-0 right-0 top-0 h-full bg-gradient-to-b"
					style={{
						backgroundImage: `linear-gradient(to top, rgb(253, 255, 247), rgb(253, 255, 247), rgba(253, 255, 247, 0.85), rgba(253, 255, 247, 0))`
					}}
				/>
			)}
			<Base
				className={cn(isAbsolute && "absolute bottom-0 left-0 right-0 top-0", className)}
				title="Your are anonymous."
				description={`To view ${viewing} you must authenticate a wallet or select an account to view as.`}
				{...props}
			>
				<Button sizing="sm" onClick={() => handleSidebar("authenticating")}>
					Login
				</Button>
			</Base>
		</>
	)
}

const Loading: FC<
	Omit<HTMLAttributes<HTMLDivElement>, "title" | "description"> & {
		index: number
	}
> = ({ index, className, ...props }) => {
	const { column } = useColumnData(index)

	if (!column) return null

	return (
		<div className={cn("absolute bottom-0 left-0 right-0 top-0", className)} {...props}>
			<Loader size={14} className="animate-spin" />
		</div>
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
					backgroundImage: `linear-gradient(to top, rgb(253, 255, 247), rgb(253, 255, 247), rgba(253, 255, 247, 0.85), rgba(253, 255, 247, 0))`
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
	const {
		is: { authenticating },
		handleSidebar
	} = useSidebar()

	if (isEmpty === false) return null

	return (
		<>
			<div
				className="pointer-events-none absolute left-0 right-0 top-0 h-full bg-gradient-to-b"
				style={{
					backgroundImage: `linear-gradient(to top, rgb(253, 255, 247), rgb(253, 255, 247), rgba(253, 255, 247, 0.85), rgba(253, 255, 247, 0))`
				}}
			/>

			<Base
				className={cn("absolute bottom-0 left-0 right-0 top-0", className)}
				title="Nothing to see here, yet."
				description={`When this account has ${isViewing} they will appear here.`}
				{...props}
			>
				{isReceivable && (
					<Button variant={authenticating ? "primaryDisabled" : "primary"} sizing="sm" onClick={() => handleSidebar("authenticating")}>
						{authenticating ? "Depositing..." : "Deposit"}
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
			<div
				className="pointer-events-none absolute left-0 right-0 top-0 z-[9999] h-full bg-gradient-to-b"
				style={{
					backgroundImage: `linear-gradient(to top, rgb(253, 255, 247), rgb(253, 255, 247), rgba(253, 255, 247, 0.85), rgba(253, 255, 247, 0))`
				}}
			/>

			<Base
				className={cn("absolute bottom-0 left-0 right-0 top-0 z-[9999]", className)}
				title="Nothing to see here, yet."
				description={" Go ahead and create a Plug from scratch or view the Plugs of another account."}
				{...props}
			>
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
	const { column } = useColumnData(index)
	if (!column || isEmpty === false) return null
	return (
		<Base
			className={cn("my-52", className)}
			title="No actions added, yet."
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
			<div
				className="pointer-events-none absolute left-0 right-0 top-0 h-full bg-gradient-to-b"
				style={{
					backgroundImage: `linear-gradient(to top, rgb(253, 255, 247), rgb(253, 255, 247), rgba(253, 255, 247, 0.85), rgba(253, 255, 247, 0))`
				}}
			/>

			<Base
				className={cn("absolute bottom-0 left-0 right-0 top-0", className)}
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
	Loading,
	EmptySearch,
	EmptyAssets,
	EmptyPlugs,
	EmptyPlug,
	EmptyPage,
	EmptyActivity
})
