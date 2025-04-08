import { FC, ReactNode, useEffect, useRef, useState } from "react"

import { ChartBar, LogOut, Plus, ScanFace, Settings, Wallet, X } from "lucide-react"
import { ColumnStats } from "@/components/app/columns/utils/column-stats"
import { ColumnWallet } from "@/components/app/columns/utils/column-wallet"
import { Header } from "@/components/app/layout/header"
import { Avatar } from "@/components/app/sockets/profile"
import { Image } from "@/components/app/utils/image"
import { cn } from "@/lib"
import { useSocket } from "@/state/authentication"
import { usePlugActions } from "@/state/plugs"
import { useSidebar } from "@/state/sidebar"
import { useAccount } from "@/lib/hooks/account/useAccount"
import { useDisconnect } from "@/lib/hooks/account/useDisconnect"
import { ColumnAuthenticate } from "@/components/app/columns/authenticate/column"
import { ColumnSettings } from "../columns/settings/column"

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
			"group relative mx-auto mr-auto flex h-10 w-10 cursor-pointer flex-row items-center justify-center gap-4 rounded-sm border-[1px] border-plug-green/10 bg-white p-4 py-2 transition-all duration-200 ease-in-out",
			isActive && "bg-plug-green/10 hover:bg-white",
			isPrimary
				? "border-plug-yellow bg-plug-yellow text-plug-green hover:brightness-105"
				: "text-black hover:bg-plug-green/10",
			className
		)}
		{...props}
	>
		<div className={cn(!isActive && "opacity-80")}>{icon}</div>

		<p
			className={cn(
				"mr-auto whitespace-nowrap font-bold transition-all duration-200 ease-in-out",
				!isActive && "opacity-80",
				isExpanded === false ? "hidden" : "group-hover:opacity-100"
			)}
		>
			{isExpanded ? title : "."}
		</p>
	</div>
)

export const ConsoleSidebarPane = () => {
	const resizeRef = useRef<HTMLDivElement>(null)

	const { is, width, handleActivePane, resize } = useSidebar()
	const { socket } = useSocket()

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
			{(is.authenticating || is.stats || is.searching || is.settings) && (
				<div ref={resizeRef} className="flex">
					<div
						className="relative mr-0 flex flex-col overflow-hidden"
						style={{
							width: `${width}px`
						}}
					>
						<div className="relative z-[30] w-full rounded-t-lg border-b-[1px] border-plug-green/10 bg-white px-4">
							<Header
								label={
									is.settings ? "Settings" :
										is.stats
											? "Stats"
											: socket.id.startsWith("0x")
												? "Wallet"
												: "Login"
								}
								size="md"
								icon={
									is.settings ? (
										<Settings
											size={14}
											className="m-1 opacity-40 transition-all duration-200 ease-in-out group-hover:opacity-60"
										/>

									) : is.stats ? (
										<ChartBar
											size={14}
											className="m-1 opacity-40 transition-all duration-200 ease-in-out group-hover:opacity-60"
										/>
									) : socket.id.startsWith("0x") ? (
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
								nextLabel={socket.id.startsWith("0x") ? <X size={14} /> : undefined}
							/>
						</div>

						<div className="h-full overflow-y-scroll">
							{is.settings ? (
								<ColumnSettings index={0} className="px-6 py-4" />
							) : is.stats ? (
								<ColumnStats index={0} />
							) : socket.id.startsWith("0x") ? (
								<ColumnWallet index={0} />
							) : socket.id.startsWith("0x") === false ? (
								<ColumnAuthenticate index={0} />
							) : (
								<></>
							)}
						</div>
					</div>
					<div className="relative h-full cursor-col-resize">
						<div className="h-full w-[1px] bg-plug-green/10" />
						<div
							className="absolute -left-4 -right-4 bottom-0 top-0 z-[999]"
							onMouseDown={e => {
								e.preventDefault()
								setIsResizing(true)
							}}
						/>
					</div>
				</div>
			)}
		</>
	)
}

export const ConsoleSidebar = () => {
	const { address } = useAccount()
	const { disconnect } = useDisconnect(true)

	const { is, handleSidebar: sidebar } = useSidebar()
	const { socket, avatar } = useSocket()
	const { add } = usePlugActions()

	const showRestrictedOptions = !!socket.id.startsWith("0x") && !!socket?.identity?.onboardingAt

	return (
		<div className="flex h-full w-max select-none flex-row bg-transparent">
			<div className="flex h-full flex-col items-center border-r-[1px] border-plug-green/10 p-2 pt-2">
				<div className="flex w-full flex-col items-start gap-2 p-2">
					<button
						className="relative mb-4 h-10 w-10 rounded-sm bg-plug-green/5 transition-all duration-200 ease-in-out"
						onClick={() => sidebar("authenticating")}
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

					{showRestrictedOptions && (
						<>
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
								onClick={() => add()}
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
								onClick={() => sidebar("stats")}
							/>
						</>
					)}
				</div>

				<div className="mt-auto flex w-full flex-col items-center gap-2 p-2">
					{(socket || address) && (
						<>
							<ConsoleSidebarAction
								className={cn(is.expanded && "pr-16")}
								icon={
									<Settings
										size={14}
										className="rotate-180 opacity-60 transition-all duration-200 ease-in-out group-hover:opacity-100"
									/>
								}
								title="Settings"
								isExpanded={is.expanded}
								isActive={is.settings}
								onClick={() => sidebar("settings")}
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
		</div>
	)
}
