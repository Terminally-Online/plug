import { FC, useMemo } from "react"

import { Sentence } from "@/components"
import { usePlugs } from "@/contexts"
import { cn } from "@/lib"

const SentenceReview: FC<{ index: number; item: string; actionIndex: number; review: boolean }> = ({
	index,
	item,
	actionIndex,
	review
}) => {
	const { plug, actions } = usePlugs(item)

	const isReady = useMemo(
		() => plug && actions && actions[actionIndex].values.every(value => Boolean(value)),
		[plug, actions, actionIndex]
	)

	return (
		<>
			<Sentence
				className={cn(
					isReady ? "border-plug-green hover:border-plug-green" : "border-plug-red hover:border-plug-red"
				)}
				index={index}
				item={item}
				actionIndex={actionIndex}
				preview={true}
			/>

			{review && (
				<p className={cn("mt-1 text-sm font-bold", isReady ? "opacity-40" : "text-plug-red")}>
					{isReady ? "Ready" : "Missing required values"}
				</p>
			)}
		</>
	)
}

export const ActionPreview: FC<{ index: number; item: string; review?: boolean }> = ({
	index,
	item,
	review = false
}) => {
	const { actions } = usePlugs(item)

	return (
		<div className="flex flex-col gap-2">
			{actions.map((_, actionIndex) => (
				<div key={actionIndex} className="relative z-[4] flex flex-col">
					<SentenceReview index={index} item={item} actionIndex={actionIndex} review={review} />
				</div>
			))}
		</div>
	)
}
