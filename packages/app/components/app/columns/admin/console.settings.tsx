import { FC, HTMLAttributes } from "react"

import {
	BookUser,
	Cable,
	CalendarCheck,
	CalendarClock,
	CheckCircle,
	Clipboard,
	Computer,
	Flower2,
	Glasses,
	Globe,
	Handshake,
	Hash,
	MessageCircleIcon,
	Puzzle,
	User
} from "lucide-react"

import plugCore from "@terminallyonline/plug-core/package.json"

import { formatAddress, useClipboard } from "@/lib"
import app from "@/package.json"
import { useSocket } from "@/state/authentication"
import { Flag, useFlags } from "@/state/flags"

export const ConsoleSettings: FC<HTMLAttributes<HTMLDivElement> & { index: number }> = ({ index, ...props }) => {
	const { getFlag } = useFlags()
	const { socket } = useSocket()

	const { copied: socketAddressCopied, handleCopied: handleSocketAddressCopied } = useClipboard(
		socket?.socketAddress ?? ""
	)

	if (!socket) return null

	return (
		<div {...props}>
			<div className="flex flex-row items-center gap-4 font-bold">
				<p className="opacity-40">Socket</p>
				<div className="h-[2px] w-full bg-plug-green/10" />
			</div>
			<p className="flex flex-row items-center justify-between gap-2 font-bold">
				<BookUser size={14} className="opacity-60" />
				<span className="opacity-40">Address</span>{" "}
				<span
					className="group ml-auto flex cursor-pointer flex-row items-center gap-4"
					onClick={handleSocketAddressCopied}
				>
					{formatAddress(socket?.socketAddress)}
					{socketAddressCopied ? (
						<CheckCircle size={14} className="opacity-40 group-hover:opacity-100" />
					) : (
						<Clipboard size={14} className="opacity-40 group-hover:opacity-100" />
					)}
				</span>
			</p>
			<p className="flex flex-row items-center justify-between gap-2 font-bold">
				<Hash size={14} className="opacity-60" />
				<span className="opacity-40">Nonce</span>{" "}
				<span className="group ml-auto flex flex-row items-center gap-4">1738</span>
			</p>
			<p className="flex flex-row items-center justify-between gap-2 font-bold">
				<Flower2 size={14} className="opacity-60" />
				<span className="opacity-40">Salt</span>{" "}
				<span className="group ml-auto flex cursor-pointer flex-row items-center gap-4">
					{formatAddress(socket?.salt ?? "")}
				</span>
			</p>
			<p className="flex flex-row items-center justify-between gap-2 font-bold">
				<Puzzle size={14} className="opacity-60" />
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
				<CalendarCheck size={14} className="opacity-60" />
				<span className="opacity-40">Created</span>{" "}
				<span className="group ml-auto flex flex-row items-center gap-4">
					{socket?.createdAt.toLocaleDateString()}
				</span>
			</p>
			<p className="flex flex-row items-center justify-between gap-2 font-bold">
				<CalendarClock size={14} className="opacity-60" />
				<span className="opacity-40">Updated</span>{" "}
				<span className="group ml-auto flex flex-row items-center gap-4">
					{socket?.updatedAt.toLocaleDateString()}
				</span>
			</p>

			<div className="mt-4 flex flex-row items-center gap-4 font-bold">
				<p className="opacity-40">Identity</p>
				<div className="h-[2px] w-full bg-plug-green/10" />
			</div>
			<p className="flex flex-row items-center justify-between gap-2 font-bold">
				<User size={14} className="opacity-60" />
				<span className="opacity-40">ENS</span>{" "}
				<span className="group ml-auto flex flex-row items-center gap-4">
					{socket?.identity?.ens?.name || "None"}
				</span>
			</p>
			<p className="flex flex-row items-center justify-between gap-2 font-bold">
				<MessageCircleIcon size={14} className="opacity-60" />
				<span className="opacity-40">Farcaster Id</span>{" "}
				<span className="group ml-auto flex flex-row items-center gap-4">
					{socket?.identity?.farcasterId || "None"}
				</span>
			</p>
			<p className="flex flex-row items-center justify-between gap-2 font-bold">
				<Handshake size={14} className="opacity-60" />
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
				<Computer size={14} className="opacity-60" />
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
				<Globe size={14} className="opacity-60" />
				<span className="opacity-40">App</span>{" "}
				<span className="group ml-auto flex flex-row items-center gap-4">v{app.version}</span>
			</p>
			<p className="flex flex-row items-center justify-between gap-2 font-bold">
				<Cable size={14} className="opacity-60" />
				<span className="opacity-40">Cord</span>{" "}
				<span className="group ml-auto flex flex-row items-center gap-4">v9.1.4</span>
			</p>
			<p className="flex flex-row items-center justify-between gap-2 font-bold">
				<Glasses size={14} className="opacity-60" />
				<span className="opacity-40">Core</span>{" "}
				<span className="group ml-auto flex flex-row items-center gap-4">v{plugCore.version}</span>
			</p>
		</div>
	)
}
