import { useSession } from "next-auth/react"

import { ChevronLeft, Ellipsis, GitFork, Plus, Share } from "lucide-react"

import BlockiesSvg from "blockies-react-svg"

import { Button, Container, Header, Image } from "@/components"
import { cardColors, cn, formatAddress, formatTimeSince, formatTitle } from "@/lib"
import { COLUMNS, useColumnData, useColumnStore, usePlugStore, useSocket } from "@/state"

const HomePageHeader = () => {
	const { column, handle } = useColumnStore(COLUMNS.MOBILE_INDEX)
	const { handle: plugHandle } = usePlugStore()
	const { socket, avatar } = useSocket()

	if (!column) return null

	return (
		<Header
			size="lg"
			label={
				<>
					{socket ? (
						<button className="flex flex-row items-center gap-2" onClick={() => handle.frame("auth")}>
							{avatar ? (
								<Image
									src={avatar}
									alt="ENS Avatar"
									width={24}
									height={24}
									className="h-6 w-6 rounded-sm"
								/>
							) : (
								<BlockiesSvg className="h-6 w-6 rounded-sm" address={socket.id} />
							)}
						</button>
					) : (
						<div
							className="flex h-6 w-6 flex-row items-center justify-center rounded-sm"
							style={{
								backgroundImage: "linear-gradient(30deg, #00E100, #A3F700)"
							}}
						>
							<Image src="/white-icon.svg" alt="Logo" width={662} height={616} className="h-3 w-auto" />
						</div>
					)}

					<button
						className={cn(
							"text-lg font-bold transition-all duration-200 ease-in-out",
							column.key !== COLUMNS.KEYS.HOME ? "opacity-40 hover:opacity-100" : ""
						)}
						onClick={() =>
							handle.navigate({
								index: -1,
								key: COLUMNS.KEYS.HOME
							})
						}
					>
						Home
					</button>

					<button
						className={cn(
							"mr-auto text-lg font-bold transition-all duration-200 ease-in-out",
							column.key !== COLUMNS.KEYS.ACTIVITY ? "opacity-40 hover:opacity-100" : ""
						)}
						onClick={() =>
							handle.navigate({
								index: -1,
								key: COLUMNS.KEYS.ACTIVITY
							})
						}
					>
						Activity
					</button>
				</>
			}
			nextOnClick={() => plugHandle.plug.add({ from: column.key })}
			nextLabel={<Plus size={14} />}
		/>
	)
}

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
				onBack={
					column.from
						? () =>
								handle.navigate({
									index: -1,
									key: column.from ?? ""
								})
						: undefined
				}
				icon={
					<div
						className="h-6 w-6 min-w-6 rounded-md bg-grayscale-100"
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

const DynamicPageHeader = () => {
	const { column, handle } = useColumnStore(COLUMNS.MOBILE_INDEX)
	const { handle: plugHandle } = usePlugStore(column?.item ?? "")

	if (!column) return null

	return (
		<Header
			size="lg"
			label={
				<>
					<Button
						variant="secondary"
						className="rounded-sm p-1"
						onClick={() =>
							handle.navigate({
								index: -1,
								key: column?.from ?? COLUMNS.KEYS.HOME
							})
						}
					>
						<ChevronLeft size={14} />
					</Button>

					<Button
						className="mr-auto text-lg font-bold transition-all duration-200 ease-in-out"
						onClick={() => handle.navigate({ index: -1, key: COLUMNS.KEYS.ACTIVITY })}
					>
						{formatTitle(column?.key.toLowerCase() ?? "")}
					</Button>
				</>
			}
			nextOnClick={() => plugHandle.plug.add({ from: column.key })}
			nextLabel={<Plus size={14} />}
		/>
	)
}

export const PageHeader = () => {
	const { column } = useColumnData(COLUMNS.MOBILE_INDEX)

	return (
		<Container>
			{[COLUMNS.KEYS.HOME, COLUMNS.KEYS.ACTIVITY].includes(column?.key) ? (
				<HomePageHeader />
			) : column?.key === COLUMNS.KEYS.PLUG ? (
				<PlugHeader />
			) : (
				<DynamicPageHeader />
			)}
		</Container>
	)
}
