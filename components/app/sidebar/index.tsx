import { useSession } from "next-auth/react"
import { FC, ReactNode, useEffect, useRef, useState } from "react"

import { motion } from "framer-motion"
import { Eye, LogOut, PanelRightOpen, Plus, ScanFace, Search, SearchIcon, X } from "lucide-react"

import { Avatar, ColumnAuthenticate, ColumnSearch, ColumnViewAs, Header, Image } from "@/components"
import { usePlugs } from "@/contexts"
import { cn, useConnect } from "@/lib"
import { useDisconnect } from "@/lib/hooks/wallet/useDisconnect"
import { useSocket } from "@/state"
import { useSidebar } from "@/state/sidebar"

const ConsoleSidebarAction: FC<
	React.HTMLAttributes<HTMLDivElement> & {
		icon: ReactNode
		isExpanded: boolean
		isPrimary?: boolean
		isActive?: boolean
	}
> = ({ icon, isExpanded, isPrimary = false, isActive = false, className, title, ...props }) => {
	return (
		<div
			className={cn(
				"group mr-auto flex h-8 w-full cursor-pointer select-none flex-row items-center justify-center gap-4 p-4 transition-all duration-200 ease-in-out",
				className
			)}
			{...props}
		>
			<div
				className={cn(
					"group flex h-8 cursor-pointer flex-row items-center justify-center gap-4 rounded-sm border-[1px] border-grayscale-100 bg-white p-4 px-2 transition-all duration-200 ease-in-out group-hover:bg-grayscale-0",
					isActive && "bg-grayscale-0 hover:bg-white",
					isPrimary &&
					"group-hover: border-plug-yellow bg-gradient-to-tr from-plug-green to-plug-yellow text-white shadow-[0_0_16px_rgba(0,255,0,1)] group-hover:shadow-[0_0_8px_rgba(0,255,0,1)]"
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
}

export const ConsoleSidebar = () => {
	const resizeRef = useRef<HTMLDivElement>(null)

	const { account } = useConnect()
	const { disconnect } = useDisconnect(true)
	const { data: session } = useSession()

	const { avatar, socket } = useSocket()
	const { handle: handlePlugs } = usePlugs("NOT_IMPLEMENTED")
	const {
		is,
		width,
		handleActivePane,
		toggleExpanded,
		toggleAuthenticating,
		toggleSearching,
		toggleViewingAs,
		resize
	} = useSidebar()

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
		<div className="flex h-full w-max select-none flex-row bg-transparent">
			<div className="flex h-full w-max flex-col items-center border-r-[1px] border-grayscale-100 py-4">
				<div className={cn("flex w-full flex-col gap-4 p-4")}>
					{session && (
						<button
							className="relative mx-4 mb-4 h-10 w-10 rounded-sm bg-grayscale-0 transition-all duration-200 ease-in-out"
							onClick={() => toggleAuthenticating()}
						>
							<motion.div
								initial={{ opacity: 0 }}
								animate={{ opacity: 1 }}
								exit={{ opacity: 0 }}
								transition={{ duration: 0.2 }}
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
							</motion.div>
						</button>
					)}

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
						onClick={() => handlePlugs.plug.add({ index: 0 })}
					/>

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
						onClick={toggleSearching}
					/>

					<ConsoleSidebarAction
						icon={
							<Eye
								size={14}
								className="opacity-60 transition-all duration-200 ease-in-out group-hover:opacity-100"
							/>
						}
						title="View As"
						isExpanded={is.expanded}
						isActive={is.viewingAs}
						onClick={toggleViewingAs}
					/>
				</div>

				<div className="mt-auto flex w-full flex-col items-center gap-4 p-4">
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

					{account.address && (
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
					)}
				</div>
			</div>

			{(is.authenticating || is.viewingAs || is.searching) && (
				<div ref={resizeRef} className="flex">
					<div
						className="m-2 mr-0 flex flex-col overflow-hidden rounded-lg border-[1px] border-grayscale-100"
						style={{
							width: `${width}px`
						}}
					>
						<div className="relative z-[30] w-full rounded-t-lg border-b-[1px] border-grayscale-100 px-4">
							<Header
								label={is.viewingAs ? "View As" : is.searching ? "Search" : "Login"}
								size="md"
								icon={
									is.searching ? (
										<SearchIcon
											size={14}
											className="m-1 opacity-40 transition-all duration-200 ease-in-out group-hover:opacity-60"
										/>
									) : is.viewingAs ? (
										<Eye
											size={14}
											className="m-1 opacity-40 transition-all duration-200 ease-in-out group-hover:opacity-60"
										/>
									) : <ScanFace
										size={14}
										className="m-1 opacity-40 transition-all duration-200 ease-in-out group-hover:opacity-60"
									/>
								}
								nextPadded={false}
								nextOnClick={() => handleActivePane(null)}
								nextLabel={<X size={14} />}
							/>
						</div>

						<div className="h-full">
							{is.authenticating && <ColumnAuthenticate index={0} />}
							{is.searching && <ColumnSearch index={0} className="px-4" />}
							{is.viewingAs && <ColumnViewAs />}
						</div>
					</div>

					<div
						className="h-full cursor-col-resize px-2"
						onMouseDown={e => {
							e.preventDefault()
							setIsResizing(true)
						}}
					>
						<div className="h-full w-[1px] bg-grayscale-100" />
					</div>
				</div>
			)}
		</div>
	)
}
