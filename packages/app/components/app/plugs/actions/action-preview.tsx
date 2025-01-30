import { FC } from "react"

import { Sentence } from "@/components/app/plugs/sentences/sentence"
import { Actions } from "@/lib"
import { usePlugData } from "@/state/plugs"

export const ActionPreview: FC<{ index: number; item: string; actions?: Actions; errors?: Array<string | null> }> = ({
	index,
	item,
	actions,
	errors = []
}) => {
	const { actions: plugActions } = usePlugData(item)

	actions = actions ?? plugActions

	return (
		<div className="flex flex-col">
			{actions.map((action, actionIndex) => (
				<>
					<Sentence
						key={`${index}-${actionIndex}`}
						index={index}
						item={item}
						actionIndex={actionIndex}
						action={action}
						preview
						error={errors && errors[actionIndex] ? true : false}
					/>

					{errors && errors[actionIndex] && (
						<p className="text-sm font-bold text-plug-red">Error: {errors[actionIndex]}</p>
					)}
				</>
			))}
		</div>
	)
}
