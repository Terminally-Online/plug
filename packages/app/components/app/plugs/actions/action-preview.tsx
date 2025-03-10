import { FC } from "react"

import { Sentence } from "@/components/app/plugs/sentences/sentence"
import { SchemasRequestActions } from "@/lib"
import { plugByIdAtom } from "@/state/plugs"
import { useAtom } from "jotai"

export const ActionPreview: FC<{ index: number; item: string; actions?: SchemasRequestActions; errors?: Array<string | null> }> = ({
	index,
	item,
	actions,
	errors = []
}) => {
	const [plug] = useAtom(plugByIdAtom(item))

	actions = actions ?? plug?.actions ?? []

	return (
		<div className="flex flex-col gap-2">
			{actions.map((action, actionIndex) => (
				<div key={`${index}-${actionIndex}`}>
					<Sentence
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
				</div>
			))}
		</div>
	)
}
