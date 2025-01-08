import { FC, HTMLAttributes, useMemo } from "react"

import { ChevronLeft, SearchIcon } from "lucide-react"

import { Button, Callout, Container, Header, PlugGrid, Search, Tags } from "@/components"
import { cn, useSearch } from "@/lib"
import { api } from "@/server/client"
import { COLUMNS, useColumnStore } from "@/state"

export const PlugsDiscover: FC<HTMLAttributes<HTMLDivElement> & { index?: number }> = ({
	index = -1,
	className,
	...props
}) => {
	const { handle } = useColumnStore(index)
	const { search, tag, handleSearch, handleTag } = useSearch()
	const { data: plugs } = api.plugs.all.useQuery({
		target: "others"
	})

	const visiblePlugs = useMemo(() => {
		if (plugs === undefined || plugs.length === 0) {
			return Array(12).fill(undefined)
		}

		// Filter plugs based on search text
		return plugs.filter(plug => {
			if (!search) return true
			return plug.name.toLowerCase().includes(search.toLowerCase())
		})
	}, [plugs, search])

	return (
		<div className={cn("flex flex-col gap-2", className)} {...props}>
			<Container className="border-grayscale-100 fixed left-0 right-0 top-0 z-[10] border-b-[1px] bg-white md:hidden">
				<Header
					size="lg"
					label={
						<div className="flex flex-row items-center gap-4">
							<Button
								variant="secondary"
								className="rounded-sm p-1"
								onClick={() => handle.navigate({ index: -1, key: COLUMNS.KEYS.HOME })}
							>
								<ChevronLeft size={14} />
							</Button>
							<span className="font-bold">Discover</span>
						</div>
					}
				/>
			</Container>

			<div className="mt-16 md:mt-0">
				{plugs && plugs.length > 0 && (
					<Container className="mb-6">
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
					<PlugGrid index={index} className="mb-4" from={COLUMNS.KEYS.DISCOVER} plugs={visiblePlugs} />
				</Container>

				<Callout.EmptyPlugs index={index} isEmpty={(search === "" && plugs && plugs.length === 0) || false} />
			</div>
		</div>
	)
}