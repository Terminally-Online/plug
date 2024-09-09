import { useState } from "react"

import Image from "next/image"

import { signOut } from "next-auth/react"

import Avatar from "boring-avatars"
import { AnimatePresence, motion } from "framer-motion"
import { BookUser, ClipboardCheck, LogOut, PanelRightOpen, Plus, SearchIcon, Zap } from "lucide-react"
import { useDisconnect } from "wagmi"

import { Button } from "@/components"
import { usePlugs, useSockets } from "@/contexts"
import { cn, useClipboard, VIEW_KEYS } from "@/lib"

export const ConsoleSidebar = () => {
	const { address, avatar, socket, handle: handleSocket } = useSockets()
	const { handle: handlePlugs } = usePlugs("NOT_IMPLEMENTED")
	const { copied, handleCopied } = useClipboard(socket?.socketAddress ?? "")

	const [expanded, setExpanded] = useState(false)

	const { disconnect } = useDisconnect({
		mutation: {
			onSuccess: () => signOut()
		}
	})

	return (
		<div className="mr-2 flex h-full w-max flex-col items-center border-r-[1px] border-grayscale-100 bg-white py-4">
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
							{avatar ? (
								<Image
									src={avatar}
									alt="ENS Avatar"
									width={64}
									height={64}
									className="h-full w-full rounded-sm"
								/>
							) : (
								<div className="overflow-hidden rounded-sm">
									<Avatar
										name={address}
										variant="beam"
										size={"100%"}
										colors={["#00E100", "#A3F700"]}
										square
									/>
								</div>
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
									<ClipboardCheck size={14} className="opacity-40" />
								</motion.div>
							)}
						</AnimatePresence>
					</button>
				)}

				<div
					className="group flex cursor-pointer flex-row items-center gap-4 px-2"
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
				</div>

				<div
					className="group flex cursor-pointer flex-row items-center gap-4 px-2"
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
				</div>

				<div
					className="group flex cursor-pointer flex-row items-center gap-4 px-2"
					onClick={() =>
						handleSocket.columns.add({
							key: VIEW_KEYS.VIEW_AS,
							index: 0
						})
					}
				>
					<Button
						variant="secondary"
						onClick={() =>
							expanded
								? handleSocket.columns.add({
										key: VIEW_KEYS.VIEW_AS,
										index: 0
									})
								: {}
						}
						sizing="sm"
						className="rounded-sm p-1 outline-none group-hover:bg-grayscale-100 group-hover:text-opacity-100"
					>
						<BookUser
							size={14}
							className="opacity-60 transition-all duration-200 ease-in-out group-hover:opacity-100"
						/>
					</Button>
					{expanded && (
						<p className="opacity-40 transition-all duration-200 ease-in-out group-hover:opacity-80">
							View As
						</p>
					)}
				</div>

				<div
					className="group flex cursor-pointer flex-row items-center gap-4 px-2"
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
				</div>
			</div>

			<div className="mt-auto flex w-full flex-col items-start gap-4">
				<div
					className="group flex cursor-pointer flex-row items-center gap-4 px-6"
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
				</div>

				{address && (
					<>
						<div className="h-[1px] w-full bg-grayscale-100" />
						<div
							className="group flex cursor-pointer flex-row items-center gap-4 px-6"
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
						</div>
					</>
				)}
			</div>
		</div>
	)
}
