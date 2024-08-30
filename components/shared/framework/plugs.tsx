import { FC, HTMLAttributes } from "react"

import { PlugZap, Puzzle } from "lucide-react"

import { Header, PlugGrid } from "@/components"
import { useSockets } from "@/contexts"
import { VIEW_KEYS } from "@/lib"
import { api } from "@/server/client"

export const Plugs: FC<HTMLAttributes<HTMLDivElement> & { id: string; hideEmpty?: boolean }> = ({
	id,
	hideEmpty = false,
	...props
}) => {
	const { page, handle } = useSockets()

	const { data: discoveryPlugs } = api.plug.all.useQuery({
		target: "others",
		limit: 2
	})

	const { data: myPlugs } = api.plug.all.useQuery({
		target: "mine",
		limit: 4
	})

	if (page === undefined) return null

	return (
		<div {...props}>
			{(!hideEmpty || (hideEmpty && (discoveryPlugs?.length ?? 0) > 0)) && (
				<>
					<Header
						size="md"
						icon={<Puzzle size={14} className="opacity-40" />}
						label="Discover"
						nextOnClick={() =>
							handle.columns.navigate({
								id: page.id,
								key: VIEW_KEYS.DISCOVER
							})
						}
						nextLabel="See All"
					/>

					<PlugGrid id={id} from={VIEW_KEYS.HOME} plugs={discoveryPlugs} />
				</>
			)}

			{(!hideEmpty || (hideEmpty && (myPlugs?.length ?? 0) > 0)) && (
				<>
					<Header
						size="md"
						icon={<PlugZap size={14} className="opacity-40" />}
						label="My Plugs"
						nextOnClick={() =>
							handle.columns.navigate({
								id: page.id,
								key: VIEW_KEYS.MY_PLUGS
							})
						}
						nextLabel="See All"
					/>
					<PlugGrid id={id} from={VIEW_KEYS.HOME} plugs={myPlugs} />
				</>
			)}
		</div>
	)
}
