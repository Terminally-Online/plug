import { Plus } from "lucide-react"

import { Container, Header } from "@/components"
import { Plugs } from "@/components/shared/framework/plugs"
import { usePlugs } from "@/contexts/PlugProvider"
import { routes } from "@/lib"

const Page = () => {
	const { handle } = usePlugs()

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

				<Plugs />
			</Container>
		</>
	)
}

export default Page
