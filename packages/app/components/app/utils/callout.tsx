import { FC, HTMLAttributes, type JSX, PropsWithChildren, ReactNode } from "react"

import { motion } from "framer-motion"
import { Loader } from "lucide-react"

import { useAtom } from "jotai"

import { Button } from "@/components/shared/buttons/button"
import { cn, greenGradientStyle } from "@/lib"
import { useSocket } from "@/state/authentication"
import { columnByIndexAtom, } from "@/state/columns"
import { useSidebar } from "@/state/sidebar"

const Base: FC<
	PropsWithChildren & {
		title: ReactNode | JSX.Element | string
		description?: JSX.Element | string
		className?: string
	}
> = ({ title, description, children, className }) => (
	<motion.div
		className={cn("relative z-[9999] flex h-full flex-col items-center justify-center gap-2 text-center font-bold", className)}
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
	const [column] = useAtom(columnByIndexAtom(index))

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

const EmptyOverlay = () => <>
	<div
		className="z-[9999] pointer-events-none absolute left-0 right-0 top-0 h-2/3 bg-gradient-to-b"
		style={{
			backgroundImage: `linear-gradient(to top, rgb(253, 255, 247), rgb(253, 255, 247), rgba(253, 255, 247, 0.85), rgba(253, 255, 247, 0))`
		}}
	/>
	<div className="pointer-events-none absolute left-0 right-0 top-2/3 h-1/3 bg-plug-white" />
</>

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
			<EmptyOverlay />

			<Base
				className={cn("absolute bottom-0 left-0 right-0 top-0", className)}
				title="Nothing to see here, yet."
				description={`When this Socket has ${isViewing} they will automatically appear here.`}
				{...props}
			/>
		</>
	)
}

const EmptyActivity: FC<
	Omit<HTMLAttributes<HTMLDivElement>, "title" | "description"> & {
		index: number
		isEmpty: boolean
	}
> = ({ index, isEmpty, className, ...props }) => {
	if (isEmpty === false) return null

	return (
		<>
			<EmptyOverlay />

			<Base
				className={cn("absolute bottom-0 left-0 right-0 top-0", className)}
				title="Nothing to see here, yet."
				description="When you create and run Plugs, their activity will appear here."
				{...props}
			/>
		</>
	)
}

const EmptyPlugs: FC<
	Omit<HTMLAttributes<HTMLDivElement>, "title" | "description"> & {
		index: number
		isEmpty: boolean
	}
> = ({ index, isEmpty, className, ...props }) => {
	if (isEmpty === false) return null

	return (
		<>
			<EmptyOverlay />

			<Base
				className={cn("absolute bottom-0 left-0 right-0 top-0 z-[9999]", className)}
				title="Nothing to see here, yet."
				description={" Go ahead and create a Plug from scratch or view the Plugs of another account."}
				{...props}
			/>
		</>
	)
}
const EmptyPlug: FC<
	Omit<HTMLAttributes<HTMLDivElement>, "title" | "description"> & {
		index: number
		isEmpty: boolean
	}
> = ({ index, isEmpty, ...props }) => {
	const [column] = useAtom(columnByIndexAtom(index))

	if (!column || isEmpty === false) return null

	return (
		<Base
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
