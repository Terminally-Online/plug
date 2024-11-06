import { FC } from "react"

import { Accordion, Image } from "@/components"
import { usePlugs } from "@/contexts"
import { ActionSchema, formatTitle, getValues } from "@/lib"
import { useColumns } from "@/state"

export const ActionItem: FC<{
	index: number
	item: string
	actionName: string
	action: ActionSchema
	image?: boolean
}> = ({ index, item, actionName, action }) => {
	const { plug, actions, handle } = usePlugs(item)
	const { frame } = useColumns(index)

	if (!plug) return null

	return (
		<Accordion
			onExpand={() => {
				handle.action.edit({
					id: plug.id,
					actions: JSON.stringify([
						...actions,
						{
							protocol: action.schema.protocol,
							action: actionName,
							values: getValues(action.schema[actionName].sentence)
						}
					])
				})
				frame()
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

				<p className="font-bold">{formatTitle(actionName)}</p>
			</div>
		</Accordion>
	)
}
