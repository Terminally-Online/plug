import { FC } from "react"

import { Accordion, Image } from "@/components"
import { ActionSchema, formatTitle, getValues } from "@/lib"
import { useColumnStore, usePlugStore } from "@/state"

export const ActionItem: FC<{
	index: number
	item: string
	actionName: string
	protocol: string
	action: ActionSchema
	image?: boolean
}> = ({ index, item, protocol, actionName, action }) => {
	const { handle } = useColumnStore(index)
	const { plug, actions, handle: plugHandle } = usePlugStore(item)

	if (!plug) return null

	return (
		<Accordion
			onExpand={() => {
				plugHandle.action.edit({
					id: plug.id,
					actions: JSON.stringify([
						...actions,
						{
							protocol,
							action: actionName,
							values: getValues(action.schema[actionName].sentence)
						}
					])
				})
				handle.frame()
			}}
		>
			<div className="flex flex-row items-center gap-2">
				<div className="relative h-6 w-10 min-w-10">
					<Image
						src={action.metadata.icon}
						alt={"icon"}
						width={128}
						height={128}
						className="absolute left-1/2 top-1/2 mr-2 h-12 w-12 -translate-x-1/2 -translate-y-1/2 rounded-full blur-2xl filter"
					/>
					<Image
						src={action.metadata.icon}
						alt={"icon"}
						width={128}
						height={128}
						className="absolute left-1/2 top-1/2 mr-2 h-6 w-6 -translate-x-1/2 -translate-y-1/2 rounded-sm"
					/>
				</div>

				<p className="font-bold">
					<span className="opacity-40">{formatTitle(protocol)}: </span>
					{formatTitle(actionName)}
				</p>
			</div>
		</Accordion>
	)
}
