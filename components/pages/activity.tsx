import { Container, Header, SocketActivity } from "@/components"
import { useSockets } from "@/contexts"

export const PageActivity = () => {
	const { page } = useSockets()

	if (page === undefined) return null

	return (
		<Container>
			<SocketActivity id={page.id} />
		</Container>
	)
}
