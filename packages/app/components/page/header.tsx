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
		<div className="flex flex-col border-b-[1px] border-plug-green/10">
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
					<div className="flex flex-row items-center gap-2">
						<Button
							variant="secondary"
							className="h-8 w-8 rounded-sm p-2 transition-colors hover:bg-plug-green/5"
							onClick={e => {
								e.stopPropagation()
								plugHandle.plug.fork({
									plug: plug.id,
									index: -1,
									from: column.key
								})
							}}
						>
							<GitFork size={14} />
						</Button>

						<Button
							variant="secondary"
							className="h-8 w-8 rounded-sm p-2 transition-colors hover:bg-plug-green/5"
							onClick={e => {
								e.stopPropagation()
								handle.frame("share")
							}}
						>
							<Share size={14} />
						</Button>

						{own && (
							<Button
								variant="secondary"
								className="h-8 w-8 rounded-sm p-2 transition-colors hover:bg-plug-green/5"
							>
								<Ellipsis size={14} />
							</Button>
						)}
					</div>
				}
			/>
		</div>
	)
}

const DiscoverHeader = () => {
	const { handle } = useColumnStore(COLUMNS.MOBILE_INDEX)
	return <Header size="lg" onBack={() => handle.navigate({ index: -1, key: COLUMNS.KEYS.HOME })} label="Discover" />
}

const MyPlugsHeader = () => {
	const { handle } = useColumnStore(COLUMNS.MOBILE_INDEX)
	return <Header size="lg" onBack={() => handle.navigate({ index: -1, key: COLUMNS.KEYS.HOME })} label="My Plugs" />
}

export const PageHeader = () => {
	const { column } = useColumnData(COLUMNS.MOBILE_INDEX)

	if (!column) return null

	return (
		<Container>
			{column.key === COLUMNS.KEYS.PLUG && <PlugHeader />}
			{column.key === COLUMNS.KEYS.DISCOVER && <DiscoverHeader />}
			{column.key === COLUMNS.KEYS.MY_PLUGS && <MyPlugsHeader />}
		</Container>
	)
}
