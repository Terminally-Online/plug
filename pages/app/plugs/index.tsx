import { PlugZap, Plus, Puzzle, ToyBrick } from "lucide-react"

import { Container, Header } from "@/components/app"
import { PlugGrid } from "@/components/app/plugs/grid"
import { usePlugs } from "@/contexts/PlugProvider"
import { routes } from "@/lib/constants"
import { api } from "@/server/client"

const Page = () => {
	const { handle } = usePlugs()

	const { data: othersPlugs } = api.plug.all.useQuery({
		target: "others",
		limit: 4
	})
	const { data: myPlugs } = api.plug.all.useQuery({
		target: "mine",
		limit: 12
	})

	return (
		<>
			<Container>
				<Header
					size="lg"
					back={routes.app.index}
					label="Plugs"
					nextOnClick={() => handle.plug.add(routes.app.plugs.index)}
					nextLabel={<Plus size={14} />}
				/>

				<Header
					size="md"
					icon={<Puzzle size={14} className="opacity-60" />}
					label="Discover"
					nextHref={routes.app.plugs.templates}
					nextLabel="See All"
				/>
				<PlugGrid from={routes.app.plugs.index} plugs={othersPlugs} />

				<Header
					size="md"
					icon={<PlugZap size={14} className="opacity-60" />}
					label="My Plugs"
					nextHref={routes.app.plugs.mine}
					nextLabel="See All"
				/>
				<PlugGrid from={routes.app.plugs.index} plugs={myPlugs} />

				{/* <Header
					size="md"
					icon={<ToyBrick size={14} className="opacity-60" />}
					label="Protocol Actions"
				/> */}
			</Container>
		</>
	)
}

export default Page
