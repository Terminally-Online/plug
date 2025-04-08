import { useSession } from "next-auth/react"

import { useAtom } from "jotai"

import { Container } from "@/components/app/layout/container"
import { PlugsDiscover } from "@/components/app/plugs/discover"
import { PlugsMine } from "@/components/app/plugs/mine"
import { Plug } from "@/components/app/plugs/plug"
import { SocketActivity } from "@/components/app/sockets/activity/activity-list"
import { SocketAssets } from "@/components/app/sockets/assets"
import { Plugs } from "@/components/shared/framework/plugs"
import { useSocket } from "@/state/authentication"
import { columnByIndexAtom, COLUMNS } from "@/state/columns"
import { ColumnAuthenticate } from "../app/columns/authenticate/column"

const ProfileContent = () => {
	const { data: sessionData } = useSession()
	const { socket } = useSocket()

	if (!socket || !sessionData?.user.id) return null

	return (
		<div className="flex h-[calc(100vh-8rem)] flex-col">
			<div className="flex-1 overflow-y-auto p-4">
				<SocketAssets index={COLUMNS.MOBILE_INDEX} address={sessionData.user.id} hasTokens hasCollectibles />
			</div>
		</div>
	)
}

export const PageContent = () => {
	const [column] = useAtom(columnByIndexAtom(COLUMNS.MOBILE_INDEX))

	if (!column) return null

	switch (column.key) {
		case COLUMNS.KEYS.PROFILE:
			return (
				<Container>
					<ProfileContent />
				</Container>
			)
		case COLUMNS.KEYS.PLUG:
			return column.item ? <Plug index={COLUMNS.MOBILE_INDEX} item={column.item} from={column.from} /> : null
		case COLUMNS.KEYS.HOME:
			return (
				<Container className="mb-24">
					<Plugs hideEmpty={true} />
				</Container>
			)
		case COLUMNS.KEYS.DISCOVER:
			return <PlugsDiscover className="pt-4" />
		case COLUMNS.KEYS.MY_PLUGS:
			return <PlugsMine className="pt-4" />
		case COLUMNS.KEYS.ACTIVITY:
			return (
				<Container className="pt-4">
					<SocketActivity />
				</Container>
			)
		case COLUMNS.KEYS.AUTHENTICATE:
			return (
				<Container className="pt-4">
					<ColumnAuthenticate index={COLUMNS.MOBILE_INDEX} />
				</Container>
			)
		default:
			return null
	}
}
