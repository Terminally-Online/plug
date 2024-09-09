import { FC, HTMLAttributes, useEffect, useMemo, useState } from "react"

import { useMotionValueEvent, useScroll } from "framer-motion"
import { SearchIcon } from "lucide-react"

import { Workflow } from "@prisma/client"

import { Callout, Container, PlugGrid, Search, Tags } from "@/components"
import { useSockets } from "@/contexts"
import { cn, useSearch, VIEW_KEYS } from "@/lib"
import { api } from "@/server/client"

export const PlugsMine: FC<HTMLAttributes<HTMLDivElement> & { id: string }> = ({ id, className, ...props }) => {
	const { address, isExternal, column } = useSockets(id)
	const { search, tag, handleSearch, handleTag } = useSearch()
	const { scrollYProgress } = useScroll()

	const [plugs, setPlugs] = useState<{
		count?: number
		plugs: Array<Workflow>
	}>({ plugs: [] })

	const representative = isExternal && column && column.viewAs ? column.viewAs.id : address

	const { fetchNextPage, isLoading } = api.plug.infinite.useInfiniteQuery(
		{
			address: isExternal && column && column.viewAs ? column.viewAs.id : address,
			search,
			tag,
			limit: 20
		},
		{
			enabled: representative !== undefined,
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
		if (plugs === undefined || plugs.count === 0) return Array(12).fill(undefined)
		return plugs.plugs
	}, [plugs])

	useMotionValueEvent(scrollYProgress, "change", latest => {
		if (!plugs || isLoading || latest < 0.8) return
		if ((plugs.count ?? 0) > plugs.plugs.length) fetchNextPage()
	})

	return (
		<div className={cn("relative flex h-full flex-col gap-2", className)} {...props}>
			{visiblePlugs.length > 0 && (
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
				<PlugGrid id={id} className="mb-4" from={VIEW_KEYS.MY_PLUGS} plugs={visiblePlugs} />
			</Container>

			<Callout.EmptyPlugs id={id} isEmpty={search === "" && tag === "" && plugs && plugs.count === 0} />
		</div>
	)
}
