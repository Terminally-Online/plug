import { useState } from "react"

import { useMotionValueEvent, useScroll } from "framer-motion"
import { Plus, SearchIcon } from "lucide-react"

import { Workflow } from "@prisma/client"

import { Container, Header, PlugGrid, Search, Tags } from "@/components"
import { usePlugs } from "@/contexts/PlugProvider"
import { routes, useSearch } from "@/lib"
import { api } from "@/server/client"

// Notes: We do not use the contextual plugs here because we want search without
// manipulating the global context as well as we want to do incremental
// rendering. If we utilize the global state and someone had hundreds of workflows
// this page would take an eternity to load.
//
// Conclusion: The tradeoff we make by doing this is that this list will not live update.
// Realistically, that is not super important because the underlying state will have regardless.
const Page = () => {
	const { scrollYProgress } = useScroll()
	const {
		search,
		debouncedSearch,
		tag,
		handleSearch,
		handleTag,
		handleReset
	} = useSearch()
	const { handle } = usePlugs()

	const [plugs, setPlugs] = useState<{
		count?: number
		plugs: Array<Workflow>
	}>({ plugs: [] })

	const { fetchNextPage, isLoading } = api.plug.infinite.useInfiniteQuery(
		{
			mine: true,
			search: debouncedSearch,
			tag,
			limit: 20
		},
		{
			getNextPageParam(lastPage) {
				return lastPage.nextCursor
			},
			onSuccess(data) {
				setPlugs(() => ({
					count: data.pages[data.pages.length - 1].count,
					plugs: data.pages.flatMap(page => page.plugs)
				}))
			}
		}
	)

	useMotionValueEvent(scrollYProgress, "change", latest => {
		if (!plugs || isLoading || latest < 0.8) return

		if ((plugs.count ?? 0) > plugs.plugs.length) {
			fetchNextPage()
		}
	})

	return (
		<>
			<Container>
				<Header
					size="lg"
					back={routes.app.plugs.index}
					label="My Plugs"
					nextOnClick={() => handle.plug.add(routes.app.plugs.mine)}
					nextLabel={<Plus size={14} className="opacity-60" />}
				/>

				<Search
					icon={<SearchIcon size={14} className="opacity-60" />}
					placeholder="Search Plugs"
					search={search}
					handleSearch={handleSearch}
					clear={true}
				/>
			</Container>

			<Tags tag={tag} handleTag={handleTag} />

			<Container>
				<PlugGrid
					className="mb-4"
					from={routes.app.plugs.mine}
					search={search || tag}
					handleReset={handleReset}
					plugs={plugs.plugs}
				/>
			</Container>
		</>
	)
}

export default Page
