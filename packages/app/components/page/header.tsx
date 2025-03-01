import { useSession } from "next-auth/react"

import { Ellipsis, GitFork, LogOut, Share } from "lucide-react"


import { Container } from "@/components/app/layout/container"
import { Header } from "@/components/app/layout/header"
import { Button } from "@/components/shared/buttons/button"
import { cardColors } from "@/lib"
import { useDisconnect } from "@/lib/hooks/wallet/useDisconnect"
import { columnByIndexAtom, COLUMNS, useColumnActions } from "@/state/columns"
import { usePlugStore } from "@/state/plugs"
import { useAtom } from "jotai"

const PlugHeader = () => {
	const { data: session } = useSession()

	const [column] = useAtom(columnByIndexAtom(COLUMNS.MOBILE_INDEX))
	const { frame, navigate } = useColumnActions(COLUMNS.MOBILE_INDEX)
	const { plug, handle: plugHandle } = usePlugStore(column?.item ?? "")

	const own = plug !== undefined && session && session.address === plug.socketId

	if (!column || !plug) return null

	return (
		<div className="flex flex-col border-b-[1px] border-plug-green/10">
			<Header
				size="lg"
				onBack={() =>
					navigate({
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
				nextOnClick={own ? () => frame("manage") : () => { }}
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
								frame("share")
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
	const { navigate } = useColumnActions(COLUMNS.MOBILE_INDEX)
	return <Header size="lg" onBack={() => navigate({ index: -1, key: COLUMNS.KEYS.HOME })} label="Discover" />
}

const MyPlugsHeader = () => {
	const { navigate } = useColumnActions(COLUMNS.MOBILE_INDEX)

	return <Header size="lg" onBack={() => navigate({ index: -1, key: COLUMNS.KEYS.HOME })} label="My Plugs" />
}

const AuthenticateHeader = () => {
	return (
		<div className="flex flex-col border-b-[1px] border-plug-green/10">
			<Header size="lg" label="Login" />
		</div>
	)
}

const ProfileHeader = () => {
	const { disconnect } = useDisconnect(true)

	const [column] = useAtom(columnByIndexAtom(COLUMNS.MOBILE_INDEX))

	if (!column) return null

	return (
		<div className="flex flex-col border-b-[1px] border-plug-green/10">
			<Header
				size="lg"
				label="Profile"
				nextLabel={
					<button
						className="flex items-center gap-2 rounded-md p-2 text-red-500 transition-colors hover:bg-red-50"
						onClick={() => disconnect()}
					>
						<LogOut size={16} />
					</button>
				}
			/>
		</div>
	)
}

export const PageHeader = () => {
	const [column] = useAtom(columnByIndexAtom(COLUMNS.MOBILE_INDEX))
	const { data: session } = useSession()

	if (!column) return null

	if (!session?.user.id?.startsWith("0x")) {
		return (
			<Container className="sticky top-0 z-10 border-b border-plug-green/10 bg-white">
				<AuthenticateHeader />
			</Container>
		)
	}

	return (
		<Container className="sticky top-0 z-10 border-b border-plug-green/10 bg-white">
			{column.key === COLUMNS.KEYS.PLUG && <PlugHeader />}
			{column.key === COLUMNS.KEYS.DISCOVER && <DiscoverHeader />}
			{column.key === COLUMNS.KEYS.MY_PLUGS && <MyPlugsHeader />}
			{column.key === COLUMNS.KEYS.PROFILE && <ProfileHeader />}
		</Container>
	)
}
