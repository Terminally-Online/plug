import type { FC } from "react"

import Link from "next/link"

import { colors } from "@/lib/constants"
import { Workflow } from "@/server/api/routers/plug"

type Props = { from: string; plug: Workflow }

export const PlugGridItem: FC<Props> = ({ from, plug }) => (
	<Link
		href={`/app/plugs/${plug.id}?from=${from}`}
		className="flex min-h-[128px] flex-col justify-end rounded-lg p-4 text-left text-white"
		style={{ backgroundColor: colors[plug.color as keyof typeof colors] }}
	>
		<span className="font-bold">
			{plug.name === "" ? "Untitled Plug" : plug.name}
		</span>
	</Link>
)
