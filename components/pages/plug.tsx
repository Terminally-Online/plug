import { Container, Plug } from "@/components"
import { useSockets } from "@/contexts"

export const PagePlug = () => {
	const { page } = useSockets()

	if (!page || !page.item) return null

	return (
		<Container>
			<Plug id={page.id} item={page.item} />
		</Container>
	)
}
