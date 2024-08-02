import { AccountFrame, Container, Plugs, SocketAssets } from "@/components"

export const PageHome = () => (
	<>
		<Container>
			<Plugs hideEmpty={true} />
			<SocketAssets />
		</Container>

		<AccountFrame />
	</>
)
