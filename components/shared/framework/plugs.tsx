import { FC, HTMLAttributes } from "react"

import { PlugZap, Puzzle } from "lucide-react"

import { api } from "@/server/client"

import { Header, PlugGrid } from "@/components"
import { VIEW_KEYS } from "@/lib"
import { useColumns } from "@/state"

export const Plugs: FC<HTMLAttributes<HTMLDivElement> & { index?: number; hideEmpty?: boolean }> = ({
	index = -1,
	hideEmpty = false,
	...props
}) => {
	const { navigate } = useColumns(index)

	const { data: discoveryPlugs } = api.plug.all.useQuery({
		target: "others",
		limit: 2
	})

	const { data: myPlugs } = api.plug.all.useQuery({
		target: "mine",
		limit: 4
	})

	return (
		<div {...props}>
			<Header
				size="md"
				icon={<Puzzle size={14} className="opacity-40" />}
				label="Discover"
				nextOnClick={() =>
					navigate({
						index,
						key: VIEW_KEYS.DISCOVER
					})
				}
				nextLabel="See All"
			/>

			<PlugGrid index={index} from={VIEW_KEYS.HOME} plugs={discoveryPlugs} />

			<Header
				size="md"
				icon={<PlugZap size={14} className="opacity-40" />}
				label="My Plugs"
				nextOnClick={() =>
					navigate({
						index,
						key: VIEW_KEYS.MY_PLUGS
					})
				}
				nextLabel="See All"
			/>

			<PlugGrid index={index} from={VIEW_KEYS.HOME} plugs={myPlugs} />
		</div>
	)
}
