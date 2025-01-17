import { FC, HTMLAttributes } from "react"

import { SearchIcon } from "lucide-react"

import { Search } from "@/components/app/inputs/search"
import { Tags } from "@/components/app/inputs/tags"
import { Container } from "@/components/app/layout/container"
import { PlugGrid } from "@/components/app/plugs/grid/grid"
import { Callout } from "@/components/app/utils/callout"
import { cn, useSearch } from "@/lib"
import { COLUMNS } from "@/state/columns"
import { usePlugStore } from "@/state/plugs"

export const PlugsMine: FC<HTMLAttributes<HTMLDivElement> & { index?: number }> = ({
	index = -1,
	className,
	...props
}) => {
	const { search, tag, handleSearch, handleTag } = useSearch()
	const { plugs } = usePlugStore()

	return (
		<div className={cn("flex flex-col gap-2", className)} {...props}>
			{(search !== "" || (plugs && plugs.length > 0)) && (
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

			{plugs.some(plug => Boolean(plug)) && <Tags tag={tag} handleTag={handleTag} />}

			<Callout.EmptySearch
				isEmpty={(search !== "" || tag !== "") && plugs && plugs.length === 0}
				search={search || tag}
				handleSearch={handleSearch}
			/>

			<Container>
				<PlugGrid index={index} className="mb-4" from={COLUMNS.KEYS.MY_PLUGS} plugs={plugs} />
			</Container>

			<Callout.EmptyPlugs index={index} isEmpty={search === "" && tag === "" && plugs && plugs.length === 0} />
		</div>
	)
}
