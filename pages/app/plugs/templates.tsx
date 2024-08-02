import { Plus } from "lucide-react"

import { Container, Header, PlugsDiscover } from "@/components"
import { usePlugs } from "@/contexts/PlugProvider"
import { routes } from "@/lib"

const Page = () => {
	const { handle } = usePlugs()

	return (
		<>
			<Container>
				<Header
					size="lg"
					back={routes.app.plugs.index}
					label="Discover"
					nextOnClick={() =>
						handle.plug.add(routes.app.plugs.templates)
					}
					nextLabel={<Plus size={14} />}
				/>
			</Container>

			<PlugsDiscover />
		</>
	)
}

export default Page
