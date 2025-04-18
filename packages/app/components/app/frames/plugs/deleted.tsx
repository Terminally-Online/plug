import { FC } from "react"

import { Trash2 } from "lucide-react"

import { useAtom, useAtomValue } from "jotai"

import { Frame } from "@/components/app/frames/base"
import { Button } from "@/components/shared/buttons/button"
import { columnByIndexAtom, isFrameAtom } from "@/state/columns"
import { usePlugActions } from "@/state/plugs"

export const DeletedFrame: FC<{ index: number }> = ({ index }) => {
	const [column] = useAtom(columnByIndexAtom(index))
	const isFrame = useAtomValue(isFrameAtom)(column, "deleted")

	const { add } = usePlugActions()

	return (
		<Frame index={index} className="z-[2]" icon={<Trash2 size={18} />} label="Plug Deleted" visible={isFrame}>
			<p className="w-full opacity-60">
				This content you were viewing is no longer available. It may have been deleted or made private by the
				creator.
			</p>

			<Button className="mt-4 w-full" onClick={() => add()}>
				Create New
			</Button>
		</Frame>
	)
}
