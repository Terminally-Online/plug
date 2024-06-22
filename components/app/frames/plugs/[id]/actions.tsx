import { useState } from "react"

import { Blocks, SearchIcon } from "lucide-react"

import { ActionList } from "@/components/app/plugs/actions"
import { Search } from "@/components/inputs"
import { useFrame } from "@/contexts"

import { Frame } from "../../base"

export const ActionsFrame = () => {
	const { frameVisible, handleFrameVisible } = useFrame()
	const [search, setSearch] = useState("")
	return (
		<Frame
			className="scrollbar-hide z-[1] h-[calc(100vh-80px)] overflow-y-auto"
			icon={<Blocks size={18} className="opacity-60" />}
			label="Add Action"
			visible={frameVisible === "actions"}
			handleVisibleToggle={() => handleFrameVisible(undefined)}
		>
			<div className="flex flex-col gap-8">
				<Search
					icon={<SearchIcon size={14} className="opacity-60" />}
					placeholder="Search protocols and actions"
					search={search}
					handleSearch={setSearch}
				/>
				<ActionList
					handleNestedToggle={() => handleFrameVisible(undefined)}
				/>
			</div>
		</Frame>
	)
}
