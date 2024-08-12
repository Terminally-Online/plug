import { useSession } from "next-auth/react"

import { Button, Container, Plugs, SocketAssets } from "@/components"
import { useFrame } from "@/contexts"

export const PageHome = () => {
	const { handleFrame } = useFrame({ id: "global" })
	const { data: session } = useSession()

	return (
		<>
			<Container>
				<Plugs hideEmpty={true} />

				{session?.address ? (
					<SocketAssets id="global" />
				) : (
					<>
						<p>Authenticate to proceed.</p>
						<Button onClick={() => handleFrame("auth")}>
							Authenticate
						</Button>
					</>
				)}
			</Container>
		</>
	)
}
