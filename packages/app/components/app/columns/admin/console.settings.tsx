import { useSession } from "next-auth/react"
import Link from "next/link"
import { FC, HTMLAttributes } from "react"

import { useChainId } from "wagmi"

import {
	BookUser,
	Cable,
	CalendarCheck,
	CalendarClock,
	Computer,
	Copy,
	ExternalLink,
	Flower2,
	Glasses,
	Globe,
	Handshake,
	Hash,
	MessageCircleIcon,
	Puzzle,
	User,
	Waypoints
} from "lucide-react"

import plugCore from "@terminallyonline/plug-core/package.json"

import { ChainId, formatAddress, getBlockExplorerAddress, getChainName } from "@/lib"
import app from "@/package.json"
import { useSocket } from "@/state/authentication"
import { Flag, useFlags } from "@/state/flags"

export const ConsoleSettings: FC<HTMLAttributes<HTMLDivElement> & { index: number }> = ({ index, ...props }) => {
	const { getFlag } = useFlags()
	const { data: session } = useSession()
	const { socket } = useSocket()

	const chainId = useChainId()

	if (!socket) return null

	return (
		<div {...props}>
			<div className="flex flex-row items-center gap-4 font-bold">
				<p className="opacity-40">Wallet</p>
				<div className="h-[2px] w-full bg-plug-green/10" />
			</div>
			<p className="flex flex-row items-center justify-between gap-2 font-bold">
				<BookUser size={14} className="opacity-20" />
				<span className="opacity-40">Address</span>{" "}
				<span className="group ml-auto flex flex-row items-center gap-1">
					<span
						className="cursor-pointer"
						onClick={() => navigator.clipboard.writeText(session?.address ?? "")}
					>
						{formatAddress(session?.address ?? "")}
					</span>
					<ExternalLink
						size={14}
						className="cursor-pointer opacity-20"
						onClick={() =>
							window.open(getBlockExplorerAddress(chainId as ChainId, session?.address), "_blank")
						}
					/>
				</span>
			</p>
			<p className="flex flex-row items-center justify-between gap-2 font-bold">
				<Waypoints size={14} className="opacity-20" />
				<span className="opacity-40">Connected Chain</span>{" "}
				<span className="group ml-auto flex flex-row items-center gap-4">
					{getChainName(chainId as ChainId)} ({chainId})
				</span>
			</p>

			<div className="mt-4 flex flex-row items-center gap-4 font-bold">
				<p className="opacity-40">Socket</p>
				<div className="h-[2px] w-full bg-plug-green/10" />
			</div>
			<p className="flex flex-row items-center justify-between gap-2 font-bold">
				<BookUser size={14} className="opacity-20" />
				<span className="opacity-40">Address</span>{" "}
				<span className="group ml-auto flex flex-row items-center gap-1">
					<span
						className="cursor-pointer"
						onClick={() => navigator.clipboard.writeText(socket?.socketAddress ?? "")}
					>
						{formatAddress(socket?.socketAddress)}
					</span>
					<ExternalLink
						size={14}
						className="cursor-pointer opacity-20"
						onClick={() =>
							window.open(getBlockExplorerAddress(chainId as ChainId, socket?.socketAddress), "_blank")
						}
					/>
				</span>
			</p>
			<p className="flex flex-row items-center justify-between gap-2 font-bold">
				<Hash size={14} className="opacity-20" />
				<span className="opacity-40">Nonce</span>{" "}
				<span className="group ml-auto flex flex-row items-center gap-4">1738</span>
			</p>
			<p className="flex flex-row items-center justify-between gap-2 font-bold">
				<Flower2 size={14} className="opacity-20" />
				<span className="opacity-40">Salt</span>{" "}
				<span className="group ml-auto flex cursor-pointer flex-row items-center gap-4">
					{formatAddress(socket?.salt ?? "")}
				</span>
			</p>
			<p className="flex flex-row items-center justify-between gap-2 font-bold">
				<Puzzle size={14} className="opacity-20" />
				<span className="opacity-40">Implementation</span>{" "}
				<span className="group ml-auto flex cursor-pointer flex-row items-center gap-4">
					{socket?.implementation ? formatAddress(socket?.implementation) : "None"}
				</span>
			</p>

			<div className="mt-4 flex flex-row items-center gap-4 font-bold">
				<p className="opacity-40">Account</p>
				<div className="h-[2px] w-full bg-plug-green/10" />
			</div>
			<p className="flex flex-row items-center justify-between gap-2 font-bold">
				<CalendarCheck size={14} className="opacity-20" />
				<span className="opacity-40">Created</span>
				<span className="group ml-auto flex flex-row items-center gap-4">
					{new Date(socket?.createdAt).toLocaleDateString()}
				</span>
			</p>
			<p className="flex flex-row items-center justify-between gap-2 font-bold">
				<CalendarClock size={14} className="opacity-20" />
				<span className="opacity-40">Updated</span>{" "}
				<span className="group ml-auto flex flex-row items-center gap-4">
					{new Date(socket?.updatedAt).toLocaleDateString()}
				</span>
			</p>

			<div className="mt-4 flex flex-row items-center gap-4 font-bold">
				<p className="opacity-40">Identity</p>
				<div className="h-[2px] w-full bg-plug-green/10" />
			</div>
			<p className="flex flex-row items-center justify-between gap-2 font-bold">
				<User size={14} className="opacity-20" />
				<span className="opacity-40">ENS</span>{" "}
				<span className="group ml-auto flex flex-row items-center gap-4">
					{socket?.identity?.ens?.name || "None"}
				</span>
			</p>
			<p className="flex flex-row items-center justify-between gap-2 font-bold">
				<MessageCircleIcon size={14} className="opacity-20" />
				<span className="opacity-40">Farcaster Id</span>{" "}
				<span className="group ml-auto flex flex-row items-center gap-4">
					{socket?.identity?.farcasterId || "None"}
				</span>
			</p>
			<p className="flex flex-row items-center justify-between gap-2 font-bold">
				<Handshake size={14} className="opacity-20" />
				<span className="opacity-40">Referred By</span>{" "}
				<span className="group ml-auto flex flex-row items-center gap-4">
					{formatAddress(socket?.identity?.referrerId ?? "") || "None"}
				</span>
			</p>

			<div className="mt-4 flex flex-row items-center gap-4 font-bold">
				<p className="opacity-40">Flags</p>
				<div className="h-[2px] w-full bg-plug-green/10" />
			</div>
			<p className="flex flex-row items-center justify-between gap-2 font-bold">
				<Computer size={14} className="opacity-20" />
				<span className="opacity-40">Show Application</span>{" "}
				<span className="group ml-auto flex flex-row items-center gap-4">
					{getFlag(Flag.SHOW_PWA) ? "Yes" : "No"}
				</span>
			</p>

			<div className="mt-4 flex flex-row items-center gap-4 font-bold">
				<p className="opacity-40">Versions</p>
				<div className="h-[2px] w-full bg-plug-green/10" />
			</div>
			<p className="flex flex-row items-center justify-between gap-2 font-bold">
				<Globe size={14} className="opacity-20" />
				<span className="opacity-40">App</span>{" "}
				<span className="group ml-auto flex flex-row items-center gap-4">v{app.version}</span>
			</p>
			<p className="flex flex-row items-center justify-between gap-2 font-bold">
				<Cable size={14} className="opacity-20" />
				<span className="opacity-40">Cord</span>{" "}
				<span className="group ml-auto flex flex-row items-center gap-4">v9.1.4</span>
			</p>
			<p className="flex flex-row items-center justify-between gap-2 font-bold">
				<Glasses size={14} className="opacity-20" />
				<span className="opacity-40">Core</span>{" "}
				<span className="group ml-auto flex flex-row items-center gap-4">v{plugCore.version}</span>
			</p>
		</div>
	)
}
