import type { FC } from "react"

import Link from "next/link"

import { cardColors } from "@/lib"
import { Workflow } from "@/server/api/routers/plug"

type Props = { from: string; plug: Workflow }

export const PlugGridItem: FC<Props> = ({ from, plug }) => {
	const backgroundImage = cardColors[plug.color]

	return (
		<Link
			href={`/app/plugs/${plug.id}?from=${from}`}
			className="flex min-h-[128px] flex-col justify-end rounded-lg p-4 text-left text-white"
			style={{
				backgroundImage
			}}
		>
			<span className="font-bold">
				{plug.name === "" ? "Untitled Plug" : plug.name}
			</span>
		</Link>
	)
}
