import { FC, useMemo } from "react"

import { Sentence } from "@/components"
import { usePlugs } from "@/contexts"
import { cn } from "@/lib"

const SentenceReview: FC<{ index: number; item: string; actionIndex: number }> = ({ index, item, actionIndex }) => {
	const { plug, actions } = usePlugs(item)

	const isReady = useMemo(
		() => plug && actions && actions[actionIndex].values.every(value => Boolean(value)),
		[plug, actions, actionIndex]
	)

	return (
		<Sentence
			className={cn(
				isReady ? "border-plug-green hover:border-plug-green" : "border-plug-red hover:border-plug-red"
			)}
			index={index}
			item={item}
			actionIndex={actionIndex}
			preview={true}
		/>
	)
}

export const ActionPreview: FC<{ index: number; item: string }> = ({ index, item }) => {
	const { actions } = usePlugs(item)

	return (
		<div className="flex flex-col gap-2">
			{actions.map((_, actionIndex) => (
				<SentenceReview key={actionIndex} index={index} item={item} actionIndex={actionIndex} />
			))}
		</div>
	)
}
