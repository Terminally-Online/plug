import Image from "next/image"
import { FC } from "react"

import { Accordion } from "@/components"
import { usePlugs } from "@/contexts"
import { categories, formatTitle, getValues, actions as staticActions } from "@/lib"
import { useFrame } from "@/state"

export const ActionItem: FC<{
	index: number
	item: string
	categoryName: keyof typeof categories
	actionName: keyof (typeof staticActions)[keyof typeof categories]
	image?: boolean
}> = ({ index, item, categoryName, actionName, image = false }) => {
	const { handleFrame } = useFrame({ index })
	const { plug, actions, handle } = usePlugs(item)

	if (!plug) return null

	return (
		<Accordion
			onExpand={() => {
				handle.action.edit({
					id: plug.id,
					actions: JSON.stringify([
						...actions,
						{
							categoryName,
							actionName,
							values: getValues(categoryName, actionName)
						}
					])
				})
				handleFrame("")
			}}
		>
			<div className="flex flex-row items-center gap-2">
				{image && (
					<div className="relative h-6 w-10 min-w-10">
						<Image
							src={`/protocols/${categoryName}.png`}
							alt={categoryName}
							width={128}
							height={128}
							className="absolute left-1/2 top-1/2 mr-2 h-12 w-12 -translate-x-1/2 -translate-y-1/2 rounded-full blur-2xl filter"
						/>
						<Image
							src={`/protocols/${categoryName}.png`}
							alt={categoryName}
							width={128}
							height={128}
							className="absolute left-1/2 top-1/2 mr-2 h-6 w-6 -translate-x-1/2 -translate-y-1/2 rounded-sm"
						/>
					</div>
				)}

				<p className="font-bold">{formatTitle(actionName)}</p>
			</div>
		</Accordion>
	)
}
