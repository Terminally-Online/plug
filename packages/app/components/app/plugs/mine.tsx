import { FC, HTMLAttributes, useMemo } from "react"

import { useMotionValueEvent, useScroll } from "framer-motion"
import { SearchIcon } from "lucide-react"

import { useAtom } from "jotai"

import { Search } from "@/components/app/inputs/search"
import { Tags } from "@/components/app/inputs/tags"
import { Container } from "@/components/app/layout/container"
import { PlugGrid } from "@/components/app/plugs/grid/grid"
import { Callout } from "@/components/app/utils/callout"
import { cn, useSearch } from "@/lib"
import { api } from "@/server/client"
import { useSocket } from "@/state/authentication"
import { COLUMNS } from "@/state/columns"
import { plugOrderAtom, plugsMapAtom } from "@/state/plugs"

export const PlugsMine: FC<HTMLAttributes<HTMLDivElement> & { index?: number }> = ({
	index = -1,
	className,
	...props
}) => {
	const { socket } = useSocket()
	const { search, tag, handleSearch, handleTag } = useSearch()
	const { scrollYProgress } = useScroll()

	const [plugsMap, setPlugsMap] = useAtom(plugsMapAtom)
	const [order, setOrder] = useAtom(plugOrderAtom)

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
				// Update both map and order
				const newPlugsMap: Record<string, any> = {}
				const newOrder: string[] = []

				data.pages.forEach(page => {
					page.plugs.forEach(plug => {
						newPlugsMap[plug.id] = plug
						if (!order.includes(plug.id)) {
							newOrder.push(plug.id)
						}
					})
				})

				setPlugsMap(prev => ({ ...prev, ...newPlugsMap }))
				if (newOrder.length > 0) {
					setOrder(prev => [...prev, ...newOrder])
				}
			}
		}
	)

	const visiblePlugs = useMemo(() => {
		if (!socket?.id || (Object.keys(plugsMap).length === 0 && search === "")) {
			return Array(12).fill(undefined)
		}

		return order
			.map(id => plugsMap[id])
			.filter(plug => {
				if (!plug) return false
				if (plug.socketId !== socket.id) return false
				if (search && !plug.name.toLowerCase().includes(search.toLowerCase())) return false
				if (tag !== "" && !plug.tags?.includes(tag)) return false
				return true
			})
	}, [plugsMap, order, socket?.id, search, tag])

	useMotionValueEvent(scrollYProgress, "change", latest => {
		if (isLoading || latest < 0.8) return
		const totalPlugs = Object.values(plugsMap).filter(p => p.socketId === socket?.id).length
		if (totalPlugs > visiblePlugs.length) fetchNextPage()
	})

	const isEmpty = visiblePlugs.length === 0 && search === "" && tag === ""
	const isEmptySearch = (search !== "" || tag !== "") && visiblePlugs.length === 0

	return (
		<div className={cn("flex flex-col gap-2", className)} {...props}>
			{(search !== "" || visiblePlugs.some(Boolean)) && (
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

			<Callout.EmptySearch isEmpty={isEmptySearch} search={search || tag} handleSearch={handleSearch} />

			<Container>
				<PlugGrid index={index} className="mb-4" from={COLUMNS.KEYS.MY_PLUGS} plugs={visiblePlugs} />
			</Container>

			<Callout.EmptyPlugs index={index} isEmpty={isEmpty} />
		</div>
	)
}
