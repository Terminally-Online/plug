import { FC } from "react"

import { Workflow } from "@prisma/client"

import { Accordion } from "@/components/shared"
import { colors, formatTitle } from "@/lib"
import { COLUMNS, useColumnStore } from "@/state"

type Props = { index: number; from: string; plug: Workflow | undefined }

export const PlugGridItem: FC<Props> = ({ index, from, plug }) => {
	const { handle } = useColumnStore()

	const loading = plug === undefined

	return (
		<Accordion
			onExpand={
				plug
					? () =>
							handle.navigate({
								index,
								key: COLUMNS.KEYS.PLUG,
								item: plug.id,
								from
							})
					: undefined
			}
			loading={loading}
			className="relative flex min-h-[128px] w-full flex-col justify-end text-left"
		>
			<div
				className="absolute -bottom-full -left-1/4 h-full w-full rounded-full blur-[100px] filter"
				style={{
					backgroundColor: plug ? colors[plug.color as keyof typeof colors] : undefined
				}}
			/>

			<div className="relative">
				{plug ? (
					<span className="font-bold">{plug.name === "" ? "Untitled Plug" : formatTitle(plug.name)}</span>
				) : (
					<span className="invisible font-bold">.</span>
				)}
			</div>
		</Accordion>
	)
}
