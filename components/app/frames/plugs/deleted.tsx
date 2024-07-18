import { Trash2 } from "lucide-react"

import { Button, Frame } from "@/components"
import { useFrame, usePlugs } from "@/contexts"

export const DeletedFrame = () => {
	const { frameVisible } = useFrame()
	const { plug, handle } = usePlugs()

	if (!plug) return null

	return (
		<Frame
			className="z-[2]"
			icon={<Trash2 size={18} />}
			label="Plug Deleted"
			visible={frameVisible === "deleted"}
		>
			<p className="w-full opacity-60">
				This content you were viewing is no longer available. It may
				have been deleted or made private by the creator.
			</p>

			<Button className="mt-4 w-full" onClick={() => handle.plug.add()}>
				Create New
			</Button>
		</Frame>
	)
}
