import { FC, HTMLAttributes, useEffect, useMemo, useState } from "react"

import { useMotionValueEvent, useScroll } from "framer-motion"
import { SearchIcon } from "lucide-react"

import { api } from "@/server/client"

import { Workflow } from "@prisma/client"

import { Callout, Container, PlugGrid, Search, Tags } from "@/components"
import { cn, useSearch, VIEW_KEYS } from "@/lib"

export const PlugsDiscover: FC<HTMLAttributes<HTMLDivElement> & { index?: number }> = ({
	index = -1,
	className,
	...props
}) => {
	// const { scrollYProgress } = useScroll()
	const { search, tag, handleSearch, handleTag } = useSearch()

	// const [plugs, setPlugs] = useState<{
	// 	count?: number
	// 	plugs: Array<Workflow>
	// }>({ plugs: [] })

	const { data: plugs } = api.plug.all.useQuery({
		target: "others"
	})

	const visiblePlugs = useMemo(() => {
		if (plugs === undefined || plugs.length === 0) return Array(12).fill(undefined)
		return plugs
	}, [plugs])

	// useMotionValueEvent(scrollYProgress, "change", latest => {
	// 	if (!plugs || isLoading || latest < 0.8) return
	// 	if ((plugs.count ?? 0) > plugs.plugs.length) fetchNextPage()
	// })

	return (
		<div className={cn("relative flex h-full flex-col gap-2", className)} {...props}>
			{plugs && plugs.length > 0 && (
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
				isEmpty={(search !== "" && plugs && plugs.length === 0) || false}
				search={search}
				handleSearch={handleSearch}
			/>

			<Container>
				<PlugGrid index={index} className="mb-4" from={VIEW_KEYS.DISCOVER} plugs={visiblePlugs} />
			</Container>

			<Callout.EmptyPlugs index={index} isEmpty={(search === "" && plugs && plugs.length === 0) || false} />
		</div>
	)
}
