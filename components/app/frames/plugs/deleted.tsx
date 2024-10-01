import { FC } from "react"

import { Trash2 } from "lucide-react"

import { Button, Frame } from "@/components"
import { usePlugs } from "@/contexts"
import { useColumns } from "@/state"

export const DeletedFrame: FC<{ index: number }> = ({ index }) => {
	const { isFrame } = useColumns(index, "deleted")
	const { handle } = usePlugs()

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
