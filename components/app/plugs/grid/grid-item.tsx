import type { FC } from "react"

import Link from "next/link"

import { usePage } from "@/contexts"
import { cardColors } from "@/lib"
import { Workflow } from "@/server/api/routers/plug"

type Props = { from: string; plug: Workflow }

export const PlugGridItem: FC<Props> = ({ from, plug }) => {
	const { handlePage } = usePage()

	const backgroundImage = cardColors[plug.color]

	return (
		<button
			onClick={() => handlePage({ key: "plug", id: plug.id, from })}
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
