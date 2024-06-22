import { Plus, SearchIcon } from "lucide-react"

import { Container, Header } from "@/components/app"
import { PlugGrid } from "@/components/app/plugs/grid"
import { Search } from "@/components/inputs/search"
import { Tags } from "@/components/inputs/tags"
import { usePlugs } from "@/contexts/PlugProvider"
import { routes } from "@/lib/constants"

const Page = () => {
	const { search, actions } = usePlugs()

	return (
		<>
			<Container>
				<Header
					size="lg"
					back={routes.app.plugs.index}
					label="My Plugs"
					nextOnClick={() =>
						actions.plug.handleAdd(routes.app.plugs.mine)
					}
					nextLabel={<Plus size={14} className="opacity-60" />}
				/>
				<Search
					icon={<SearchIcon size={14} className="opacity-60" />}
					placeholder="Search Plugs"
					search={search}
					handleSearch={actions.plug.handleSearch}
				/>
			</Container>

			<Tags />

			<Container className="mb-[20px]">
				<PlugGrid from={routes.app.plugs.mine} />
			</Container>
		</>
	)
}

export default Page
