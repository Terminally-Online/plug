import { FC } from "react"

import Image from "next/image"

import { Info } from "lucide-react"

import { Button } from "@/components"
import { useFrame, usePlugs } from "@/contexts"
import {
	categories,
	formatTitle,
	getValues,
	actions as staticActions
} from "@/lib"

export const ActionItem: FC<{
	id: string
	categoryName: keyof typeof categories
	actionName: keyof (typeof staticActions)[keyof typeof categories]
	image?: boolean
}> = ({ id, categoryName, actionName, image = false }) => {
	const { handleFrame } = useFrame({
		id,
		key: `${categoryName}-${actionName}`
	})
	const { plug, actions, handle } = usePlugs(id)

	if (!plug) return null

	return (
		<>
			<div className="flex flex-row items-center gap-2">
				{image && (
					<Image
						src={`/protocols/${categoryName}.png`}
						alt={categoryName}
						width={32}
						height={32}
						className="mr-2 h-6 w-6 rounded-md"
					/>
				)}

				<Button
					variant="secondary"
					sizing="md"
					className="w-full px-6 text-left"
					onClick={() => {
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

						handleFrame()
					}}
				>
					{formatTitle(actionName)}
				</Button>

				<button className="ml-2" onClick={() => handleFrame()}>
					<Info size={14} />
				</button>
			</div>
		</>
	)
}
