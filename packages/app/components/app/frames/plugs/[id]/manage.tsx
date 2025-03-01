import { FC } from "react"

import { PencilLine, Settings } from "lucide-react"

import { Frame } from "@/components/app/frames/base"
import { Search } from "@/components/app/inputs/search"
import { Button } from "@/components/shared/buttons/button"
import { useDebounce } from "@/lib"
import { columnByIndexAtom, isFrameAtom } from "@/state/columns"
import { usePlugStore } from "@/state/plugs"
import { useAtom, useAtomValue } from "jotai"

export const ManagePlugFrame: FC<{ index: number; item: string; from?: string }> = ({ index, item, from }) => {
	const [column] = useAtom(columnByIndexAtom(index))
	const frameKey = "manage"
	const isFrame = useAtomValue(isFrameAtom)(column, frameKey)

	const { plug, handle } = usePlugStore(item)

	const handleNameChange = (newName: string) => {
		if (!plug || !newName || newName === plug.name) return

		handle.plug.edit({ ...plug, name: newName, namedAt: new Date() })
	}

	const [name, _, handleName] = useDebounce(plug?.name ?? "", 1000, handleNameChange)

	if (!plug) return null

	return (
		<Frame
			index={index}
			className="z-[2]"
			icon={<Settings size={18} className="opacity-40" />}
			label="Manage Plug"
			visible={isFrame}
			hasChildrenPadding={false}
		>
			<div className="mb-4 flex flex-col gap-4 px-6">
				<Search
					icon={<PencilLine size={14} />}
					placeholder="Plug name"
					search={name}
					handleSearch={handleName}
				/>

				<Button
					variant="destructive"
					className="w-full py-4"
					onClick={() => handle.plug.delete({ plug: plug.id, index, from })}
				>
					Delete
				</Button>
			</div>
		</Frame>
	)
}
