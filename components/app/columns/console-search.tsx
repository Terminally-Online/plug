import { FC, HTMLAttributes } from "react"

import { Plug, SearchIcon } from "lucide-react"

import { useDebounce, VIEW_KEYS } from "@/lib"
import { api } from "@/server/client"

import { Search } from "../inputs"
import { PlugGrid } from "../plugs"

export const ConsoleSearch: FC<
	HTMLAttributes<HTMLDivElement> & { id: string }
> = ({ id, ...props }) => {
	const [search, debounced, handleSearch] = useDebounce("", 500)

	const { data: results } = api.misc.search.useQuery(debounced, {
		enabled: search !== ""
	})

	return (
		<div {...props}>
			<Search
				className="mb-4"
				icon={<SearchIcon size={14} className="opacity-60" />}
				placeholder="Search protocols, actions, or assets"
				search={search}
				handleSearch={handleSearch}
			/>

			{results && results.plugs.length > 0 && (
				<div className="flex flex-col gap-2">
					<p className="flex flex-row items-center gap-2 font-bold">
						<Plug size={14} className="opacity-40" />
						<span>Plugs</span>
					</p>

					<PlugGrid
						id={id}
						from={VIEW_KEYS.SEARCH}
						plugs={results.plugs}
					/>
				</div>
			)}

			{/* {results && results.collectibles.length > 0 && (
				<SocketCollectionList
					id={id}
					collectibles={results.collectibles}
				/>
			)} */}

			{/* {results && results.tokens.length > 0 && (
				<SocketTokenList id={id} tokens={results.tokens} />
			)} */}
		</div>
	)
}
