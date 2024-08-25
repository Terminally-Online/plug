import { PageActivity, PageDiscover, PageEarnings, PageHome, PageMine, PagePlug } from "@/components"
import { useSockets } from "@/contexts"
import { VIEW_KEYS } from "@/lib"

export const PageContent = () => {
	const { page } = useSockets()

	if (page === undefined) return null

	switch (page.key) {
		case VIEW_KEYS.HOME:
			return <PageHome />
		case VIEW_KEYS.ACTIVITY:
			return <PageActivity />
		case VIEW_KEYS.DISCOVER:
			return <PageDiscover />
		case VIEW_KEYS.MY_PLUGS:
			return <PageMine />
		case VIEW_KEYS.PLUG:
			return <PagePlug />
		case VIEW_KEYS.EARNINGS:
			return <PageEarnings />
		default:
			return <></>
	}
}
