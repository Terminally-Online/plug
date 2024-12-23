import { useSession } from "next-auth/react"
import { FC, ReactNode, useEffect, useRef, useState } from "react"

import { Cat, ChartBar, LogOut, PanelRightOpen, Plus, ScanFace, Search, SearchIcon, Wallet, X } from "lucide-react"

import {
	Avatar,
	ColumnAuthenticate,
	ColumnCompanion,
	ColumnSearch,
	ColumnStats,
	ColumnWallet,
	Header,
	Image
} from "@/components"
import { cn, useConnect } from "@/lib"
import { useDisconnect } from "@/lib/hooks/wallet/useDisconnect"
import { usePlugStore, useSidebar, useSocket } from "@/state"

const ConsoleSidebarAction: FC<
	React.HTMLAttributes<HTMLDivElement> & {
		icon: ReactNode
		isExpanded: boolean
		isPrimary?: boolean
		isActive?: boolean
	}
> = ({ icon, isExpanded, isPrimary = false, isActive = false, className, title, ...props }) => (
	<div
		className={cn(
			"group mr-auto flex h-8 w-full cursor-pointer select-none flex-row items-center justify-center gap-4 p-2 transition-all duration-200 ease-in-out",
			className
		)}
		{...props}
	>
		<div
			className={cn(
				"group relative flex h-8 cursor-pointer flex-row items-center justify-center gap-4 rounded-sm border-[1px] border-plug-green/10 bg-white p-4 px-2 transition-all duration-200 ease-in-out group-hover:bg-plug-green/5",
				isActive && "bg-plug-green/5 hover:bg-white",
				isPrimary && "group-hover: border-plug-yellow bg-plug-yellow text-plug-green"
			)}
		>
			{icon}
		</div>

		<p
			className={cn(
				"mr-auto whitespace-nowrap font-bold opacity-40 transition-all duration-200 ease-in-out",
				isExpanded === false ? "hidden" : "group-hover:opacity-80"
			)}
		>
			{isExpanded ? title : "."}
		</p>
	</div>
)

const ConsoleSidebarPane = () => {
	const resizeRef = useRef<HTMLDivElement>(null)

	const { data: session } = useSession()
	const { is, width, handleActivePane, resize } = useSidebar()

	const [isResizing, setIsResizing] = useState(false)

	useEffect(() => {
		const getBoundedWidth = (width: number) => Math.min(Math.max(width, 380), 620)

		const handleMouseMove = (e: MouseEvent) => {
			if (!resizeRef.current || !isResizing) return

			resize(getBoundedWidth(e.clientX - resizeRef.current.getBoundingClientRect().left))
		}

		const handleMouseUp = () => {
			setIsResizing(false)
		}

		if (isResizing) {
			window.addEventListener("mousemove", handleMouseMove)
			window.addEventListener("mouseup", handleMouseUp)
		}

		return () => {
			window.removeEventListener("mousemove", handleMouseMove)
			window.removeEventListener("mouseup", handleMouseUp)
		}
	}, [isResizing, resize])

	return (
		<>
			{(is.authenticating || is.stats || is.companion || is.searching) && (
				<div ref={resizeRef} className="flex">
					<div
						className="relative m-2 mr-0 flex flex-col overflow-hidden rounded-lg border-[1px] border-plug-green/10"
						style={{
							width: `${width}px`
						}}
					>
						<div className="relative z-[30] w-full rounded-t-lg border-b-[1px] border-plug-green/10 bg-white px-4">
							<Header
								label={
									is.companion
										? "Companion"
										: is.stats
											? "Stats"
											: is.searching
												? "Search"
												: session?.user.id
													? "Wallet"
													: "Login"
								}
								size="md"
								icon={
									is.searching ? (
										<SearchIcon
											size={14}
											className="m-1 opacity-40 transition-all duration-200 ease-in-out group-hover:opacity-60"
										/>
									) : is.stats ? (
										<ChartBar
											size={14}
											className="m-1 opacity-40 transition-all duration-200 ease-in-out group-hover:opacity-60"
										/>
									) : is.authenticating && session?.user.id ? (
										<Wallet
											size={14}
											className="m-1 opacity-40 transition-all duration-200 ease-in-out group-hover:opacity-60"
										/>
									) : (
										<ScanFace
											size={14}
											className="m-1 opacity-40 transition-all duration-200 ease-in-out group-hover:opacity-60"
										/>
									)
								}
								nextPadded={false}
								nextOnClick={() => handleActivePane(null)}
								nextLabel={<X size={14} />}
							/>
						</div>

						<div className="h-full overflow-y-scroll">
							{is.searching ? (
								<ColumnSearch index={0} className="px-4" />
							) : is.stats ? (
								<ColumnStats index={0} />
							) : is.companion ? (
								<ColumnCompanion index={0} />
							) : session?.user.id.startsWith("0x") ? (
								<ColumnWallet index={0} />
							) : session?.user.id.startsWith("0x") === false ? (
								<ColumnAuthenticate index={0} />
							) : (
								<></>
							)}
						</div>
					</div>

					<div
						className="h-full cursor-col-resize pl-2"
						onMouseDown={e => {
							e.preventDefault()
							setIsResizing(true)
						}}
					>
						<div className="h-full w-[1px] bg-plug-green/10" />
					</div>
				</div>
			)}
		</>
	)
}

export const ConsoleSidebar = () => {
	const { account } = useConnect()
	const { disconnect } = useDisconnect(true)
	const { data: session } = useSession()
	const { socket } = useSocket()

	const { avatar } = useSocket()
	const { handle: handlePlugs } = usePlugStore("NOT_IMPLEMENTED")
	const { is, toggleExpanded, handleSidebar } = useSidebar()

	const showRestrictedOptions = account.isAuthenticated && socket?.identity?.referrerId !== null

	return (
		<div className="flex h-full w-max select-none flex-row bg-transparent">
			<div className="flex h-full w-max flex-col items-center border-r-[1px] border-plug-green/10 py-4">
				<div className={cn("flex w-full flex-col gap-4 p-2")}>
					{session && (
						<button
							className="relative mx-2 mb-4 h-10 w-10 rounded-sm bg-plug-green/5 transition-all duration-200 ease-in-out"
							onClick={() => handleSidebar("authenticating")}
						>
							{avatar ? (
								<Image
									src={avatar}
									alt="ENS Avatar"
									width={64}
									height={64}
									className="h-full w-full rounded-sm"
								/>
							) : (
								<Avatar name={socket?.id ?? ""} />
							)}
						</button>
					)}

					{showRestrictedOptions && (
						<ConsoleSidebarAction
							icon={
								<Plus
									size={14}
									className="transition-all duration-200 ease-in-out group-hover:opacity-100"
								/>
							}
							title="New Plug"
							isExpanded={is.expanded}
							isPrimary={true}
							onClick={() => handlePlugs.plug.add()}
						/>
					)}

					{showRestrictedOptions && (
						<>
							<ConsoleSidebarAction
								icon={
									<Search
										size={14}
										className="opacity-60 transition-all duration-200 ease-in-out group-hover:opacity-100"
									/>
								}
								title="Search"
								isExpanded={is.expanded}
								isActive={is.searching}
								onClick={() => handleSidebar("searching")}
							/>

							<ConsoleSidebarAction
								icon={
									<ChartBar
										size={14}
										className="opacity-60 transition-all duration-200 ease-in-out group-hover:opacity-100"
									/>
								}
								title="Stats"
								isExpanded={is.expanded}
								isActive={is.stats}
								onClick={() => handleSidebar("stats")}
							/>

							<ConsoleSidebarAction
								icon={
									<Cat
										size={14}
										className="opacity-60 transition-all duration-200 ease-in-out group-hover:opacity-100"
									/>
								}
								title="Companion"
								isExpanded={is.expanded}
								isActive={is.companion}
								onClick={() => handleSidebar("companion")}
							/>
						</>
					)}
				</div>

				<div className="mt-auto flex w-full flex-col items-center gap-4 p-2">
					{account.isAuthenticated && (
						<>
							<ConsoleSidebarAction
								className={cn(is.expanded && "pr-16")}
								icon={
									<PanelRightOpen
										size={14}
										className="rotate-180 opacity-60 transition-all duration-200 ease-in-out group-hover:opacity-100"
									/>
								}
								title="Collapse"
								isExpanded={is.expanded}
								onClick={toggleExpanded}
							/>

							<ConsoleSidebarAction
								className={cn(is.expanded && "pr-16")}
								icon={
									<LogOut
										size={14}
										className="rotate-180 opacity-60 transition-all duration-200 ease-in-out group-hover:opacity-100"
									/>
								}
								title="Logout"
								isExpanded={is.expanded}
								onClick={() => disconnect()}
							/>
						</>
					)}
				</div>
			</div>

			<ConsoleSidebarPane />
		</div>
	)
}
