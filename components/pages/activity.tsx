import { FileCog } from "lucide-react"

import { Container, Header, SocketActivity } from "@/components"

export const PageActivity = () => (
	<Container>
		<Header
			size="md"
			icon={<FileCog size={14} className="opacity-40" />}
			label="Runs"
		/>

		<SocketActivity id="global" />
	</Container>
)
