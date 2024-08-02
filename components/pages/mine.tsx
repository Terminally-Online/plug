import { FC, HTMLAttributes, useState } from "react"

import { useMotionValueEvent, useScroll } from "framer-motion"
import { SearchIcon } from "lucide-react"

import { Workflow } from "@prisma/client"

import { Container, PlugGrid, Search, Tags } from "@/components"
import { useSearch } from "@/lib"
import { api } from "@/server/client"

export const PageMine: FC<
	HTMLAttributes<HTMLDivElement> & { column?: boolean }
> = ({ column = false, ...props }) => {
	const { scrollYProgress } = useScroll()
	const {
		search,
		debouncedSearch,
		tag,
		handleSearch,
		handleTag,
		handleReset
	} = useSearch()

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
		<div {...props}>
			<Container>
				<Search
					icon={<SearchIcon size={14} className="opacity-60" />}
					placeholder="Search Plugs"
					search={search}
					handleSearch={handleSearch}
					clear={true}
				/>
			</Container>

			{(plugs?.count ?? 0) > 0 && (
				<Tags tag={tag} handleTag={handleTag} />
			)}

			<Container>
				<PlugGrid
					className="mb-4"
					from={"mine"}
					search={search || tag}
					handleReset={handleReset}
					plugs={plugs.plugs}
				/>
			</Container>
		</div>
	)
}
