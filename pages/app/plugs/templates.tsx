import { useState } from "react"

import { useMotionValueEvent, useScroll } from "framer-motion"
import { Earth, Gem, Plus, SearchIcon } from "lucide-react"

import { Workflow } from "@prisma/client"

import { Container, Header, PlugGrid, Search, Tags } from "@/components"
import { usePlugs } from "@/contexts/PlugProvider"
import { routes, useSearch } from "@/lib"
import { api } from "@/server/client"

const Page = () => {
	const { scrollYProgress } = useScroll()
	const { handle } = usePlugs()
	const {
		search,
		debouncedSearch,
		tag,
		handleSearch,
		handleTag,
		handleReset
	} = useSearch()

	const [communityPlugs, setCommunityPlugs] = useState<{
		count?: number
		plugs: Array<Workflow>
	}>({ plugs: [] })

	const { data: curatedPlugs } = api.plug.all.useQuery({
		target: "curated",
		limit: 4
	})
	const { fetchNextPage, isLoading } = api.plug.infinite.useInfiniteQuery(
		{
			search: debouncedSearch,
			tag,
			limit: 20
		},
		{
			getNextPageParam(lastPage) {
				return lastPage.nextCursor
			},
			onSuccess(data) {
				setCommunityPlugs(() => ({
					count: data.pages[data.pages.length - 1].count,
					plugs: data.pages.flatMap(page => page.plugs)
				}))
			}
		}
	)

	useMotionValueEvent(scrollYProgress, "change", latest => {
		if (!communityPlugs || isLoading || latest < 0.8) return

		if ((communityPlugs.count ?? 0) > communityPlugs.plugs.length) {
			fetchNextPage()
		}
	})

	return (
		<>
			<Container>
				<Header
					size="lg"
					back={routes.app.plugs.index}
					label="Discover"
					nextOnClick={() =>
						handle.plug.add(routes.app.plugs.templates)
					}
					nextLabel={<Plus size={14} />}
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
				{!search && !tag && curatedPlugs && curatedPlugs.length > 0 && (
					<>
						<Header
							size="md"
							icon={<Gem size={14} className="opacity-40" />}
							label="Curated"
						/>
						<PlugGrid
							from={routes.app.plugs.templates}
							count={4}
							plugs={curatedPlugs}
						/>
					</>
				)}

				<Header
					size="md"
					icon={<Earth size={14} className="opacity-40" />}
					label="Community"
				/>
				<PlugGrid
					className="mb-4"
					from={routes.app.plugs.templates}
					search={search || tag}
					handleReset={handleReset}
					plugs={communityPlugs.plugs}
				/>
			</Container>
		</>
	)
}

export default Page
