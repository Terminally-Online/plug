import { FC, useMemo } from "react"

import { Sentence } from "@/components"
import { usePlugs } from "@/contexts"
import { Action, Actions, cn } from "@/lib"

const SentenceReview: FC<{ index: number; item: string; actionIndex: number; action: Action }> = ({
	index,
	item,
	actionIndex,
	action
}) => {
	const { plug } = usePlugs(item)

	const isReady = useMemo(() => plug && action.values.every(value => Boolean(value)), [plug, action])

	return (
		<Sentence
			className={cn(
				isReady ? "border-plug-green hover:border-plug-green" : "border-plug-red hover:border-plug-red"
			)}
			index={index}
			item={item}
			actionIndex={actionIndex}
			action={action}
			preview={true}
		/>
	)
}

export const ActionPreview: FC<{ index: number; item: string; actions?: Actions }> = ({ index, item, actions }) => {
	const { actions: plugActions } = usePlugs(item)

	actions = actions ?? plugActions

	return (
		<div className="flex flex-col gap-2">
			{actions.map((action, actionIndex) => (
				<SentenceReview key={actionIndex} index={index} item={item} actionIndex={actionIndex} action={action} />
			))}
		</div>
	)
}
