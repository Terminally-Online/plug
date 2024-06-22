import { useState } from "react"

import { Earth, Gem, Plus, SearchIcon } from "lucide-react"

import { Container, Header } from "@/components/app"
import { Search } from "@/components/inputs/search"
import { Tags } from "@/components/inputs/tags"
import { usePlugs } from "@/contexts/PlugProvider"
import { routes } from "@/lib/constants"

const Page = () => {
	const { actions } = usePlugs()

	const [search, setSearch] = useState("")

	return (
		<>
			<Container>
				<Header
					size="lg"
					back={routes.app.plugs.index}
					label="Templates"
					nextOnClick={() =>
						actions.plug.handleAdd(routes.app.plugs.templates)
					}
					nextLabel={<Plus size={14} className="opacity-60" />}
				/>
				<Search
					icon={<SearchIcon size={14} className="opacity-60" />}
					placeholder="Search templates"
					search={search}
					handleSearch={setSearch}
				/>
			</Container>

			<Tags />

			<Container>
				<Header
					size="md"
					icon={<Gem size={14} className="opacity-60" />}
					label="Curated"
				/>
				{/* Show the currently curated Plugs. */}

				<Header
					size="md"
					icon={<Earth size={14} className="opacity-60" />}
					label="Community"
				/>
				{/* Show the community plugs that have been vetted for safety. */}
			</Container>
		</>
	)
}

export default Page
