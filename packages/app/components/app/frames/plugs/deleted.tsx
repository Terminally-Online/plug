import { FC } from "react"

import { Trash2 } from "lucide-react"

import { Frame } from "@/components/app/frames/base"
import { Button } from "@/components/shared/buttons/button"
import { useColumnStore } from "@/state/columns"
import { usePlugStore } from "@/state/plugs"

export const DeletedFrame: FC<{ index: number }> = ({ index }) => {
	const { isFrame } = useColumnStore(index, "deleted")
	const { handle } = usePlugStore()

	return (
		<Frame index={index} className="z-[2]" icon={<Trash2 size={18} />} label="Plug Deleted" visible={isFrame}>
			<p className="w-full opacity-60">
				This content you were viewing is no longer available. It may have been deleted or made private by the
				creator.
			</p>

			<Button className="mt-4 w-full" onClick={() => handle.plug.add()}>
				Create New
			</Button>
		</Frame>
	)
}
