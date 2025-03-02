import { FC, HTMLAttributes, useMemo } from "react"

import { SearchIcon } from "lucide-react"

import { Search } from "@/components/app/inputs/search"
import { Tags } from "@/components/app/inputs/tags"
import { Container } from "@/components/app/layout/container"
import { PlugGrid } from "@/components/app/plugs/grid/grid"
import { Callout } from "@/components/app/utils/callout"
import { cn, useSearch } from "@/lib"
import { COLUMNS } from "@/state/columns"
import { plugsAtom } from "@/state/plugs"
import { useAtomValue } from "jotai"

export const PlugsMine: FC<HTMLAttributes<HTMLDivElement> & { index?: number }> = ({
	index = -1,
	className,
	...props
}) => {
	const { search, tag, handleSearch, handleTag } = useSearch()

	const plugs = useAtomValue(plugsAtom)

	const visiblePlugs = useMemo(() => {
		if (!plugs || plugs.length === 0)
			return Array(12).fill(undefined)

		if (!search) return plugs

		return plugs.filter(plug => plug.name.toLowerCase().includes(search.toLowerCase()))
	}, [plugs, search])

	return (
		<div className={cn("flex flex-col gap-2", className)} {...props}>
			{(plugs && plugs.length > 0) && (
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
				isEmpty={(search !== "" || tag !== "") && visiblePlugs && visiblePlugs.length === 0}
				search={search || tag}
				handleSearch={handleSearch}
			/>

			<Container>
				<PlugGrid index={index} className="mb-4" from={COLUMNS.KEYS.MY_PLUGS} plugs={visiblePlugs} />
			</Container>

			<Callout.EmptyPlugs index={index} isEmpty={(search === "" && plugs && plugs.length === 0) || false} />
		</div>
	)
}
