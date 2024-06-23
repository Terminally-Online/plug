import { useState } from "react"

import { PlugZap, Plus, Puzzle, SearchIcon, ToyBrick } from "lucide-react"

import { Container, Header } from "@/components/app"
import { PlugGrid } from "@/components/app/plugs/grid"
import { Search, Tags } from "@/components/inputs"
import { usePlugs } from "@/contexts/PlugProvider"
import { routes } from "@/lib/constants"

const Page = () => {
	const { search, handle } = usePlugs()

	return (
		<>
			<Container>
				<Header
					size="lg"
					back={routes.app.index}
					label="Plugs"
					nextOnClick={() => handle.plug.add(routes.app.plugs.index)}
					nextLabel={<Plus size={14} className="opacity-60" />}
				/>
				<Search
					icon={<SearchIcon size={14} className="opacity-60" />}
					placeholder="Search templates"
					search={search}
					handleSearch={handle.search}
				/>
			</Container>

			<Tags />

			<Container>
				<Header
					size="md"
					icon={<Puzzle size={14} className="opacity-60" />}
					label="Templates"
					nextHref={routes.app.plugs.templates}
					nextLabel="See All"
				/>
				<PlugGrid from={routes.app.plugs.index} count={4} />

				<Header
					size="md"
					icon={<PlugZap size={14} className="opacity-60" />}
					label="My Plugs"
					nextHref={routes.app.plugs.mine}
					nextLabel="See All"
				/>
				<PlugGrid from={routes.app.plugs.index} count={6} />

				<Header
					size="md"
					icon={<ToyBrick size={14} className="opacity-60" />}
					label="Protocol Actions"
				/>
			</Container>
		</>
	)
}

export default Page
