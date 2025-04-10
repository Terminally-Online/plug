import { FC, HTMLAttributes, useState } from "react"

import {
	BookUser,
	Cable,
	CalendarCheck,
	CalendarClock,
	Computer,
	Glasses,
	Globe,
	Handshake,
	Hash,
	MessageCircleIcon,
	Puzzle,
	Sigma,
	User,
	Waypoints
} from "lucide-react"

import plugCore from "@terminallyonline/plug-core/package.json"

import { ChainId, formatAddress, getChainName } from "@/lib"
import app from "@/package.json"
import { useSocket } from "@/state/authentication"
import { Flag, useFlags } from "@/state/flags"
import { useAccount } from "@/lib/hooks/account/useAccount"
import { ChainImage } from "../../sockets/chains/chain.image"
import { connectedChains } from "@/contexts"
import { SocketDeployFrame } from "../../frames/socket/deploy/frame"
import { ColumnSettingsDeploymentItem } from "./deployment/item"
import { useResponse } from "@/lib/hooks/useResponse"
import { api } from "@/server/client"

export const ColumnSettings: FC<HTMLAttributes<HTMLDivElement> & { index: number }> = ({ index, ...props }) => {
	const { getFlag } = useFlags()

	const { user, chainId } = useAccount()
	const { socket } = useSocket()

	const [killed, setKilled] = useState(false)

	useResponse(() => api.solver.killer.killed.useQuery(undefined), {
		onSuccess: data => setKilled(data.killed)
	})

	if (!socket) return null

	return (
		<>
			<div {...props}>
				<div className="flex flex-row items-center gap-4 font-bold">
					<p className="opacity-40">Session</p>
					<div className="h-[2px] w-full bg-plug-green/10" />
				</div>
				{user && <p className="flex flex-row items-center justify-between gap-2 font-bold">
					<BookUser size={14} className="opacity-20" />
					<span className="opacity-40">Id</span>{" "}
					<span
						className="group ml-auto flex flex-row items-center gap-1 cursor-pointer"
						onClick={() => navigator.clipboard.writeText(user.id ?? "")}
					>
						{formatAddress(user.id)}
					</span>
				</p>}
				{chainId && <p className="flex flex-row items-center justify-between gap-2 font-bold">
					<Waypoints size={14} className="opacity-20" />
					<span className="opacity-40">Chain</span>{" "}
					<span className="group ml-auto flex flex-row items-center gap-2">
						<ChainImage chainId={chainId as ChainId} size="xs" />
						{getChainName(chainId as ChainId)} (<span className="opacity-40">{chainId}</span>)
					</span>
				</p>}

				<div className="mt-4 flex flex-row items-center gap-4 font-bold">
					<p className="opacity-40">Socket</p>
					<div className="h-[2px] w-full bg-plug-green/10" />
				</div>
				<p className="flex flex-row items-center justify-between gap-2 font-bold">
					<BookUser size={14} className="opacity-20" />
					<span className="opacity-40">Address</span>{" "}
					<span
						className="group ml-auto flex flex-row items-center gap-1 cursor-pointer"
						onClick={() => navigator.clipboard.writeText(socket?.socketAddress ?? "")}
					>
						{formatAddress(socket?.socketAddress)}
					</span>
				</p>
				<p className="flex flex-row items-center justify-between gap-2 font-bold">
					<Hash size={14} className="opacity-20" />
					<span className="opacity-40">Nonce</span>{" "}
					<span className="group ml-auto flex flex-row items-center gap-4">1738</span>
				</p>
				<p className="flex flex-row items-center justify-between gap-2 font-bold">
					<Puzzle size={14} className="opacity-20" />
					<span className="opacity-40">Implementation</span>{" "}
					<span
						className="group ml-auto flex cursor-pointer flex-row items-center gap-4"
						onClick={() => navigator.clipboard.writeText(socket?.deploymentImplementation ?? "")}
					>
						{socket?.deploymentImplementation ? formatAddress(socket?.deploymentImplementation) : "None"}
					</span>
				</p>

				<div className="mt-4 flex flex-row items-center gap-4 font-bold">
					<p className="opacity-40">Status</p>
					<div className="h-[2px] w-full bg-plug-green/10" />
				</div>
				<p className="flex flex-row items-center justify-between gap-2 font-bold">
					<span className="opacity-40">Solver</span>
					<span className="group ml-auto flex flex-row items-center gap-4">
						{killed ? "Halted" : "Operational"}
					</span>
				</p>

				{connectedChains.map((chain, chainIndex) => (
					<ColumnSettingsDeploymentItem 
						key={chainIndex} 
						index={index} 
						chainId={chain.id} 
						factory={socket.deploymentFactory} 
						address={socket.socketAddress} 
					/>
				))}

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

			{connectedChains.map((chain, chainIndex) => (
				<SocketDeployFrame
					key={chainIndex}
					index={index}
					chainId={chain.id}
				/>
			))}
		</>
	)
}
