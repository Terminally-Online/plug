import { FC, useMemo } from "react"

import { Sentence } from "@/components/app/plugs/sentences/sentence"
import { Actions } from "@/lib"
import { workflowByIdAtom } from "@/state/plugs"
import { useAtomValue } from "jotai"

export const ActionPreview: FC<{ index: number; item: string; actions?: Actions; errors?: Array<string | null> }> = ({
	index,
	item,
	actions,
	errors = []
}) => {
	const plug = useAtomValue(workflowByIdAtom)(item)
	const plugActions: Actions = useMemo(() => {
		if (!plug) return []
		try {
			return JSON.parse(plug.actions)
		} catch {
			return []
		}
	}, [plug])

	actions = actions ?? plugActions

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
