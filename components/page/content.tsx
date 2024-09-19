import { ColumnSearch, Container, Plug, Plugs, PlugsDiscover, PlugsMine, Search, SocketActivity } from "@/components"
import { MOBILE_INDEX, VIEW_KEYS } from "@/lib"
import { useColumns } from "@/state"

export const PageContent = () => {
	const { column } = useColumns(MOBILE_INDEX)

	if (!column) return null

	switch (column.key) {
		case VIEW_KEYS.HOME:
			return (
				<Container>
					<Plugs hideEmpty={true} />
				</Container>
			)
		case VIEW_KEYS.DISCOVER:
			return <PlugsDiscover className="pt-4" />
		case VIEW_KEYS.MY_PLUGS:
			return <PlugsMine className="pt-4" />
		case VIEW_KEYS.PLUG:
			return (
				<Container>
					<Plug item={column.item} />
				</Container>
			)
		case VIEW_KEYS.SEARCH:
			return (
				<Container className="pt-4">
					<ColumnSearch index={column.index} />
				</Container>
			)
		case VIEW_KEYS.ACTIVITY:
			return <SocketActivity />
		default:
			return <></>
	}
}
