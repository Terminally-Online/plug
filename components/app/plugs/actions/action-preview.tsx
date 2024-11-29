import { FC } from "react"

import { Sentence } from "@/components"
import { Action, Actions, cn } from "@/lib"
import { usePlugData } from "@/state"

const SentenceReview: FC<{ index: number; item: string; actionIndex: number; action: Action }> = ({
	index,
	item,
	actionIndex,
	action
}) => {
	const isReady = false
	// const isReady = useMemo(() => action.values.every(value => Boolean(value)), [action])

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
	const { actions: plugActions } = usePlugData(item)

	actions = actions ?? plugActions

	return (
		<div className="flex flex-col gap-2">
			{actions.map((action, actionIndex) => (
				<SentenceReview key={actionIndex} index={index} item={item} actionIndex={actionIndex} action={action} />
			))}
		</div>
	)
}
