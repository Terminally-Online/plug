import { PlugsMine } from "@/components"
import { useSockets } from "@/contexts"

export const PageMine = () => {
	const { page } = useSockets()

	if (!page) return null

	return <PlugsMine id={page.id} />
}
