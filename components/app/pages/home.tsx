import { AccountFrame, Container, SocketAssets } from "@/components"
import { Plugs } from "@/components/shared/framework/plugs"

export const PageHome = () => (
	<>
		<Container>
			<Plugs hideEmpty={true} />
			<SocketAssets />
		</Container>

		<AccountFrame />
	</>
)
