import { FC } from "react"

import { Sentence } from "@/components"
import { Action, Actions } from "@/lib"
import { usePlugData } from "@/state"

export const ActionPreview: FC<{ index: number; item: string; actions?: Actions }> = ({ index, item, actions }) => {
	const { actions: plugActions } = usePlugData(item)

	actions = actions ?? plugActions

	return (
		<div className="flex flex-col gap-2">
			{actions.map((action, actionIndex) => (
				<Sentence
					key={actionIndex}
					index={index}
					item={item}
					actionIndex={actionIndex}
					action={action}
					preview={true}
				/>
			))}
		</div>
	)
}
