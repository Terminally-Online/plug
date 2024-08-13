import { PlugsDiscover } from "@/components"
import { useSockets } from "@/contexts"

export const PageDiscover = () => {
	const { page } = useSockets()

	if (!page) return null

	return <PlugsDiscover id={page.id} />
}
