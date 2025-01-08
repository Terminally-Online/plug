import { FC, HTMLAttributes } from "react"

import { CheckCircle, Clipboard } from "lucide-react"

import { formatAddress, useClipboard } from "@/lib"
import { Flag, useFlags, useSocket } from "@/state"

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

			<p className="flex justify-between font-bold">
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

			<p className="flex flex-row items-center justify-between font-bold">
				<span className="opacity-40">Nonce</span>{" "}
				<span className="group ml-auto flex flex-row items-center gap-4">1738</span>
			</p>

			<p className="flex flex-row items-center justify-between font-bold">
				<span className="opacity-40">Salt</span>{" "}
				<span className="group ml-auto flex cursor-pointer flex-row items-center gap-4">
					{formatAddress(socket?.salt ?? "")}
				</span>
			</p>

			<p className="flex justify-between font-bold">
				<span className="opacity-40">Implementation</span>{" "}
				<span className="group ml-auto flex cursor-pointer flex-row items-center gap-4">
					{socket?.implementation ? formatAddress(socket?.implementation) : "None"}
				</span>
			</p>

			<div className="mt-4 flex flex-row items-center gap-4 font-bold">
				<p className="opacity-40">Account</p>
				<div className="h-[2px] w-full bg-plug-green/10" />
			</div>

			<p className="flex justify-between font-bold">
				<span className="opacity-40">Created</span>{" "}
				<span className="group ml-auto flex flex-row items-center gap-4">
					{socket?.createdAt.toLocaleDateString()}
				</span>
			</p>

			<p className="flex justify-between font-bold">
				<span className="opacity-40">Updated</span>{" "}
				<span className="group ml-auto flex flex-row items-center gap-4">
					{socket?.updatedAt.toLocaleDateString()}
				</span>
			</p>

			<div className="mt-4 flex flex-row items-center gap-4 font-bold">
				<p className="opacity-40">Identity</p>
				<div className="h-[2px] w-full bg-plug-green/10" />
			</div>

			<p className="flex justify-between font-bold">
				<span className="opacity-40">ENS</span>{" "}
				<span className="group ml-auto flex flex-row items-center gap-4">
					{socket?.identity?.ens?.name || "None"}
				</span>
			</p>

			<p className="flex justify-between font-bold">
				<span className="opacity-40">Farcaster Id</span>{" "}
				<span className="group ml-auto flex flex-row items-center gap-4">
					{socket?.identity?.farcasterId || "None"}
				</span>
			</p>
			<p className="flex justify-between font-bold">
				<span className="opacity-40">Referred By</span>{" "}
				<span className="group ml-auto flex flex-row items-center gap-4">
					{formatAddress(socket?.identity?.referrerId ?? "") || "None"}
				</span>
			</p>

			<div className="mt-4 flex flex-row items-center gap-4 font-bold">
				<p className="opacity-40">Flags</p>
				<div className="h-[2px] w-full bg-plug-green/10" />
			</div>
			<p className="flex justify-between font-bold">
				<span className="opacity-40">Show Application</span>{" "}
				<span className="group ml-auto flex flex-row items-center gap-4">
					{getFlag(Flag.SHOW_PWA) ? "Yes" : "No"}
				</span>
			</p>
		</div>
	)
}
