import { FC, HTMLAttributes } from "react"

import { PlugZap, Puzzle } from "lucide-react"

import { Header, PlugGrid } from "@/components/app"
import { usePage } from "@/contexts"
import { routes } from "@/lib"
import { api } from "@/server/client"

export const Plugs: FC<
	HTMLAttributes<HTMLDivElement> & { hideEmpty?: boolean }
> = ({ hideEmpty = false, ...props }) => {
	const { handlePage } = usePage()

	const { data: discoveryPlugs } = api.plug.all.useQuery({
		target: "others",
		limit: 4
	})
	const { data: myPlugs } = api.plug.all.useQuery({
		target: "mine",
		limit: 12
	})

	return (
		<div {...props}>
			{!hideEmpty ||
				(hideEmpty && (discoveryPlugs?.length ?? 0) > 0 && (
					<>
						<Header
							size="md"
							icon={<Puzzle size={14} className="opacity-40" />}
							label="Discover"
							nextOnClick={() => handlePage("discover")}
							nextLabel="See All"
						/>

						<PlugGrid
							from={routes.app.index}
							plugs={discoveryPlugs}
						/>
					</>
				))}

			{!hideEmpty ||
				(hideEmpty && (myPlugs?.length ?? 0) > 0 && (
					<>
						<Header
							size="md"
							icon={<PlugZap size={14} className="opacity-40" />}
							label="My Plugs"
							nextHref={routes.app.plugs.mine}
							nextLabel="See All"
						/>
						<PlugGrid from={routes.app.index} plugs={myPlugs} />
					</>
				))}
		</div>
	)
}
