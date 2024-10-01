import { FC } from "react"

import { Workflow } from "@/server/api/routers/plug"

import { Accordion } from "@/components/shared"
import { cn, colors, formatTitle, VIEW_KEYS } from "@/lib"
import { useColumns } from "@/state"

type Props = { index: number; from: string; plug: Workflow | undefined }

export const PlugGridItem: FC<Props> = ({ index, from, plug }) => {
	const { navigate } = useColumns()

	const loading = plug === undefined

	return (
		<Accordion
			onExpand={
				plug
					? () =>
							navigate({
								index,
								key: VIEW_KEYS.PLUG,
								item: plug.id,
								from
							})
					: undefined
			}
			className={cn(
				"relative flex min-h-[128px] w-full flex-col justify-end text-left",
				loading
					? "animate-loading bg-gradient-animated bg-[length:200%_200%]"
					: "transition-all duration-200 ease-in-out",
				// loading === false && "bg-white hover:border-white hover:bg-grayscale-0",
				loading === false ? "cursor-pointer" : "cursor-default"
			)}
		>
			<div
				className="absolute -right-1/4 -top-3/4 h-full w-full rounded-full blur-[60px] filter"
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
