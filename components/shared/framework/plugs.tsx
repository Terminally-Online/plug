import { FC, HTMLAttributes } from "react"

import { PlugZap, Puzzle } from "lucide-react"

import { Header, PlugGrid } from "@/components/app"
import { routes } from "@/lib"
import { api } from "@/server/client"

export const Plugs: FC<HTMLAttributes<HTMLDivElement>> = ({ ...props }) => {
	const { data: othersPlugs } = api.plug.all.useQuery({
		target: "others",
		limit: 4
	})
	const { data: myPlugs } = api.plug.all.useQuery({
		target: "mine",
		limit: 12
	})

	return (
		<div {...props}>
			<Header
				size="md"
				icon={<Puzzle size={14} className="opacity-40" />}
				label="Discover"
				nextHref={routes.app.plugs.templates}
				nextLabel="See All"
			/>
			<PlugGrid from={routes.app.plugs.index} plugs={othersPlugs} />

			<Header
				size="md"
				icon={<PlugZap size={14} className="opacity-40" />}
				label="My Plugs"
				nextHref={routes.app.plugs.mine}
				nextLabel="See All"
			/>
			<PlugGrid from={routes.app.plugs.index} plugs={myPlugs} />
		</div>
	)
}
