import { useSession } from "next-auth/react"
import { FC, ReactNode, useEffect, useRef, useState } from "react"

import { Cat, ChartBar, Code, LogOut, PanelRightOpen, Plus, ScanFace, Wallet, X } from "lucide-react"

import { ColumnAuthenticate } from "@/components/app/columns/utils/column-authenticate"
import { ColumnCompanion } from "@/components/app/columns/utils/column-companion"
import { ColumnStats } from "@/components/app/columns/utils/column-stats"
import { ColumnWallet } from "@/components/app/columns/utils/column-wallet"
import { Header } from "@/components/app/layout/header"
import { Avatar } from "@/components/app/sockets/profile"
import { Image } from "@/components/app/utils/image"
import { cn, useConnect } from "@/lib"
import { useDisconnect } from "@/lib/hooks/wallet/useDisconnect"
import { useSocket } from "@/state/authentication"
import { Flag, useFlags } from "@/state/flags"
import { usePlugStore } from "@/state/plugs"
import { useSidebar } from "@/state/sidebar"

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
			"group relative mr-auto flex h-10 w-full cursor-pointer flex-row items-center justify-center gap-4 rounded-sm border-[1px] border-plug-green/10 bg-white p-4 py-2 transition-all duration-200 ease-in-out",
			isActive && "bg-plug-green/10 hover:bg-white",
			isPrimary
				? "border-plug-yellow bg-plug-yellow text-plug-green hover:brightness-105"
				: "text-black hover:bg-plug-green/10",
			className
		)}
		{...props}
	>
		<div className="opacity-60">{icon}</div>

		<p
			className={cn(
				"mr-auto whitespace-nowrap font-bold opacity-80 transition-all duration-200 ease-in-out",
				isExpanded === false ? "hidden" : "group-hover:opacity-100"
			)}
		>
			{isExpanded ? title : "."}
		</p>
	</div>
)

export const ConsoleSidebarPane = () => {
	const resizeRef = useRef<HTMLDivElement>(null)

	const { data: session } = useSession()
	const { account } = useConnect()
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
			{(is.authenticating || !account.isAuthenticated || is.stats || is.companion || is.searching) && (
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
											: session?.user.id.startsWith("0x")
												? "Wallet"
												: "Login"
								}
								size="md"
								icon={
									is.stats ? (
										<ChartBar
											size={14}
											className="m-1 opacity-40 transition-all duration-200 ease-in-out group-hover:opacity-60"
										/>
									) : session?.user.id.startsWith("0x") ? (
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
								nextLabel={account.isAuthenticated ? <X size={14} /> : undefined}
							/>
						</div>

						<div className="h-full overflow-y-scroll">
							{is.stats ? (
								<ColumnStats index={0} />
							) : is.companion ? (
								<ColumnCompanion index={0} />
							) : session?.user.id.startsWith("0x") ? (
								<ColumnWallet index={0} />
							) : session?.user.id.startsWith("0x") === false || !account.isAuthenticated ? (
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
	const { handleFlag, getFlag } = useFlags()
	const { socket } = useSocket()

	const { avatar } = useSocket()
	const { handle: handlePlugs } = usePlugStore("NOT_IMPLEMENTED")
	const { is, toggleExpanded, handleSidebar } = useSidebar()

	const showRestrictedOptions = account.isAuthenticated && socket?.identity?.referrerId !== null

	return (
		<div className="flex h-full w-max select-none flex-row bg-transparent">
			<div className="flex h-full flex-col items-center border-r-[1px] border-plug-green/10 p-4 pt-2">
				<div className="flex w-full flex-col items-start gap-2 p-2">
					<button
						className="relative mb-4 h-12 w-12 rounded-sm bg-plug-green/5 transition-all duration-200 ease-in-out"
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
								onClick={() => handlePlugs.plug.add()}
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

				<div className="mt-auto flex w-full flex-col items-center gap-2 p-2">
					{(account.address || account.isAuthenticated) && (
						<>
							{socket?.admin && (
								<ConsoleSidebarAction
									className={cn(
										is.expanded && "pr-16",
										getFlag(Flag.SHOW_DEVELOPER) &&
											"border-plug-yellow bg-plug-yellow text-plug-green"
									)}
									icon={
										<Code
											size={14}
											className="opacity-60 transition-all duration-200 ease-in-out group-hover:opacity-100"
										/>
									}
									title="Developer"
									isExpanded={is.expanded}
									onClick={() => handleFlag(Flag.SHOW_DEVELOPER, !getFlag(Flag.SHOW_DEVELOPER))}
								/>
							)}

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

			{/* <ConsoleSidebarPane /> */}
		</div>
	)
}
