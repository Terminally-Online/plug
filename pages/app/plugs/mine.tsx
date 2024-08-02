import { Plus } from "lucide-react"

import { Container, Header } from "@/components"
import { PlugsMine } from "@/components/app/plugs/mine"
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
					label="My Plugs"
					nextOnClick={() => handle.plug.add(routes.app.plugs.mine)}
					nextLabel={<Plus size={14} className="opacity-60" />}
				/>
			</Container>

			<PlugsMine />
		</>
	)
}

export default Page
