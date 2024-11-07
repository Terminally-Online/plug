import { FC, useMemo } from "react"

import { Blocks, SearchIcon } from "lucide-react"

import { ActionItem, Frame, Search } from "@/components"
import { usePlugs } from "@/contexts"
import { useDebounce } from "@/lib"
import { useActions, useColumns } from "@/state"

export const ActionsFrame: FC<{ index: number; item: string }> = ({ index, item }) => {
	const { column, isFrame } = useColumns(index, `${index}-${item}-actions`)
	const [actions] = useActions()
	const { actions: plugActions } = usePlugs(item)

	const [search, debouncedSearch, handleDebounce] = useDebounce("")

	const allFilteredActions = useMemo(
		() =>
			Object.keys(actions).flatMap(protocol => {
				if (protocol.toLowerCase().includes(debouncedSearch.toLowerCase())) {
					return Object.keys(actions[protocol].schema).map(action => ({
						protocol: protocol,
						action: action
					}))
				}

				return Object.keys(actions[protocol].schema)
					.filter(actionName => actionName.toLowerCase().includes(debouncedSearch.toLowerCase()))
					.map(action => ({
						protocol,
						action
					}))
			}),
		[actions, debouncedSearch]
	)

	if (!column) return null

	return (
		<Frame
			index={index}
			icon={<Blocks size={18} className="opacity-60" />}
			label="Add Action"
			visible={plugActions.length === 0 || isFrame}
			hasChildrenPadding={false}
		>
			<div className="flex flex-col gap-4 px-6">
				<Search
					icon={<SearchIcon size={14} />}
					placeholder="Search protocols and actions"
					search={search}
					handleSearch={handleDebounce}
					clear={true}
				/>

				<div className="mb-4 flex flex-col gap-2">
					{allFilteredActions.length > 0 ? (
						allFilteredActions.map(({ protocol, action }) => (
							<ActionItem
								key={`${protocol}-${action}`}
								index={index}
								item={item}
								actionName={action}
								action={actions[protocol]}
								image={true}
							/>
						))
					) : (
						<>
							<div className="mx-auto my-8 flex h-full max-w-[80%] flex-col gap-2 text-center">
								<p className="text-lg font-bold">No matching actions.</p>
								<p className="font-bold opacity-40">Search for another action or protocol.</p>
							</div>
						</>
					)}
				</div>
			</div>
		</Frame>
	)
}
