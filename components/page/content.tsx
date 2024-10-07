import {
	ColumnAuthenticate,
	ColumnSearch,
	Container,
	Plug,
	Plugs,
	PlugsDiscover,
	PlugsMine,
	SocketActivity,
	SocketAssets,
	SocketProfile
} from "@/components"
import { COLUMN_KEYS, MOBILE_INDEX, useColumns } from "@/state"

export const PageContent = () => {
	const { column } = useColumns(MOBILE_INDEX)

	if (!column) return null

	switch (column.key) {
		case COLUMN_KEYS.HOME:
			return (
				<Container className="mb-24">
					<Plugs hideEmpty={true} />
				</Container>
			)
		case COLUMN_KEYS.DISCOVER:
			return <PlugsDiscover className="pt-4" />
		case COLUMN_KEYS.MY_PLUGS:
			return <PlugsMine className="pt-4" />
		case COLUMN_KEYS.PLUG:
			return (
				<Container>
					<Plug item={column.item} />
				</Container>
			)
		case COLUMN_KEYS.ACTIVITY:
			return (
				<Container className="pt-4">
					<SocketActivity />
				</Container>
			)
		case COLUMN_KEYS.PROFILE:
			return (
				<Container className="pt-4">
					<SocketProfile />
					<SocketAssets />
				</Container>
			)
		case COLUMN_KEYS.AUTHENTICATE:
			return (
				<Container className="pt-4">
					<ColumnAuthenticate index={column.index} />
				</Container>
			)
		default:
			return <></>
	}
}
