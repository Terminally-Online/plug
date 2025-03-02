import { FC } from "react"

import { PencilLine, Settings } from "lucide-react"

import { Frame } from "@/components/app/frames/base"
import { Search } from "@/components/app/inputs/search"
import { Button } from "@/components/shared/buttons/button"
import { useDebounce } from "@/lib"
import { columnByIndexAtom, isFrameAtom } from "@/state/columns"
import { useAtom, useAtomValue } from "jotai"
import { plugByIdAtom, usePlugActions } from "@/state/plugs"

export const ManagePlugFrame: FC<{ index: number; item: string; from?: string }> = ({ index, item, from }) => {
	const [column] = useAtom(columnByIndexAtom(index))
	const frameKey = "manage"
	const isFrame = useAtomValue(isFrameAtom)(column, frameKey)

	const [plug] = useAtom(plugByIdAtom(item))
	const { edit, delete: deletePlug } = usePlugActions()

	const handleNameChange = (newName: string) => {
		if (!plug || !newName || newName === plug.name) return

		edit({ ...plug, name: newName, namedAt: new Date() })
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
					onClick={() => deletePlug({ plug: plug.id, index, from })}
				>
					Delete
				</Button>
			</div>
		</Frame>
	)
}
