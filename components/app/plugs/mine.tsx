import { FC, HTMLAttributes, useMemo, useState } from "react"

import { useMotionValueEvent, useScroll } from "framer-motion"
import { SearchIcon } from "lucide-react"

import { Workflow } from "@prisma/client"

import { Callout, Container, PlugGrid, Search, Tags } from "@/components"
import { cn, useSearch } from "@/lib"
import { api } from "@/server/client"
import { COLUMNS, useColumnData, useSocket } from "@/state"

export const PlugsMine: FC<HTMLAttributes<HTMLDivElement> & { index?: number }> = ({
	index = -1,
	className,
	...props
}) => {
	const { socket } = useSocket()
	const { search, tag, handleSearch, handleTag } = useSearch()
	const { scrollYProgress } = useScroll()

	const [plugs, setPlugs] = useState<{
		count?: number
		plugs: Array<Workflow>
	}>({ plugs: [] })

	const { fetchNextPage, isLoading } = api.plugs.infinite.useInfiniteQuery(
		{
			address: socket?.id,
			search,
			tag,
			limit: 40
		},
		{
			enabled: socket?.id !== undefined,
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

	const visiblePlugs = useMemo(() => {
		if (plugs === undefined || (plugs.count === 0 && search === "")) return Array(12).fill(undefined)

		return plugs.plugs
	}, [plugs, search])

	useMotionValueEvent(scrollYProgress, "change", latest => {
		if (!plugs || isLoading || latest < 0.8) return
		if ((plugs.count ?? 0) > plugs.plugs.length) fetchNextPage()
	})

	return (
		<div className={cn("flex flex-col gap-2", className)} {...props}>
			{(search !== "" || (plugs && plugs.plugs.length > 0)) && (
				<Container>
					<Search
						icon={<SearchIcon size={14} className="opacity-60" />}
						placeholder="Search Plugs"
						search={search}
						handleSearch={handleSearch}
						clear={true}
					/>
				</Container>
			)}

			{visiblePlugs.some(plug => Boolean(plug)) && <Tags tag={tag} handleTag={handleTag} />}

			<Callout.EmptySearch
				isEmpty={(search !== "" || tag !== "") && plugs && plugs.count === 0}
				search={search || tag}
				handleSearch={handleSearch}
			/>

			<Container>
				<PlugGrid index={index} className="mb-4" from={COLUMNS.KEYS.MY_PLUGS} plugs={visiblePlugs} />
			</Container>

			<Callout.EmptyPlugs index={index} isEmpty={search === "" && tag === "" && plugs && plugs.count === 0} />
		</div>
	)
}
