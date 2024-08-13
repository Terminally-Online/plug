import { useSession } from "next-auth/react"

import { Button, Container, Plugs, SocketAssets } from "@/components"
import { useFrame, useSockets } from "@/contexts"

export const PageHome = () => {
	const { page } = useSockets()
	const { handleFrame } = useFrame({ id: page?.id })
	const { data: session } = useSession()

	if (page === undefined) return null

	return (
		<>
			<Container>
				<Plugs id={page.id} hideEmpty={true} />

				{session?.address ? (
					<SocketAssets id={page.id} />
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
