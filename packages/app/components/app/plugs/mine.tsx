import { FC, HTMLAttributes, useMemo } from "react"

import { SearchIcon } from "lucide-react"

import { useAtomValue } from "jotai"

import { Search } from "@/components/app/inputs/search"
import { Tags } from "@/components/app/inputs/tags"
import { Container } from "@/components/app/layout/container"
import { PlugGrid } from "@/components/app/plugs/grid/grid"
import { Callout } from "@/components/app/utils/callout"
import { cn, useSearch } from "@/lib"
import { COLUMNS } from "@/state/columns"
import { plugsAtom } from "@/state/plugs"
import { useSocket } from "@/state/authentication"
import { PLACEHOLDER_PLUGS } from "@/lib/constants/placeholder/plugs"

export const PlugsMine: FC<HTMLAttributes<HTMLDivElement> & { index?: number }> = ({
	index = -1,
	className,
	...props
}) => {
	const { search, tag, handleSearch, handleTag } = useSearch()
	const { socket } = useSocket()

	const plugs = useAtomValue(plugsAtom)

	const visiblePlugs = useMemo(() => {
		const my = plugs.filter(plug => plug.socketId === socket.id)

		if (search !== "" && my.length == 0) return Array(12).fill(undefined)

		if (!my || my.length === 0) return PLACEHOLDER_PLUGS

		if (search) return my

		return my.filter(plug => plug.name.toLowerCase().includes(search.toLowerCase()))
	}, [plugs, search, socket.id])

	return (
		<div className={cn("flex flex-col gap-2", className)} {...props}>
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

			<Container>
				<PlugGrid index={index} className="mb-4" from={COLUMNS.KEYS.MY_PLUGS} plugs={visiblePlugs} />
			</Container>

			<Callout.EmptySearch
				isEmpty={(search !== "" || tag !== "") && visiblePlugs && visiblePlugs.length === 0}
				search={search || tag}
				handleSearch={handleSearch}
			/>
			<Callout.EmptyPlugs index={index} isEmpty={(search === "" && plugs && plugs.length === 0) || false} />
		</div>
	)
}
