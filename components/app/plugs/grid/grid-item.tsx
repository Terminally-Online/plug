import { FC } from "react"

import { useSockets } from "@/contexts"
import { cardColors, VIEW_KEYS } from "@/lib"
import { Workflow } from "@/server/api/routers/plug"

type Props = { id: string; from: string; plug: Workflow }

export const PlugGridItem: FC<Props> = ({ id, from, plug }) => {
	const { handle } = useSockets()

	const backgroundImage = cardColors[plug.color]

	return (
		<button
			onClick={() =>
				handle.columns.navigate({
					id,
					key: VIEW_KEYS.PLUG,
					item: plug.id,
					from
				})
			}
			className="flex min-h-[128px] w-full flex-col justify-end rounded-lg p-4 text-left text-white"
			style={{
				backgroundImage
			}}
		>
			<span className="font-bold">
				{plug.name === "" ? "Untitled Plug" : plug.name}
			</span>
		</button>
	)
}
