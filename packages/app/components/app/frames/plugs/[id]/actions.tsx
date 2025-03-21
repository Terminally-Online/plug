import { FC, useMemo } from "react"

import { Blocks, SearchIcon } from "lucide-react"

import { useAtom, useAtomValue } from "jotai"

import { Frame } from "@/components/app/frames/base"
import { Search } from "@/components/app/inputs/search"
import { ActionItem } from "@/components/app/plugs/actions/action-item"
import { useDebounce } from "@/lib"
import { useActions } from "@/state/actions"
import { columnByIndexAtom, isFrameAtom } from "@/state/columns"

export const ActionsFrame: FC<{ index: number; item: string }> = ({ index, item }) => {
	const [column] = useAtom(columnByIndexAtom(index))
	const frameKey = `${item}-actions`
	const isFrame = useAtomValue(isFrameAtom)(column, frameKey)

	const [actions] = useActions()

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
			icon={<Blocks size={16} className="opacity-40" />}
			label="Add Action"
			visible={isFrame}
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
								protocol={protocol}
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
