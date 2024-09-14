import { FC } from "react"

import { Workflow } from "@/server/api/routers/plug"

import { useSockets } from "@/contexts"
import { cardColors, cn, VIEW_KEYS } from "@/lib"

type Props = { id: string; from: string; plug: Workflow | undefined }

export const PlugGridItem: FC<Props> = ({ id, from, plug }) => {
	const { handle } = useSockets()

	const loading = plug === undefined
	const backgroundImage = plug ? cardColors[plug.color] : undefined

	return (
		<button
			onClick={
				plug
					? () =>
							handle.columns.navigate({
								id,
								key: VIEW_KEYS.PLUG,
								item: plug.id,
								from
							})
					: undefined
			}
			className={cn(
				"flex min-h-[128px] w-full flex-col justify-end rounded-lg p-4 text-left text-white",
				loading
					? "animate-loading bg-gradient-animated bg-[length:200%_200%]"
					: "transition-all duration-200 ease-in-out",
				loading === false && "bg-white hover:border-white hover:bg-grayscale-0",
				loading === false ? "cursor-pointer" : "cursor-default"
			)}
			style={{
				backgroundImage
			}}
		>
			{plug === undefined ? (
				<span className="invisible font-bold">.</span>
			) : (
				<span className="font-bold">{plug.name === "" ? "Untitled Plug" : plug.name}</span>
			)}
		</button>
	)
}
