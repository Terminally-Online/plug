import { useState } from "react"

import Image from "next/image"

import { signOut } from "next-auth/react"

import BlockiesSvg from "blockies-react-svg"
import { AnimatePresence, motion } from "framer-motion"
import {
	Activity,
	ClipboardCheck,
	LogOut,
	PanelRightOpen,
	Plus,
	SearchIcon,
	Zap
} from "lucide-react"
import { useDisconnect } from "wagmi"

import { Button } from "@/components"
import { usePlugs, useSockets } from "@/contexts"
import { cn, useClipboard, VIEW_KEYS } from "@/lib"

export const ConsoleSidebar = () => {
	const { address, ensAvatar, socket, handle: handleSocket } = useSockets()
	const { handle: handlePlugs } = usePlugs("NOT_IMPLEMENTED")
	const { copied, handleCopied } = useClipboard(socket?.socketAddress ?? "")

	const [expanded, setExpanded] = useState(false)

	const { disconnect } = useDisconnect({
		mutation: {
			onSuccess: () => signOut({ callbackUrl: "/" })
		}
	})

	return (
		<div className="flex h-screen w-max flex-col items-center border-r-[1px] border-grayscale-100 bg-white py-4">
			<div className={cn("flex w-full flex-col gap-4 px-4")}>
				{address && (
					<button
						className="relative mb-4 h-10 w-10 rounded-sm bg-grayscale-0 transition-all duration-200 ease-in-out"
						onClick={() => handleCopied()}
					>
						<motion.div
							initial={{ opacity: 0 }}
							animate={{ opacity: 1 }}
							exit={{ opacity: 0 }}
							transition={{ duration: 0.2 }}
						>
							{ensAvatar ? (
								<Image
									src={ensAvatar}
									alt="ENS Avatar"
									width={16}
									height={16}
									className="h-full w-full rounded-sm"
								/>
							) : (
								<BlockiesSvg
									className="h-full w-full rounded-sm"
									address={address}
								/>
							)}
						</motion.div>

						<AnimatePresence>
							{copied && (
								<motion.div
									className="absolute -bottom-2 -right-2 rounded-full border-[1px] border-grayscale-0 bg-white p-1"
									initial={{ opacity: 0 }}
									animate={{ opacity: 1 }}
									exit={{ opacity: 0 }}
									transition={{ duration: 0.2 }}
								>
									<ClipboardCheck
										size={14}
										className="opacity-40"
									/>
								</motion.div>
							)}
						</AnimatePresence>
					</button>
				)}

				<button
					className="group flex flex-row items-center gap-4 px-2"
					onClick={() => handlePlugs.plug.add()}
				>
					<Button
						variant="primary"
						onClick={() => (expanded ? handlePlugs.plug.add() : {})}
						sizing="sm"
						className="rounded-sm p-1 "
					>
						<Plus
							size={14}
							className="opacity-60 transition-all duration-200 ease-in-out group-hover:opacity-100"
						/>
					</Button>
					{expanded && (
						<p className="whitespace-nowrap opacity-40 transition-all duration-200 ease-in-out group-hover:opacity-80">
							New Plug
						</p>
					)}
				</button>

				<button
					className="group flex flex-row items-center gap-4 px-2"
					onClick={() =>
						handleSocket.columns.add({
							key: VIEW_KEYS.SEARCH,
							index: 0
						})
					}
				>
					<Button
						variant="secondary"
						onClick={() =>
							expanded
								? handleSocket.columns.add({
										key: VIEW_KEYS.SEARCH,
										index: 0
									})
								: {}
						}
						sizing="sm"
						className="rounded-sm p-1 outline-none group-hover:bg-grayscale-100 group-hover:text-opacity-100"
					>
						<SearchIcon
							size={14}
							className="opacity-60 transition-all duration-200 ease-in-out group-hover:opacity-100"
						/>
					</Button>
					{expanded && (
						<p className="opacity-40 transition-all duration-200 ease-in-out group-hover:opacity-80">
							Search
						</p>
					)}
				</button>

				<button
					className="group flex flex-row items-center gap-4 px-2"
					onClick={() =>
						handleSocket.columns.add({
							key: VIEW_KEYS.ALERTS,
							index: 0
						})
					}
				>
					<Button
						variant="secondary"
						onClick={() =>
							expanded
								? handleSocket.columns.add({
										key: VIEW_KEYS.ALERTS,
										index: 0
									})
								: {}
						}
						sizing="sm"
						className="rounded-sm p-1 outline-none group-hover:bg-grayscale-100 group-hover:text-opacity-100"
					>
						<Zap
							size={14}
							className="opacity-60 transition-all duration-200 ease-in-out group-hover:opacity-100"
						/>
					</Button>
					{expanded && (
						<p className="opacity-40 transition-all duration-200 ease-in-out group-hover:opacity-80">
							Alerts
						</p>
					)}
				</button>
			</div>

			<div className="mt-auto flex w-full flex-col items-start gap-4">
				<button
					className="group flex flex-row items-center gap-4 px-6"
					onClick={() => setExpanded(!expanded)}
				>
					<Button
						variant="secondary"
						onClick={() => (expanded ? setExpanded(!expanded) : {})}
						sizing="sm"
						className="rounded-sm p-1 outline-none group-hover:bg-grayscale-100 group-hover:text-opacity-100"
					>
						<PanelRightOpen
							size={14}
							className="rotate-180 opacity-60 transition-all duration-200 ease-in-out group-hover:opacity-100"
						/>
					</Button>
					{expanded && (
						<p className="whitespace-nowrap pr-16 opacity-40 transition-all duration-200 ease-in-out group-hover:opacity-80">
							Collapse Sidebar
						</p>
					)}
				</button>
				<div className="h-[1px] w-full bg-grayscale-100" />
				<button
					className="group flex flex-row items-center gap-4 px-6"
					onClick={() => disconnect()}
				>
					<Button
						variant="secondary"
						onClick={() => (expanded ? disconnect() : {})}
						sizing="sm"
						className="rounded-sm p-1 outline-none group-hover:bg-grayscale-100 group-hover:text-opacity-100"
					>
						<LogOut
							size={14}
							className="rotate-180 opacity-60 transition-all duration-200 ease-in-out group-hover:opacity-100"
						/>
					</Button>
					{expanded && (
						<p className="opacity-40 transition-all duration-200 ease-in-out group-hover:opacity-80">
							Logout
						</p>
					)}
				</button>
			</div>
		</div>
	)
}
