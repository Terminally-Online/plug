import { useSession } from "next-auth/react"

import { Button, Container, Plugs, SocketAssets } from "@/components"
import { useFrame } from "@/contexts"

export const PageHome = () => {
	const { handleFrameVisible } = useFrame()
	const { data: session } = useSession()

	return (
		<>
			<Container>
				<Plugs hideEmpty={true} />

				{session?.address ? (
					<SocketAssets />
				) : (
					<>
						<p>Authenticate to proceed.</p>
						<Button onClick={() => handleFrameVisible("auth")}>
							Authenticate
						</Button>
					</>
				)}
			</Container>
		</>
	)
}
