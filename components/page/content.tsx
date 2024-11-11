import {
	ColumnAuthenticate,
	Container,
	Plug,
	Plugs,
	PlugsDiscover,
	PlugsMine,
	SocketActivity,
	SocketAssets,
	SocketProfile
} from "@/components"
import { COLUMNS, useColumnData } from "@/state"

export const PageContent = () => {
	const { column } = useColumnData(COLUMNS.MOBILE_INDEX)

	if (!column) return null

	switch (column.key) {
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
		case COLUMNS.KEYS.PLUG:
			return (
				<Container>
					<Plug item={column.item} />
				</Container>
			)
		case COLUMNS.KEYS.ACTIVITY:
			return (
				<Container className="pt-4">
					<SocketActivity />
				</Container>
			)
		case COLUMNS.KEYS.PROFILE:
			return (
				<Container className="pt-4">
					<SocketProfile />
					<SocketAssets />
				</Container>
			)
		case COLUMNS.KEYS.AUTHENTICATE:
			return (
				<Container className="pt-4">
					<ColumnAuthenticate index={COLUMNS.MOBILE_INDEX} />
				</Container>
			)
		default:
			return <></>
	}
}
