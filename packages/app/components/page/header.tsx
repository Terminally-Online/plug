import { useSession } from "next-auth/react"

import { ChevronLeft, Ellipsis, GitFork, Plus, Share } from "lucide-react"

import BlockiesSvg from "blockies-react-svg"

import { Container } from "@/components/app/layout/container"
import { Header } from "@/components/app/layout/header"
import { Image } from "@/components/app/utils/image"
import { Button } from "@/components/shared/buttons/button"
import { cardColors, cn, formatAddress, formatTimeSince, formatTitle } from "@/lib"
import { useSocket } from "@/state/authentication"
import { COLUMNS, useColumnData, useColumnStore } from "@/state/columns"
import { usePlugStore } from "@/state/plugs"

const PlugHeader = () => {
	const { data: session } = useSession()
	const { column, handle } = useColumnStore(COLUMNS.MOBILE_INDEX)
	const { plug, handle: plugHandle } = usePlugStore(column?.item ?? "")

	const own = plug !== undefined && session && session.address === plug.socketId

	if (!column || !plug) return null

	return (
		<div className="flex min-h-[calc(100vh-80px)] flex-col">
			<Header
				size="lg"
				onBack={() =>
					handle.navigate({
						index: -1,
						key: column.from ?? COLUMNS.KEYS.HOME
					})
				}
				icon={
					<div
						className="h-6 w-6 min-w-6 rounded-md bg-plug-green/10"
						style={{
							backgroundImage: cardColors[plug.color]
						}}
					/>
				}
				label={plug.name === "" ? "Untitled Plug" : plug.name}
				nextOnClick={own ? () => handle.frame("manage") : () => {}}
				nextLabel={
					own ? (
						<Ellipsis size={14} />
					) : (
						<div className="flex flex-row items-center gap-2">
							<BlockiesSvg address={plug.socketId} className="h-5 w-5 rounded-md" />
							<p className="text-sm font-bold opacity-40">{formatAddress(plug.socketId)}</p>
						</div>
					)
				}
				nextEmpty={own === false}
			/>

			<div className="mb-4 flex flex-row items-center gap-4">
				<div className="font-bold opacity-40">Last updated {formatTimeSince(plug.updatedAt)}</div>

				<Button
					variant="secondary"
					className="group ml-auto p-1"
					onClick={() =>
						plugHandle.plug.fork({
							plug: plug.id,
							index: -1,
							from: column.key
						})
					}
				>
					<GitFork size={14} />
				</Button>

				<Button variant="secondary" className="group p-1" onClick={() => handle.frame("share")}>
					<Share size={14} />
				</Button>
			</div>
		</div>
	)
}

export const PageHeader = () => {
	const { column } = useColumnData(COLUMNS.MOBILE_INDEX)

	// Only show header when viewing a Plug
	if (!column || column.key !== COLUMNS.KEYS.PLUG) return null

	return (
		<Container>
			<PlugHeader />
		</Container>
	)
}
